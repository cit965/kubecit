package biz

import (
	"context"
	"fmt"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/go-kratos/kratos/v2/log"
)

type Cluster struct {
	ID         int
	Alias      string
	Kubeconfig string
}

type ClusterRepo interface {
	Register(ctx context.Context, cluster *Cluster) (*Cluster, error)
	List(ctx context.Context) ([]*Cluster, error)
	Get(ctx context.Context, id int) (*Cluster, error)
	Update(ctx context.Context, cluster *Cluster) (*Cluster, error)
	Delete(ctx context.Context, id int) error
}

type ClusterUseCase struct {
	repo ClusterRepo
	log  *log.Helper
}

func NewClusterUseCase(repo ClusterRepo, logger log.Logger) *ClusterUseCase {
	return &ClusterUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (c *ClusterUseCase) RegisterCluster(ctx context.Context, cluster *Cluster) (*Cluster, error) {
	clusters, err := c.repo.List(ctx)
	if err != nil {
		return nil, err
	}
	if len(clusters) > 0 {
		// check kubeconfig is exist
		for _, v := range clusters {
			if v.Kubeconfig == cluster.Kubeconfig {
				return nil, fmt.Errorf("kubeconfig is exist")
			}
		}
	}

	// Load the kubeconfig file
	config, err := clientcmd.BuildConfigFromFlags("", cluster.Kubeconfig)
	if err != nil {
		return nil, fmt.Errorf("error building kubeconfig: %s", err.Error())
	}

	// Create a Kubernetes clientset
	_, err = kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("error building kubernetes clientset: %s", err.Error())
	}

	clusterResult, err := c.repo.Register(ctx, cluster)
	if err != nil {
		return nil, err
	}
	return clusterResult, nil
}

func (c *ClusterUseCase) ListCluster(ctx context.Context) ([]*Cluster, error) {
	clusterResult, err := c.repo.List(ctx)
	if err != nil {
		return nil, err
	}
	return clusterResult, nil
}

func (c *ClusterUseCase) GetCluster(ctx context.Context, id int) (*Cluster, error) {
	clusterResult, err := c.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return clusterResult, nil
}

func (c *ClusterUseCase) UpdateCluster(ctx context.Context, cluster *Cluster) (*Cluster, error) {
	clusterResult, err := c.repo.Update(ctx, cluster)
	if err != nil {
		return nil, err
	}
	return clusterResult, nil
}

func (c *ClusterUseCase) Delete(ctx context.Context, id int) error {
	err := c.repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
