package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Cluster struct {
	Id         int
	Kubeconfig string
}

// ClusterUsecase 集群领域结构体，可以包含多个与用户业务相关的 repo
type ClusterUsecase struct {
	repo ClusterRepo
	log  *log.Helper
}

type ClusterRepo interface {
	Get(ctx context.Context, id int) (*Cluster, error)

	List(ctx context.Context) ([]*Cluster, error)
}

// NewClusterUsecase 集群领域构造方法
func NewClusterUsecase(repo ClusterRepo, logger log.Logger) *ClusterUsecase {
	return &ClusterUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (c *ClusterUsecase) List(ctx context.Context) ([]*Cluster, error) {
	return c.repo.List(ctx)
}
