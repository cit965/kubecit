package service

import (
	"context"
	"fmt"

	v1 "kubecit/api/helloworld/v1"
	"kubecit/internal/biz"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer

	uc          *biz.GreeterUsecase
	userCase    *biz.UserUsecase
	clusterCase *biz.ClusterUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase, userCase *biz.UserUsecase, clusterCase *biz.ClusterUsecase) *GreeterService {
	return &GreeterService{uc: uc, userCase: userCase, clusterCase: clusterCase}
}

// SayHello implements helloworld.GreeterServer.
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello niko" + g.Hello}, nil
}

// UserRegister register a user with username and password
func (s *GreeterService) UserRegister(ctx context.Context, in *v1.UserRegisterRequest) (*v1.UserRegisterResponse, error) {
	fmt.Println(in.Username, in.Password)
	_, err := s.userCase.RegisterUser(ctx, &biz.User{
		Username: in.Username,
		Password: in.Password,
	})
	if err != nil {
		return nil, err
	}
	return &v1.UserRegisterResponse{Result: "success"}, nil
}

func (s *GreeterService) UserList(ctx context.Context, in *v1.Empty) (*v1.UserListResponse, error) {
	users, err := s.userCase.UserList(ctx)
	if err != nil {
		return nil, err
	}
	userRes := []*v1.User{}

	for _, v := range users {
		userRes = append(userRes, &v1.User{
			Username: v.Username,
			Password: v.Password,
		})
	}
	return &v1.UserListResponse{Users: userRes}, nil
}

func (s *GreeterService) ClusterList(ctx context.Context, in *v1.Empty) (*v1.ClusterListResponse, error) {
	result, err := s.clusterCase.List(ctx)
	if err != nil {
		return nil, err
	}

	var res []*v1.Cluster
	for _, v := range result {
		tmp := &v1.Cluster{
			Kubeconfig: v.Kubeconfig,
			Id:         int32(v.Id),
		}
		res = append(res, tmp)
	}
	return &v1.ClusterListResponse{Clusters: res}, nil
}
