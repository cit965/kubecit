## ginkgo

Ginkgo 是一个 Go 测试框架，旨在帮助你有效地编写富有表现力的全方位测试。

### 入门: 第一个测试

Ginkgo与Go现有的测试基础设施挂钩. 这允许您使用 go test 运行Ginkgo套件.
这同时意味着 Ginkgo 测试可以和传统 Go testing 测试一起使用。go test 和 ginkgo 都会运行你套件内的所有测试。

要为一个包写 Ginkgo 测试的话你必须首先初始化一个 Ginkgo 测试套件。比如说你有一个名为 data 的包：

```shell
$ cd internal/data
$ ginkgo bootstrap
```

我们将生成一个名为 data_suite_test.go 的文件并包含：

```go
package data_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestData(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Data Suite")
}
```

一个空的测试套件不是非常有趣。在你可以开始添加测试到 data_suite_test.go 的时候，你很可能偏向把测试放在多个文件中 （特别是有多个文件的包）。我们添加一个测试文件到我们的 data.go 模型：
```shell
ginkgo generate data
```

这将生成一个名为 data_test.go 的文件并包含：

```go
package data_test

import (
    "/path/to/data"
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
)

var _ = Describe("Data", func() {
	BeforeEach(func() {})
	AfterSuite()
})
```

- 你应该使用 Describe 和 Context 容器来富有表现力地组织你代码的行为。
- 你可使用 BeforeEach 为你的 specs 初始化状态。使用 It 指定一个 spec。
- 要在 BeforeEach 和 It 分享状态的话你可以使用闭包变量，一般声明在 Describe 和 Context 容器最近的顶层。



## repo 单元测试

本节我们想测试 userRepo 代码的正确性(操作数据库),由于 repo 依赖 data，我们需要mock掉data。

- 在测试执行前临时启动一个数据库镜像,使用临时镜像可以保证每次测试运行的幂等性。
- 测试时候需要指定`DOCKER_HOST`环境变量。

## biz 单元测试

本节我们想测试 userUsecase 代码的正确性(注册用户逻辑)，由于 usecase 依赖 repo，我们需要mock掉repo。

- 使用 gomock 库，在 repo 上添加注解
- 执行 mockgen 命令生成 mock 代码
- 使用 生成的代码来初始化 userUsecase


