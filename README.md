# kubecit

多云管理平台，支持购买不同云厂商的云资源，支持资源发现和同步，支持k8s多集群管理

## 前置知识

- go v1.21
- kratos 框架
- k8s 基础知识
- cmdb

## 命令

在根目录执行 `make init` 安装 wire proto 等依赖

在根目录执行 `make run` 启动服务

在根目录执行 `make test` 进行单元测试，要求本地环境变量设置 docker endpoint , 例子如下

```shell
export DOCKER_HOST=unix:///Users/z/.docker/run/docker.sock & make test
```

## 生成表

1. 想要添加一张 cluster 表，可以执行如下命令：

```shell
go run -mod=mod entgo.io/ent/cmd/ent new Cluster
```

2. 在 scheme/cluster.go 文件中添加表字段如下：
```go
// Fields of the Cluster.
func (Cluster) Fields() []ent.Field {
	return []ent.Field{
		field.String("kubeconfig").
			Default("unknown"),
	}
}
```
3. 执行如下命令，生成 CRUD 代码

```shell
 go generate ./ent
```
