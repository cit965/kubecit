package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"kubecit/internal/biz"

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
