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
