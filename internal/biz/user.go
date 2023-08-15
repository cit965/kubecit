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
//
//go:generate mockgen -destination=../mocks/mrepo/user.go -package=mrepo . UserRepo
type UserRepo interface {
	Create(context.Context, *User) (*User, error)

	List(ctx context.Context) ([]*User, error)
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
	userResult, err := u.repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return userResult, nil
}

func (u *UserUsecase) UserList(ctx context.Context) ([]*User, error) {
	userResult, err := u.repo.List(ctx)
	if err != nil {
		return nil, err
	}
	return userResult, nil
}
