package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	corev1 "k8s.io/api/core/v1"
)

type Cluster struct {
	Id         int
	Kubeconfig string
}

// ClusterUsecase 集群领域结构体，可以包含多个与用户业务相关的 repo
type ClusterUsecase struct {
	repo   ClusterRepo
	log    *log.Helper
	getter K8sRepoGetter
}

type ClusterRepo interface {
	Get(ctx context.Context, id int) (*Cluster, error)

	List(ctx context.Context) ([]*Cluster, error)
}

type K8sRepo interface {
	ListNamespace(ctx context.Context) (*corev1.NamespaceList, error)
}
type K8sRepoGetter interface {
	GetRepo(kubeCfg []byte, help *log.Helper) (K8sRepo, error)
}

// NewClusterUsecase 集群领域构造方法
func NewClusterUsecase(repo ClusterRepo, getter K8sRepoGetter, logger log.Logger) *ClusterUsecase {
	return &ClusterUsecase{repo: repo, getter: getter, log: log.NewHelper(logger)}
}

func (c *ClusterUsecase) List(ctx context.Context) ([]*Cluster, error) {
	return c.repo.List(ctx)
}

func (c *ClusterUsecase) ListNamespaces(ctx context.Context, id int) ([]string, error) {
	repo, err := c.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	k8sRepo, err := c.getter.GetRepo([]byte(repo.Kubeconfig), c.log)
	if err != nil {
		return nil, err
	}
	namespaceList, err := k8sRepo.ListNamespace(ctx)
	if err != nil {
		return nil, err
	}
	var result []string
	for _, v := range namespaceList.Items {
		result = append(result, v.Name)
	}
	return result, nil
}
