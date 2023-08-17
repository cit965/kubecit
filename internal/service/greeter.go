package service

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"kubecit/internal/data"
	"time"

	v1 "kubecit/api/helloworld/v1"
	"kubecit/internal/biz"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer

	uc            *biz.GreeterUsecase
	userCase      *biz.UserUsecase
	clusterCase   *biz.ClusterUsecase
	cloudHostCase *biz.CloudHostUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase, userCase *biz.UserUsecase, clusterCase *biz.ClusterUsecase, cloudHostCase *biz.CloudHostUsecase) *GreeterService {
	return &GreeterService{uc: uc, userCase: userCase, clusterCase: clusterCase, cloudHostCase: cloudHostCase}
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

func (s *GreeterService) NamespaceList(ctx context.Context, in *v1.NamespaceReq) (*v1.NamespaceResp, error) {
	namespaces, err := s.clusterCase.ListNamespaces(ctx, int(in.Cluster))
	if err != nil {
		return nil, err
	}
	return &v1.NamespaceResp{Namespaces: namespaces}, nil
}

func (s *GreeterService) DeploymentList(ctx context.Context, in *v1.DeploymentReq) (*v1.DeploymentResp, error) {
	deployments, err := s.clusterCase.ListDeployments(ctx, int(in.Cluster), in.Namespace)
	if err != nil {
		return nil, err
	}
	return &v1.DeploymentResp{Deployments: deployments}, nil
}

func (s *GreeterService) GetInstance(ctx context.Context, in *v1.GetInstanceRequest) (*v1.GetInstanceReply, error) {
	cloudHost, err := s.cloudHostCase.Get(ctx, in.UUID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Printf("%#v", cloudHost)
	host := &v1.Host{}
	err = data.ConvertType(cloudHost, host)
	if err != nil {
		return nil, err
	}
	host.CreatedTime = timestamppb.New(cloudHost.CreatedTime)
	host.ExpitedTime = timestamppb.New(cloudHost.ExpiredTime)
	res := &v1.GetInstanceReply{
		Instance: host,
	}
	return res, nil
}

func (s *GreeterService) CreateInstance(ctx context.Context, in *v1.CreateInstanceRequest) (*v1.CreateInstanceReply, error) {
	host := &biz.CloudHost{
		UUID:               in.Instance.UUID,
		State:              in.Instance.State,
		IPV6AddressPrivate: in.Instance.IPV6AddressPrivate,
		IPV4AddressPrivate: in.Instance.IPV4AddressPrivate,
		IPV6AddressPublic:  in.Instance.IPV6AddressPublic,
		IPV4AddressPublic:  in.Instance.IPV4AddressPublic,
		Memory:             int(in.Instance.Memory),
		CPU:                int(in.Instance.CPU),
		CreatedTime:        time.Now(),
		ExpiredTime:        time.Time{},
		InstanceName:       in.Instance.InstanceName,
		ImageName:          in.Instance.ImageName,
		OSType:             in.Instance.OSType,
		Manufacturer:       in.Instance.Manufacturer,
		Zone:               in.Instance.Zone,
		SecurityGroups:     in.Instance.SecurityGroups,
		BillType:           in.Instance.BillType,
		ChargeType:         in.Instance.ChargeType,
		IsActive:           in.Instance.IsActive,
		InstanceType:       in.Instance.InstanceType,
	}

	_, err := s.cloudHostCase.Create(ctx, host)
	if err != nil {
		return nil, err
	}
	res := &v1.CreateInstanceReply{
		Instance: in.Instance,
	}

	return res, err
}
func (s *GreeterService) ListInstances(ctx context.Context, in *v1.ListInstancesRequest) (*v1.ListInstancesReply, error) {
	cloudHosts, err := s.cloudHostCase.List(ctx)
	if err != nil {
		return nil, err
	}
	res := &v1.ListInstancesReply{Total: int64(len(cloudHosts))}
	for _, v := range cloudHosts {
		instance := &v1.Host{}
		data.ConvertType(v, instance)
		instance.CreatedTime = timestamppb.New(v.CreatedTime)
		instance.ExpitedTime = timestamppb.New(v.ExpiredTime)
		res.Instances = append(res.Instances, instance)
	}
	return res, nil
}

// TODO
func (s *GreeterService) DeleteInstanceById(ctx context.Context, in *v1.DeleteInstanceRequest) (*v1.DeleteInstanceReply, error) {
	cloudHost, err := s.cloudHostCase.Delete(ctx, in.UUID)
	if err != nil {
		return nil, err
	}
	res := &v1.DeleteInstanceReply{
		Message: fmt.Sprintf("host %v delete success.", cloudHost.UUID),
	}
	return res, nil
}

// TODO
func (s *GreeterService) UpdateInstance(ctx context.Context, in *v1.UpdateInstanceRequest) (*v1.UpdateInstanceReply, error) {
	var host biz.CloudHost
	err := data.ConvertType(in.Instance, &host)
	if err != nil {
		return nil, err
	}
	res := &v1.UpdateInstanceReply{}
	instance, err := s.cloudHostCase.Update(ctx, in.UUID, &host)
	if err != nil {
		return nil, err
	}

	updatedHost := &v1.Host{}
	err = data.ConvertType(instance, updatedHost)
	if err != nil {
		return nil, err
	}
	res.Message = "update instance success"
	res.Instance = updatedHost
	return res, nil
}

func (s *GreeterService) SyncFromTencent(ctx context.Context, in *v1.SyncFromTencentRequest) (*v1.SyncFromTencentReply, error) {
	ok, total, err := s.cloudHostCase.Syncer(ctx, in.AccessKey, in.SecretKey, in.Region)
	if !ok || err != nil {
		return nil, err
	}
	return &v1.SyncFromTencentReply{Message: "sync finished", Total: total}, nil
}
