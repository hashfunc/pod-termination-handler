package kube

import (
	"fmt"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func NewClient(config *rest.Config) (*kubernetes.Clientset, error) {
	kubeClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("NewClient: %w", err)
	}

	return kubeClient, nil
}
