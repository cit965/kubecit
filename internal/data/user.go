package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"kubecit/internal/biz"
)

// userRepo 实现了 biz 层 UserRepo interface
type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo 用户数据仓库构造方法
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// Create 在用户表插入一个用户，注意返回值 为 biz.User
func (u *userRepo) Create(ctx context.Context, user *biz.User) (*biz.User, error) {
	userEnt, err := u.data.db.User.Create().SetName(user.Username).SetAge(1).SetPassword(user.Password).Save(ctx)
	if err != nil {
		fmt.Println(err)
	}
	return &biz.User{
		Username: userEnt.Name,
		Password: userEnt.Password,
		Age:      userEnt.Age,
	}, nil
}

// List 列出用户表所有用户
func (u *userRepo) List(ctx context.Context) ([]*biz.User, error) {
	users, err := u.data.db.User.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	var userResults []*biz.User
	for _, user := range users {
		userResults = append(userResults, &biz.User{
			Username: user.Name,
			Password: user.Password,
			Age:      user.Age,
		})
	}
	return userResults, nil
}
