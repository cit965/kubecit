package biz_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"kubecit/internal/biz"
	"kubecit/internal/mocks/mrepo"
)

var _ = Describe("Cluster", func() {
	var clusterCase *biz.ClusterUsecase
	var mClusterRepo *mrepo.MockClusterRepo

	BeforeEach(func() {
		mClusterRepo = mrepo.NewMockClusterRepo(ctl)
		clusterCase = biz.NewClusterUsecase(mClusterRepo, nil, nil)
	})

	It("register", func() {
		info := &biz.Cluster{
			Id:         1,
			Kubeconfig: "kubeconfig",
		}
		mClusterRepo.EXPECT().Register(ctx, gomock.Any()).Return(info, nil)
		l, err := clusterCase.RegisterCluster(ctx, info)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(err).ToNot(HaveOccurred())
		Ω(l).To(Equal(l))
	})

	It("list", func() {
		info := &biz.Cluster{
			Id:         1,
			Kubeconfig: "kubeconfig",
		}
		clusters := []*biz.Cluster{
			info,
		}
		mClusterRepo.EXPECT().List(ctx).Return(clusters, nil)
		l, err := clusterCase.List(ctx)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(err).ToNot(HaveOccurred())
		Ω(len(l)).To(Equal(1))
	})

	It("delete", func() {
		id := 1

		mClusterRepo.EXPECT().Delete(ctx, id).Return(nil)
		err := clusterCase.Delete(ctx, id)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(err).ToNot(HaveOccurred())
	})

	It("get", func() {
		info := &biz.Cluster{
			Id:         1,
			Kubeconfig: "kubeconfig",
		}

		mClusterRepo.EXPECT().Get(ctx, info.Id).Return(info, nil)
		l, err := clusterCase.GetCluster(ctx, info.Id)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(err).ToNot(HaveOccurred())
		Ω(l).To(Equal(info))
	})

	It("update", func() {
		info := &biz.Cluster{
			Id:         1,
			Kubeconfig: "kubeconfig",
		}

		mClusterRepo.EXPECT().Update(ctx, info).Return(info, nil)
		l, err := clusterCase.UpdateCluster(ctx, info)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(err).ToNot(HaveOccurred())
		Ω(l).To(Equal(info))
	})
})
