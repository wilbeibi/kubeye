package main

import (
	"os"

	"github.com/pkg/errors"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// buildOutOfClusterConfig https://github.com/bitnami-labs/kubewatch/blob/master/pkg/utils/k8sutil.go
func getOutOfClusterConfig() (kubernetes.Interface, error) {
	kubeconfigPath := os.Getenv("KUBECONFIG")
	if kubeconfigPath == "" {
		kubeconfigPath = os.Getenv("HOME") + "/.kube/config"
	}
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		return nil, errors.Wrap(err, "Cannot get kubernetes config")
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, errors.Wrap(err, "Cannot create kubernetes clientset")
	}
	return clientset, nil
}

func run() error {
	client, err := getOutOfClusterConfig()
	if err != nil {
		return errors.Wrap(err, "")
	}
	return nil
}

func main() {

}
