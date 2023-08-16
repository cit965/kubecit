package data

import (
	"context"
	"kubecit/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type clusterRepo struct {
	data *Data
	log  *log.Helper
}

func NewClusterRepo(data *Data, logger log.Logger) biz.ClusterRepo {
	return &clusterRepo{data: data, log: log.NewHelper(logger)}
}

func (c *clusterRepo) Register(ctx context.Context, cluster *biz.Cluster) (*biz.Cluster, error) {
	clusterEnt, err := c.data.db.Cluster.Create().SetAlias(cluster.Alias).SetKubeconfig(cluster.Kubeconfig).Save(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Cluster{
		ID:         clusterEnt.ID,
		Alias:      clusterEnt.Alias,
		Kubeconfig: clusterEnt.Kubeconfig,
	}, nil
}

func (c *clusterRepo) List(ctx context.Context) ([]*biz.Cluster, error) {
	clusters, err := c.data.db.Cluster.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	var clusterResults []*biz.Cluster
	for _, cluster := range clusters {
		clusterResults = append(clusterResults, &biz.Cluster{
			ID:         cluster.ID,
			Alias:      cluster.Alias,
			Kubeconfig: cluster.Kubeconfig,
		})
	}
	return clusterResults, nil
}

func (c *clusterRepo) Get(ctx context.Context, id int) (*biz.Cluster, error) {
	cluster, err := c.data.db.Cluster.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &biz.Cluster{
		ID:         cluster.ID,
		Alias:      cluster.Alias,
		Kubeconfig: cluster.Kubeconfig,
	}, nil
}

func (c *clusterRepo) Update(ctx context.Context, cluster *biz.Cluster) (*biz.Cluster, error) {
	data, err := c.data.db.Cluster.UpdateOneID(int(cluster.ID)).SetAlias(cluster.Alias).SetKubeconfig(cluster.Kubeconfig).Save(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Cluster{
		ID:         data.ID,
		Alias:      data.Alias,
		Kubeconfig: data.Kubeconfig,
	}, nil
}

func (c *clusterRepo) Delete(ctx context.Context, id int) error {
	return c.data.db.Cluster.DeleteOneID(id).Exec(ctx)
}
