package biz

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
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

//go:generate mockgen -destination=../mocks/mrepo/host.go -package=mrepo . CloudHostRepo
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

type CloudProviderRepo interface {
	GetClient(ctx context.Context, accessKey, secretKey, region string) error
	ListInstancesByVpc(ctx context.Context, vpcId string) ([]*CloudHost, error)
}

type CloudProviderUsecase struct {
	repo CloudProviderRepo
	log  *log.Helper
}

func NewCloudProviderUsecase(repo CloudProviderRepo, logger log.Logger) *CloudProviderUsecase {
	return &CloudProviderUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (c *CloudProviderUsecase) GetClient(ctx context.Context, accessKey, secretKey, region string) error {
	err := c.repo.GetClient(ctx, accessKey, secretKey, region)
	return err
}

func (c *CloudProviderUsecase) ListInstancesByVpc(ctx context.Context, vpcId string) ([]*CloudHost, error) {
	return c.repo.ListInstancesByVpc(ctx, vpcId)
}
