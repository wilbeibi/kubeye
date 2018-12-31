package client

import (
	"os"

	"github.com/pkg/errors"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
)

// Create an out of cluster client
// Use `kubernetes.Interface` in return for fake testing, refer:
// https://itnext.io/testing-kubernetes-go-applications-f1f87502b6ef
func NewOutOfClusterClient() (kubernetes.Interface, error) {
	kubeconfig := os.Getenv("KUBECONFIG")
	if kubeconfig == "" {
		kubeconfig = os.Getenv("HOME") + "/.kube/config"
	}
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, errors.Wrap(err, "getting kubeconfig for client")
	}
	return kubernetes.NewForConfig(config)
}
