package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/rest"
)

type Cluster struct {
	Id         int
	Kubeconfig string
}

// ClusterUsecase 集群领域结构体，可以包含多个与用户业务相关的 repo
type ClusterUsecase struct {
	repo   ClusterRepo
	log    *log.Helper
	getter K8sRepoGetter
}

type ClusterRepo interface {
	Get(ctx context.Context, id int) (*Cluster, error)

	List(ctx context.Context) ([]*Cluster, error)
}

type K8sRepo interface {
	// 列出某 namespace 下所有 pod
	ListPods(ctx context.Context, namespace string) (*corev1.PodList, error)
	// 重启 pod
	RestartPod(ctx context.Context, namespace, name string) error
	// 列出集群中所有的 namespace 列表
	ListNamespace(ctx context.Context) (*corev1.NamespaceList, error)
	// 列出某个 namespace 下的deployment
	ListDeployment(ctx context.Context, namespace string) (*appsv1.DeploymentList, error)
	// 列出某 namespace 下有特定标签的pod
	ListPodsByLabelSelector(ctx context.Context, namespace string, selector labels.Selector) (*corev1.PodList, error)

	ListIngress(ctx context.Context, namespace string) (*networkingv1.IngressList, error)
	ListServiceByNamespace(ctx context.Context, namespace string) (*corev1.ServiceList, error)
	ListEvents(ctx context.Context, namespace, uid string) (*corev1.EventList, error)
	GetPodLogReq(pod, namespace string, options *corev1.PodLogOptions) *rest.Request
}
type K8sRepoGetter interface {
	GetRepo(kubeCfg []byte, help *log.Helper) (K8sRepo, error)
}

// NewClusterUsecase 集群领域构造方法
func NewClusterUsecase(repo ClusterRepo, getter K8sRepoGetter, logger log.Logger) *ClusterUsecase {
	return &ClusterUsecase{repo: repo, getter: getter, log: log.NewHelper(logger)}
}

func (c *ClusterUsecase) List(ctx context.Context) ([]*Cluster, error) {
	return c.repo.List(ctx)
}

func (c *ClusterUsecase) ListDeployments(ctx context.Context, id int, namespace string) ([]string, error) {
	repo, err := c.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	k8sRepo, err := c.getter.GetRepo([]byte(repo.Kubeconfig), c.log)
	if err != nil {
		return nil, err
	}
	deploymentList, err := k8sRepo.ListDeployment(ctx, namespace)
	if err != nil {
		return nil, err
	}
	var result []string
	for _, v := range deploymentList.Items {
		result = append(result, v.Name)
	}
	return result, nil
}

func (c *ClusterUsecase) ListNamespaces(ctx context.Context, id int) ([]string, error) {
	repo, err := c.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	k8sRepo, err := c.getter.GetRepo([]byte(repo.Kubeconfig), c.log)
	if err != nil {
		return nil, err
	}
	namespaceList, err := k8sRepo.ListNamespace(ctx)
	if err != nil {
		return nil, err
	}
	var result []string
	for _, v := range namespaceList.Items {
		result = append(result, v.Name)
	}
	return result, nil
}
