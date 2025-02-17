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
func SaveConfig(conf *api.Config, path string) error {
	return clientcmd.WriteToFile(*conf, path)
}
