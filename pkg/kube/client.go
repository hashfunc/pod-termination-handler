package kube

import (
	E "github.com/IBM/fp-go/either"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func NewClient(config *rest.Config) E.Either[error, *kubernetes.Clientset] {
	return E.Eitherize1(kubernetes.NewForConfig)(config)
}
