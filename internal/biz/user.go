package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Username string
	Password string
	Age      int
}

// UserRepo is a Greater repo.
type UserRepo interface {
	Register(context.Context, *User) (*User, error)
}

// UserUsecase is a User usecase.
type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (u *UserUsecase) RegisterUser(ctx context.Context, user *User) (*User, error) {
	userResult, err := u.repo.Register(ctx, user)
	if err != nil {
		return nil, err
	}
	return userResult, nil
}
