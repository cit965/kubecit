package biz_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"kubecit/internal/biz"
	"kubecit/internal/mocks/mrepo"
)

var _ = Describe("cloudHostUsecase", func() {
	var cloudHostCase *biz.CloudHostUsecase
	var mCloudHostRepo *mrepo.MockCloudHostRepo

	BeforeEach(func() {
		mCloudHostRepo = mrepo.NewMockCloudHostRepo(ctl)
		cloudHostCase = biz.NewCloudHostUsecase(mCloudHostRepo, nil)
	})

	It("create", func() {
		host := &biz.CloudHost{
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
		}
		mCloudHostRepo.EXPECT().Create(ctx, gomock.Any()).Return(host, nil)
		l, err := cloudHostCase.Create(ctx, host)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(err).ToNot(HaveOccurred())
		Ω(l).To(Equal(l))
	})

	It("get", func() {
		host := &biz.CloudHost{
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
		}
		mCloudHostRepo.EXPECT().Get(ctx, host.InstanceId).Return(host, nil)
		h, err := cloudHostCase.Get(ctx, host.InstanceId)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(err).ToNot(HaveOccurred())
		Ω(h).To(Equal(host))
	})

	It("list", func() {
		host := &biz.CloudHost{
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
		}
		hosts := []*biz.CloudHost{
			host,
		}
		mCloudHostRepo.EXPECT().List(ctx).Return(hosts, nil)
		l, err := cloudHostCase.List(ctx)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(err).ToNot(HaveOccurred())
		Ω(len(l)).To(Equal(1))
		Ω(l[0]).To(Equal(host))
	})

	It("update", func() {
		host := &biz.CloudHost{
			InstanceId: "instanceId-xxx-xxx",
			CPU:        4,
		}

		mCloudHostRepo.EXPECT().Update(ctx, host.InstanceId, host).Return(&biz.CloudHost{
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
		}, nil)
		res, err := cloudHostCase.Update(ctx, host.InstanceId, host)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(err).ToNot(HaveOccurred())
		Ω(res).To(Equal(&biz.CloudHost{
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

	It("delete", func() {
		host := &biz.CloudHost{
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
		}

		mCloudHostRepo.EXPECT().Delete(ctx, host.InstanceId).Return(host, nil)
		l, err := cloudHostCase.Delete(ctx, host.InstanceId)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(err).ToNot(HaveOccurred())
		Ω(l).To(Equal(host))
	})
})
