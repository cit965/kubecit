package data

import (
	"context"
	"kubecit/ent"
	"kubecit/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	_ "github.com/mattn/go-sqlite3"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewUserRepo, NewClusterRepo, NewK8sRepoGetter, NewCloudHostRepo)

// Data contains config and db client
type Data struct {
	conf *conf.Data
	db   *ent.Client
}

// NewData 构造方法，初始化了数据库 client
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}

	entClient, err := ent.Open(c.Database.Driver, c.Database.Source)
	if err != nil {
		log.Fatalf("fail to open connection to db,%s", err)
	}
	if err := entClient.Schema.Create(context.Background()); err != nil {
		log.Fatalf("fail to create schema,%s", err)
	}
	return &Data{
		conf: c,
		db:   entClient,
	}, cleanup, nil
}
