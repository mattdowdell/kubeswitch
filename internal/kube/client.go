package kube

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

// ClientFromConfig creates an API client from the given configuration.
func ClientFromConfig(conf *api.Config) (*kubernetes.Clientset, error) {
	clientConf := clientcmd.NewDefaultClientConfig(*conf, &clientcmd.ConfigOverrides{})

	restConf, err := clientConf.ClientConfig()
	if err != nil {
		return nil, err
	}

	return kubernetes.NewForConfig(restConf)
}
