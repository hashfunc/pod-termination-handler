package kube

import (
	E "github.com/IBM/fp-go/either"
	F "github.com/IBM/fp-go/function"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
)

func GetConfig() E.Either[error, *rest.Config] {
	return F.Pipe1(
		getInclusterConfig(),
		E.OrElse(func(err error) E.Either[error, *rest.Config] {
			klog.Info("use local config: %w", err)
			return getLocalConfig()
		}),
	)
}

func getInclusterConfig() E.Either[error, *rest.Config] {
	return E.Eitherize0(rest.InClusterConfig)()
}

func getLocalConfig() E.Either[error, *rest.Config] {
	return E.Eitherize2(clientcmd.BuildConfigFromFlags)("", clientcmd.RecommendedHomeFile)
}
