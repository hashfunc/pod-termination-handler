package kube

import (
	E "github.com/IBM/fp-go/either"
	F "github.com/IBM/fp-go/function"
	"k8s.io/client-go/kubernetes"
)

func NewClient() E.Either[error, *kubernetes.Clientset] {
	return F.Pipe1(
		getConfig(),
		E.Chain(E.Eitherize1(kubernetes.NewForConfig)),
	)
}
