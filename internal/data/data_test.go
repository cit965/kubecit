package data_test

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/ory/dockertest/v3"
	"kubecit/internal/biz"
	"kubecit/internal/conf"
	"kubecit/internal/data"
	"log"
	"os"
)

var db *sql.DB
var cleaner func()

var _ = Describe("Data", func() {
	var ro biz.UserRepo
	BeforeEach(func() {
		dockerEndpoint := os.Getenv("DOCKER_HOST")
		pool, err := dockertest.NewPool(dockerEndpoint)
		if err != nil {
			log.Fatalf("Could not construct pool: %s", err)
		}
		// uses pool to try to connect to Docker
		err = pool.Client.Ping()
		if err != nil {
			log.Fatalf("Could not connect to Docker: %s", err)
		}

		// pulls an image, creates a container based on it and runs it
		resource, err := pool.Run("mariadb", "latest", []string{"MYSQL_ROOT_PASSWORD=secret"})
		if err != nil {
			log.Fatalf("Could not start resource: %s", err)
		}
		// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
		if err := pool.Retry(func() error {
			var err error
			db, err = sql.Open("mysql", fmt.Sprintf("root:secret@(localhost:%s)/mysql", resource.GetPort("3306/tcp")))
			if err != nil {
				return err
			}
			return db.Ping()
		}); err != nil {
			log.Fatalf("Could not connect to database: %s", err)
		}
		c := &conf.Data{
			Database: &conf.Data_Database{
				Driver: "mysql",
				Source: fmt.Sprintf("root:secret@(localhost:%s)/mysql", resource.GetPort("3306/tcp")),
			},
		}
		d, _, err := data.NewData(c, nil)
		if err != nil {
			log.Fatalf("Cound not NewData: %s", err)
		}

		ro = data.NewUserRepo(d, nil)
		// You can't defer this because os.Exit doesn't care for defer
		cleaner = func() {
			if err := pool.Purge(resource); err != nil {
				log.Fatalf("Could not purge resource: %s", err)
			}
		}

	})

	Describe("Categorizing book length", func() {
		Context("With more than 300 pages", func() {
			It("should be a novel", func() {
				ctx := context.Background()
				user, _ := ro.Register(ctx, &biz.User{
					Username: "testUser",
					Password: "xx",
					Age:      0,
				})
				Expect(user.Username).To(Equal("testUser"))
				userList, err := ro.List(ctx)
				Expect(err).To(BeNil())
				Expect(len(userList)).To(Equal(1))
			})
		})
	})
})

// 测试结束后 通过回调函数，关闭并删除 docker 创建的容器
var _ = AfterSuite(func() {
	cleaner()
})
