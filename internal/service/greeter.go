package service

import (
	"context"
	"fmt"

	v1 "kubecit/api/helloworld/v1"
	"kubecit/internal/biz"
	"kubecit/internal/data"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer

	uc                *biz.GreeterUsecase
	userCase          *biz.UserUsecase
	clusterCase       *biz.ClusterUsecase
	cloudHostCase     *biz.CloudHostUsecase
	cloudProviderCase *biz.CloudProviderUsecase
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

func (s *GreeterService) ClusterGet(ctx context.Context, in *v1.ClusterBase) (reply *v1.ClusterBase, err error) {
	cluster, err := s.clusterCase.GetCluster(ctx, int(in.Id))
	if err != nil {
		return nil, err
	}
	return &v1.ClusterBase{
		Id:         int64(cluster.Id),
		Kubeconfig: cluster.Kubeconfig,
	}, nil
}

func (s *GreeterService) ClusterRegister(ctx context.Context, in *v1.ClusterKubeconfig) (*v1.ClusterBase, error) {
	result, err := s.clusterCase.RegisterCluster(ctx, &biz.Cluster{
		Kubeconfig: in.Kubeconfig,
	})
	if err != nil {
		return nil, err
	}
	return &v1.ClusterBase{
		Id:         int64(result.Id),
		Kubeconfig: result.Kubeconfig,
	}, nil
}

func (s *GreeterService) ClusterUpdate(ctx context.Context, in *v1.ClusterBase) (reply *v1.ClusterBase, err error) {
	cluster, err := s.clusterCase.UpdateCluster(ctx, &biz.Cluster{
		Id:         int(in.Id),
		Kubeconfig: in.Kubeconfig,
	})
	if err != nil {
		return nil, err
	}
	return &v1.ClusterBase{
		Id:         int64(cluster.Id),
		Kubeconfig: cluster.Kubeconfig,
	}, nil
}

func (s *GreeterService) ClusterDelete(ctx context.Context, in *v1.ClusterBase) (reply *v1.Empty, err error) {
	err = s.clusterCase.Delete(ctx, int(in.Id))
	if err != nil {
		return nil, err
	}
	return &v1.Empty{}, nil
}
func (s *GreeterService) GetInstance(ctx context.Context, in *v1.GetInstanceRequest) (*v1.GetInstanceReply, error) {
	cloudHost, err := s.cloudHostCase.Get(ctx, in.InstanceId)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	host := &v1.Host{}
	err = data.ConvertType(cloudHost, host)
	if err != nil {
		return nil, err
	}
	res := &v1.GetInstanceReply{
		Instance: host,
	}
	return res, nil
}

func (s *GreeterService) CreateInstance(ctx context.Context, in *v1.CreateInstanceRequest) (*v1.CreateInstanceReply, error) {
	host := &biz.CloudHost{
		VpcId:            in.Instance.VpcId,
		SubnetId:         in.Instance.SubnetId,
		InstanceId:       in.Instance.InstanceId,
		InstanceName:     in.Instance.InstanceName,
		InstanceState:    in.Instance.InstanceState,
		CPU:              in.Instance.Cpu,
		Memory:           in.Instance.Memory,
		CreatedTime:      in.Instance.CreatedTime,
		InstanceType:     in.Instance.InstanceType,
		EniLimit:         in.Instance.EniLimit,
		EnilpLimit:       in.Instance.EnilpLimit,
		InstanceEniCount: in.Instance.InstanceEniCount,
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
		err := data.ConvertType(v, instance)
		if err != nil {
			return nil, err
		}
		res.Instances = append(res.Instances, instance)
	}
	return res, nil
}

func (s *GreeterService) DeleteInstanceById(ctx context.Context, in *v1.DeleteInstanceRequest) (*v1.DeleteInstanceReply, error) {
	cloudHost, err := s.cloudHostCase.Delete(ctx, in.InstanceId)
	if err != nil {
		return nil, err
	}
	res := &v1.DeleteInstanceReply{
		Message: fmt.Sprintf("host %v delete success.", cloudHost.InstanceId),
	}
	return res, nil
}

func (s *GreeterService) UpdateInstance(ctx context.Context, in *v1.UpdateInstanceRequest) (*v1.UpdateInstanceReply, error) {
	var host biz.CloudHost
	err := data.ConvertType(in.Instance, &host)
	if err != nil {
		return nil, err
	}
	res := &v1.UpdateInstanceReply{}
	instance, err := s.cloudHostCase.Update(ctx, in.InstanceId, &host)
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

func (s *GreeterService) SyncFromCloudProvider(ctx context.Context, in *v1.SyncFromCloudProviderRequest) (*v1.SyncFromCloudProviderReply, error) {
	cloudProviderCase, err := data.NewCloudProviderRepo(in.CloudProvider)
	if err != nil {
		return &v1.SyncFromCloudProviderReply{
			Message: fmt.Sprintf("unknown provider"),
			Total:   0,
		}, nil
	}
	s.cloudProviderCase = biz.NewCloudProviderUsecase(cloudProviderCase, nil)

	if err := s.cloudProviderCase.GetClient(ctx, in.AccessKey, in.SecretKey, in.Region); err != nil {
		return &v1.SyncFromCloudProviderReply{
			Message: fmt.Sprintf("sync error: %s", err),
			Total:   0,
		}, nil
	}

	res, err := s.cloudProviderCase.ListInstancesByVpc(ctx, in.VpcId)
	if err != nil {
		return &v1.SyncFromCloudProviderReply{
			Message: fmt.Sprintf("sync error: %s", err),
			Total:   0,
		}, nil
	}

	count := 0
	for _, v := range res {
		_, err := s.cloudHostCase.Create(ctx, v)
		if err != nil {
			fmt.Printf("create instance error: %s\n", err)
			continue
		}
		count++
	}
	return &v1.SyncFromCloudProviderReply{
		Message: fmt.Sprintf("sync success"),
		Total:   int64(count),
	}, nil
}
