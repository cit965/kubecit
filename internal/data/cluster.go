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
	clu, err := c.data.db.Cluster.Get(ctx, id)
	return &biz.Cluster{Kubeconfig: clu.Kubeconfig}, err
}
