package main

import "github.com/wilbeibi/cobra-demo/cmd"

/*
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
*/

func init() {
	// log initializing
}
func main() {
	cmd.Execute()
}
