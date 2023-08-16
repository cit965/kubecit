package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"kubecit/internal/biz"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

type k8sRepo struct {
	clientSet *kubernetes.Clientset
}

type k8sRepoGetter struct{}

func NewK8sRepoGetter() biz.K8sRepoGetter {
	return &k8sRepoGetter{}
}

func (k *k8sRepoGetter) GetRepo(kubeCfg []byte, help *log.Helper) (biz.K8sRepo, error) {
	return NewK8sRepo(kubeCfg, help)
}

func (k *k8sRepo) ListNamespace(ctx context.Context) (*corev1.NamespaceList, error) {
	return k.clientSet.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
}
func (k *k8sRepo) ListDeployment(ctx context.Context, namespace string) (*appsv1.DeploymentList, error) {
	return k.clientSet.AppsV1().Deployments(namespace).List(ctx, metav1.ListOptions{})
}

func (k *k8sRepo) ListPods(ctx context.Context, namespace string) (*corev1.PodList, error) {
	return k.clientSet.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{})
}

func (k *k8sRepo) RestartPod(ctx context.Context, namespace, name string) error {
	return k.clientSet.CoreV1().Pods(namespace).Delete(ctx, name, metav1.DeleteOptions{})
}

func (k *k8sRepo) ListPodsByLabelSelector(ctx context.Context, namespace string, selector labels.Selector) (*corev1.PodList, error) {
	podList, err := k.clientSet.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{
		LabelSelector: selector.String(),
	})
	if err != nil {
		return nil, err
	}
	return podList, nil
}

// ListIngress 获取 ingress
func (k *k8sRepo) ListIngress(ctx context.Context, namespace string) (*networkingv1.IngressList, error) {
	return k.clientSet.NetworkingV1().Ingresses(namespace).List(ctx, metav1.ListOptions{})
}

// ListServiceByNamespace 获取  services
func (k *k8sRepo) ListServiceByNamespace(ctx context.Context, namespace string) (*corev1.ServiceList, error) {
	return k.clientSet.CoreV1().Services(namespace).List(ctx, metav1.ListOptions{})
}

// ListEvents 获取事件
func (k *k8sRepo) ListEvents(ctx context.Context, namespace, uid string) (*corev1.EventList, error) {
	fieldSelector := fields.OneTermEqualSelector("involvedObject.uid", uid).String()
	return k.clientSet.CoreV1().Events(namespace).List(ctx, metav1.ListOptions{FieldSelector: fieldSelector})
}

// GetPodLogReq 获取 pod 日志
func (k *k8sRepo) GetPodLogReq(pod, namespace string, options *corev1.PodLogOptions) *rest.Request {
	return k.clientSet.CoreV1().Pods(namespace).GetLogs(pod, options)
}

func NewK8sRepo(kubeCfg []byte, logger *log.Helper) (biz.K8sRepo, error) {
	config, err := clientcmd.RESTConfigFromKubeConfig(kubeCfg)
	if err != nil {
		logger.Errorf("RESTConfigFromKubeConfig err: %s", err)
		return nil, err
	}
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		logger.Errorf("NewForConfig err:%s", err)
		return nil, err
	}
	return &k8sRepo{clientSet: clientSet}, nil
}
