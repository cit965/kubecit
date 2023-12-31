package biz_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"kubecit/internal/biz"
	"kubecit/internal/mocks/mrepo"
)

var _ = Describe("UserUsecase", func() {
	var userCase *biz.UserUsecase
	var mUserRepo *mrepo.MockUserRepo

	BeforeEach(func() {
		mUserRepo = mrepo.NewMockUserRepo(ctl)
		userCase = biz.NewUserUsecase(mUserRepo, nil)
	})

	It("register", func() {
		info := &biz.User{
			Username: "xxx",
			Password: "admin123456",
		}
		mUserRepo.EXPECT().Create(ctx, gomock.Any()).Return(info, nil)
		l, err := userCase.RegisterUser(ctx, info)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(err).ToNot(HaveOccurred())
		Ω(l.Username).To(Equal("xxx"))
	})

	It("list", func() {
		info := &biz.User{
			Username: "xxx",
			Password: "admin123456",
		}
		users := []*biz.User{
			info,
		}
		mUserRepo.EXPECT().List(ctx).Return(users, nil)
		l, err := userCase.UserList(ctx)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(err).ToNot(HaveOccurred())
		Ω(len(l)).To(Equal(1))
	})
})
