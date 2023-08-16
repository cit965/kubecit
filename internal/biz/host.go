package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
	"kubecit/ent"

	"reflect"
	"time"
)

type CloudHost struct {
	UUID               string
	State              string
	IPV6AddressPrivate []string
	IPV4AddressPrivate []string
	IPV6AddressPublic  []string
	IPV4AddressPublic  []string
	Memory             int
	CPU                int
	CreatedTime        time.Time
	ExpiredTime        time.Time
	InstanceName       string
	ImageName          string
	OSType             string
	Manufacturer       string
	Zone               string
	SecurityGroups     []string
	BillType           string
	ChargeType         string
	IsActive           bool
	InstanceType       string
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
		return nil, err
	}
	return h, nil
}

// TODO: implement an interface to adapter multi cloud manufacturer
func (c *CloudHostUsecase) Syncer(ctx context.Context, accessKey string, secretKey string, region string) (bool, int64, error) {
	credential := common.NewCredential(accessKey, secretKey)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "cvm.tencentcloudapi.com"
	client, _ := cvm.NewClient(credential, region, cpf)

	request := cvm.NewDescribeInstancesRequest()
	response, err := client.DescribeInstances(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return false, 0, err
	}
	if err != nil {
		return false, 0, err
	}
	var count int64
	for _, instance := range response.Response.InstanceSet {
		var cloudHost CloudHost
		cloudHost.UUID = *instance.Uuid
		cloudHost.State = *instance.InstanceState
		if instance.PublicIpAddresses != nil {
			var iPV4AddressPublic []string
			for _, v := range instance.PublicIpAddresses {
				iPV4AddressPublic = append(iPV4AddressPublic, *v)
			}
			cloudHost.IPV4AddressPublic = iPV4AddressPublic
		}
		if instance.IPv6Addresses != nil {
			var iPV6AddressPublic []string
			for _, v := range instance.IPv6Addresses {
				iPV6AddressPublic = append(iPV6AddressPublic, *v)
			}
			cloudHost.IPV6AddressPublic = iPV6AddressPublic
		}
		var IPV4AddressPrivate []string
		for _, v := range instance.PrivateIpAddresses {
			IPV4AddressPrivate = append(IPV4AddressPrivate, *v)
		}
		cloudHost.IPV4AddressPrivate = IPV4AddressPrivate
		cloudHost.Memory = int(*instance.Memory)
		cloudHost.CPU = int(*instance.CPU)
		createdTime, _ := time.Parse("2006-01-02T15:04:05Z", *instance.CreatedTime)
		cloudHost.CreatedTime = createdTime.Local()
		if instance.ExpiredTime != nil {
			expiredTime, _ := time.Parse("2006-01-02T15:04:05Z", *instance.ExpiredTime)
			cloudHost.ExpiredTime = expiredTime.Local()
		}
		cloudHost.InstanceName = *instance.InstanceName
		cloudHost.OSType = *instance.OsName
		cloudHost.Zone = *instance.Placement.Zone
		var securityGroups []string
		for _, v := range instance.SecurityGroupIds {
			securityGroups = append(securityGroups, *v)
		}
		cloudHost.SecurityGroups = securityGroups
		cloudHost.ChargeType = *instance.InstanceChargeType
		cloudHost.InstanceType = *instance.InstanceType
		cloudHost.Manufacturer = "TencentCloud"
		if instance.RenewFlag != nil {
			cloudHost.BillType = *instance.RenewFlag
		}

		host, err := c.Get(ctx, cloudHost.UUID)
		if err != nil && ent.IsNotFound(err) {
			_, err := c.Create(ctx, &cloudHost)
			if err != nil {
				fmt.Println("create host error: ", cloudHost)
				continue
			}
			count++
		}

		if host != nil && reflect.DeepEqual(host, cloudHost) {
			_, err := c.Update(ctx, cloudHost.UUID, &cloudHost)
			if err != nil {
				fmt.Println("update host error: ", cloudHost)
				continue
			}
			count++
		}

	}
	return true, count, nil
}
