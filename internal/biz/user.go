package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Id       int
	Username string
	Password string
	Age      int
}

// UserRepo 接口，定义了 data 层需要提供的能力，此接口实现者为 data/user.go 文件中的 userRepo
//
//go:generate mockgen -destination=../mocks/mrepo/user.go -package=mrepo . UserRepo
type UserRepo interface {
	Create(context.Context, *User) (*User, error)
	List(ctx context.Context) ([]*User, error)
	Delete(ctx context.Context, id int) error
}

// UserUsecase 用户领域结构体，可以包含多个与用户业务相关的 repo
type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

// NewUserUsecase 用户领域构造方法
func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

// RegisterUser 注册一个用户
func (u *UserUsecase) RegisterUser(ctx context.Context, user *User) (*User, error) {
	userResult, err := u.repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return userResult, nil
}

// UserList 列出所有用户
func (u *UserUsecase) UserLi(ctx context.Context) ([]*User, error) {
	userResult, err := u.repo.List(ctx)
	if err != nil {
		return nil, err
	}
	return userResult, nil
}
