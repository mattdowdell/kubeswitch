package kube

import (
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

// ...
func ConfigAccess(path string) clientcmd.ConfigAccess {
	rules := clientcmd.NewDefaultClientConfigLoadingRules()
	rules.ExplicitPath = path

	return &clientcmd.PathOptions{
		LoadingRules: rules,
	}
}

// ...
func UpdateConfig(access clientcmd.ConfigAccess, conf *api.Config) error {
	return clientcmd.ModifyConfig(access, *conf, true /*relativizePaths*/)
}
