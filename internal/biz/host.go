package biz

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	vpc "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vpc/v20170312"

	"kubecit/ent"
)

type CloudHost struct {
	VpcId            string
	SubnetId         string
	InstanceId       string
	InstanceName     string
	InstanceState    string
	CPU              int64
	Memory           int64
	CreatedTime      string
	InstanceType     string
	EniLimit         int64
	EnilpLimit       int64
	InstanceEniCount int64
}

type CloudHostRepo interface {
	Get(ctx context.Context, id string) (*CloudHost, error)
	Create(ctx context.Context, host *CloudHost) (*CloudHost, error)
	List(ctx context.Context) ([]*CloudHost, error)
	Delete(ctx context.Context, id string) (*CloudHost, error)
	Update(ctx context.Context, id string, host *CloudHost) (*CloudHost, error)
	Sync(ctx context.Context) (bool, error)
}

// CloudHostUsecase is a CloudHost usecase.
type CloudHostUsecase struct {
	repo CloudHostRepo
	log  *log.Helper
}

// NewCloudHostUsecase new a CloudHostRepo usecase.
func NewCloudHostUsecase(repo CloudHostRepo, logger log.Logger) *CloudHostUsecase {
	return &CloudHostUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (c *CloudHostUsecase) Get(ctx context.Context, id string) (*CloudHost, error) {
	h, err := c.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return h, nil
}

func (c *CloudHostUsecase) Create(ctx context.Context, host *CloudHost) (*CloudHost, error) {
	h, err := c.repo.Create(ctx, host)
	if err != nil {
		return nil, err
	}
	return h, nil
}

func (c *CloudHostUsecase) List(ctx context.Context) ([]*CloudHost, error) {
	hs, err := c.repo.List(ctx)
	if err != nil {
		return nil, err
	}
	return hs, nil
}

func (c *CloudHostUsecase) Delete(ctx context.Context, id string) (*CloudHost, error) {
	h, err := c.repo.Delete(ctx, id)
	if err != nil {
		return nil, err
	}
	return h, nil
}

func (c *CloudHostUsecase) Update(ctx context.Context, id string, host *CloudHost) (*CloudHost, error) {
	h, err := c.repo.Update(ctx, id, host)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return h, nil
}

// TODO: implement an interface to adapter multi cloud manufacturer
func (c *CloudHostUsecase) Syncer(ctx context.Context, accessKey string, secretKey string, region string, vpcId string) (bool, int64, error) {
	credential := common.NewCredential(accessKey, secretKey)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "vpc.tencentcloudapi.com"
	client, _ := vpc.NewClient(credential, region, cpf)

	request := vpc.NewDescribeVpcInstancesRequest()
	request.Filters = []*vpc.Filter{
		&vpc.Filter{
			Name:   common.StringPtr("vpc-id"),
			Values: common.StringPtrs([]string{vpcId}),
		},
	}
	response, err := client.DescribeVpcInstances(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return false, 0, err
	}

	var count int64
	for _, instance := range response.Response.InstanceSet {
		var cloudHost CloudHost
		cloudHost.VpcId = *instance.VpcId
		cloudHost.SubnetId = *instance.SubnetId
		cloudHost.InstanceId = *instance.InstanceId
		cloudHost.InstanceName = *instance.InstanceName
		cloudHost.InstanceState = *instance.InstanceState
		cloudHost.CPU = int64(*instance.CPU)
		cloudHost.Memory = int64(*instance.Memory)
		cloudHost.CreatedTime = *instance.CreatedTime
		cloudHost.InstanceType = *instance.InstanceType
		cloudHost.EniLimit = int64(*instance.EniLimit)
		cloudHost.EnilpLimit = int64(*instance.EniIpLimit)
		cloudHost.InstanceEniCount = int64(*instance.InstanceEniCount)

		_, err := c.Get(ctx, cloudHost.InstanceId)
		if err != nil && ent.IsNotFound(err) {
			_, err := c.Create(ctx, &cloudHost)
			if err != nil {
				fmt.Println("create host error: ", cloudHost)
				continue
			}
			count++
		}

	}
	return true, count, nil
}
