package kube

import (
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

// ...
func LoadConfig(path string) (*api.Config, error) {
	return clientcmd.LoadFromFile(path)
}

// ...
func LoadConfigAlt(path string) (clientcmd.ClientConfig) {
	rules := clientcmd.NewDefaultClientConfigLoadingRules()
	rules.ExplicitPath = path

	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		rules,
		&clientcmd.ConfigOverrides{},
	)
}

// ...
func SaveConfig(conf *api.Config, path string) error {
	return clientcmd.WriteToFile(*conf, path)
}


// ...
func ConfigAccess(path string) clientcmd.ConfigAccess {
	return &clientcmd.PathOptions{
		ExplicitFileFlag: path,
	}
}

func UpdateConfig(access clientcmd.ConfigAccess, conf *api.Config) error {
	return clientcmd.ModifyConfig(access, *conf, true /*relativizePaths*/)
}
