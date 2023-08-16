package service

import (
	"context"
	v1 "kubecit/api/helloworld/v1"
	"kubecit/internal/biz"
)

type ClusterService struct {
	v1.UnimplementedClusterServer

	cluster *biz.ClusterUseCase
}

func NewClusterService(cluster *biz.ClusterUseCase) *ClusterService {
	return &ClusterService{cluster: cluster}
}

// ClusterRegister register a cluster with name and ip
func (s *ClusterService) ClusterRegister(ctx context.Context, in *v1.ClusterRegisterRequest) (*v1.ClusterRegisterResponse, error) {
	result, err := s.cluster.RegisterCluster(ctx, &biz.Cluster{
		Alias:      in.Alias,
		Kubeconfig: in.Kubeconfig,
	})
	if err != nil {
		return nil, err
	}
	return &v1.ClusterRegisterResponse{Cluster: &v1.ClusterBase{
		Id:         int64(result.ID),
		Alias:      result.Alias,
		Kubeconfig: result.Kubeconfig,
	}}, nil
}

func (s *ClusterService) ClusterList(ctx context.Context, in *v1.ClusterListRequest) (reply *v1.ClusterListResponse, err error) {
	clusters, err := s.cluster.ListCluster(ctx)
	if err != nil {
		return nil, err
	}
	var clusterRes []*v1.ClusterBase

	for _, v := range clusters {
		clusterRes = append(clusterRes, &v1.ClusterBase{
			Id:         int64(v.ID),
			Alias:      v.Alias,
			Kubeconfig: v.Kubeconfig,
		})
	}
	return &v1.ClusterListResponse{Clusters: clusterRes}, nil
}

func (s *ClusterService) ClusterGet(ctx context.Context, in *v1.ClusterBase) (reply *v1.ClusterBase, err error) {
	cluster, err := s.cluster.GetCluster(ctx, int(in.Id))
	if err != nil {
		return nil, err
	}
	return &v1.ClusterBase{
		Id:         int64(cluster.ID),
		Alias:      cluster.Alias,
		Kubeconfig: cluster.Kubeconfig,
	}, nil
}

func (s *ClusterService) ClusterUpdate(ctx context.Context, in *v1.ClusterBase) (reply *v1.ClusterBase, err error) {
	cluster, err := s.cluster.UpdateCluster(ctx, &biz.Cluster{
		ID:         int(in.Id),
		Alias:      in.Alias,
		Kubeconfig: in.Kubeconfig,
	})
	if err != nil {
		return nil, err
	}
	return &v1.ClusterBase{
		Id:         int64(cluster.ID),
		Alias:      cluster.Alias,
		Kubeconfig: cluster.Kubeconfig,
	}, nil
}

func (s *ClusterService) ClusterDelete(ctx context.Context, in *v1.ClusterBase) (reply *v1.ClusterDeleteResponse, err error) {
	err = s.cluster.Delete(ctx, int(in.Id))
	if err != nil {
		return nil, err
	}
	return &v1.ClusterDeleteResponse{}, nil
}
