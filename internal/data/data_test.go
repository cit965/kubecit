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
	var co biz.ClusterRepo
	var ho biz.CloudHostRepo
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
		co = data.NewClusterRepo(d, nil)
		ho = data.NewCloudHostRepo(d, nil)
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
				user, _ := ro.Create(ctx, &biz.User{
					Username: "testUser",
					Password: "xx",
					Age:      0,
				})
				Expect(user.Username).To(Equal("testUser"))
				userList, err := ro.List(ctx)
				Expect(err).To(BeNil())
				Expect(len(userList)).To(Equal(1))
				err = ro.Delete(ctx, userList[0].Id)
				Expect(err).To(BeNil())
				userList2, err := ro.List(ctx)
				Expect(len(userList2)).To(Equal(0))
			})
		})
	})

	Describe("Test cluster", func() {
		Context("register,update,list,delete", func() {
			It("should be successful", func() {
				ctx := context.Background()
				log.Println("test register cluster")
				cluster, _ := co.Register(ctx, &biz.Cluster{
					Id:         1,
					Kubeconfig: "kubeconfig",
				})
				Expect(cluster.Kubeconfig).To(Equal("kubeconfig"))
				log.Println("test get cluster")
				c, _ := co.Get(ctx, cluster.Id)
				Expect(c.Kubeconfig).To(Equal(cluster.Kubeconfig))
				log.Println("test update cluster")
				cluster.Kubeconfig = "kubeconfig2"
				cluster2, err := co.Update(ctx, cluster)
				Expect(cluster2).To(Equal(cluster))
				log.Println("test list cluster")
				clusterList, err := co.List(ctx)
				Expect(err).To(BeNil())
				Expect(len(clusterList)).To(Equal(1))
				log.Println("test delete cluster")
				err = co.Delete(ctx, clusterList[0].Id)
				Expect(err).To(BeNil())
				userList2, err := co.List(ctx)
				Expect(len(userList2)).To(Equal(0))
			})
		})
	})

	Describe("Test cloudHost", func() {
		Context("create,get,list,update,delete,sync", func() {
			It("should be successful", func() {
				ctx := context.Background()
				log.Println("test create cloudHost instance")
				host, _ := ho.Create(ctx, &biz.CloudHost{
					VpcId:            "vpcId-foo-123",
					SubnetId:         "subnetId-bar-456",
					InstanceId:       "instanceId-xxx-xxx",
					InstanceName:     "test-instanceName",
					InstanceState:    "RUNNING",
					CPU:              2,
					Memory:           4096,
					CreatedTime:      "2023-08-21 17:22:12",
					InstanceType:     "small-v2",
					EniLimit:         10,
					EnilpLimit:       20,
					InstanceEniCount: 5,
				})
				Expect(host).To(Equal(&biz.CloudHost{
					VpcId:            "vpcId-foo-123",
					SubnetId:         "subnetId-bar-456",
					InstanceId:       "instanceId-xxx-xxx",
					InstanceName:     "test-instanceName",
					InstanceState:    "RUNNING",
					CPU:              2,
					Memory:           4096,
					CreatedTime:      "2023-08-21 17:22:12",
					InstanceType:     "small-v2",
					EniLimit:         10,
					EnilpLimit:       20,
					InstanceEniCount: 5,
				}))
				log.Println("test get cloudHost instance")
				host, _ = ho.Get(ctx, "instanceId-xxx-xxx")
				Expect(host).To(Equal(&biz.CloudHost{
					VpcId:            "vpcId-foo-123",
					SubnetId:         "subnetId-bar-456",
					InstanceId:       "instanceId-xxx-xxx",
					InstanceName:     "test-instanceName",
					InstanceState:    "RUNNING",
					CPU:              2,
					Memory:           4096,
					CreatedTime:      "2023-08-21 17:22:12",
					InstanceType:     "small-v2",
					EniLimit:         10,
					EnilpLimit:       20,
					InstanceEniCount: 5,
				}))
				log.Println("test update cloudHost instance")
				host, _ = ho.Update(ctx, "instanceId-xxx-xxx", &biz.CloudHost{
					CPU: 4,
				})
				Expect(host).To(Equal(&biz.CloudHost{
					VpcId:            "vpcId-foo-123",
					SubnetId:         "subnetId-bar-456",
					InstanceId:       "instanceId-xxx-xxx",
					InstanceName:     "test-instanceName",
					InstanceState:    "RUNNING",
					CPU:              4,
					Memory:           4096,
					CreatedTime:      "2023-08-21 17:22:12",
					InstanceType:     "small-v2",
					EniLimit:         10,
					EnilpLimit:       20,
					InstanceEniCount: 5,
				}))

				log.Println("test list cloudHost instances")
				hosts, _ := ho.List(ctx)
				Expect(hosts[0]).To(Equal(&biz.CloudHost{
					VpcId:            "vpcId-foo-123",
					SubnetId:         "subnetId-bar-456",
					InstanceId:       "instanceId-xxx-xxx",
					InstanceName:     "test-instanceName",
					InstanceState:    "RUNNING",
					CPU:              4,
					Memory:           4096,
					CreatedTime:      "2023-08-21 17:22:12",
					InstanceType:     "small-v2",
					EniLimit:         10,
					EnilpLimit:       20,
					InstanceEniCount: 5,
				}))

				log.Println("test delete cloudHost instances")
				host, _ = ho.Delete(ctx, "instanceId-xxx-xxx")
				Expect(host).To(Equal(&biz.CloudHost{
					VpcId:            "vpcId-foo-123",
					SubnetId:         "subnetId-bar-456",
					InstanceId:       "instanceId-xxx-xxx",
					InstanceName:     "test-instanceName",
					InstanceState:    "RUNNING",
					CPU:              4,
					Memory:           4096,
					CreatedTime:      "2023-08-21 17:22:12",
					InstanceType:     "small-v2",
					EniLimit:         10,
					EnilpLimit:       20,
					InstanceEniCount: 5,
				}))
			})
		})
	})
})

// 测试结束后 通过回调函数，关闭并删除 docker 创建的容器
var _ = AfterSuite(func() {
	cleaner()
})
