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

1. 想要添加一张 clusters 表，可以执行如下命令：

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

## mock 测试

1. 在你需要mock掉的inteface上加一下注解

```go
//go:generate mockgen -destination=../mocks/mrepo/user.go -package=mrepo . UserRepo
```
2. 在 biz 目录下执行以下命令，internal/mocks 目录下生成 mock 代码

```go
mockgen -destination=../mocks/mrepo/user.go -package=mrepo . UserRepo
```

3. 在你需要替换的地方使用如下：

```go
	BeforeEach(func() {
	    // 使用 gomock 替换 userRepo
		mUserRepo = mrepo.NewMockUserRepo(ctl)
		userCase = biz.NewUserUsecase(mUserRepo, nil)
	})

	It("Create", func() {
		info := &biz.User{
			Username: "xxx",
			Password: "admin123456",
		}
		// 设置返回值
		mUserRepo.EXPECT().Register(ctx, gomock.Any()).Return(info, nil)
		l, err := userCase.RegisterUser(ctx, info)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(err).ToNot(HaveOccurred())
		Ω(l.Username).To(Equal("xxx"))
	})
```
