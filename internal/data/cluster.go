package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kubecit/internal/biz"
)

type clusterRepo struct {
	data *Data
	log  *log.Helper
}

// NewClusterRepo 用户数据仓库构造方法
func NewClusterRepo(data *Data, logger log.Logger) biz.ClusterRepo {
	return &clusterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (c *clusterRepo) Get(ctx context.Context, id int) (*biz.Cluster, error) {
	clu, err := c.data.db.Cluster.Get(context.Background(), id)
	if err != nil {
		return nil, err
	}
	log.Info(clu)
	return &biz.Cluster{Kubeconfig: clu.Kubeconfig}, err
}

func (c *clusterRepo) List(ctx context.Context) ([]*biz.Cluster, error) {
	res, err := c.data.db.Cluster.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	var result []*biz.Cluster
	for _, v := range res {
		tmp := &biz.Cluster{Kubeconfig: v.Kubeconfig, Id: v.ID}
		result = append(result, tmp)
	}
	return result, nil
}

func (c *clusterRepo) Register(ctx context.Context, cluster *biz.Cluster) (*biz.Cluster, error) {
	clusterEnt, err := c.data.db.Cluster.Create().SetKubeconfig(cluster.Kubeconfig).Save(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Cluster{
		Id:         clusterEnt.ID,
		Kubeconfig: clusterEnt.Kubeconfig,
	}, nil
}

func (c *clusterRepo) Update(ctx context.Context, cluster *biz.Cluster) (*biz.Cluster, error) {
	data, err := c.data.db.Cluster.UpdateOneID(int(cluster.Id)).SetKubeconfig(cluster.Kubeconfig).Save(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Cluster{
		Id:         data.ID,
		Kubeconfig: data.Kubeconfig,
	}, nil
}

func (c *clusterRepo) Delete(ctx context.Context, id int) error {
	return c.data.db.Cluster.DeleteOneID(id).Exec(ctx)
}
