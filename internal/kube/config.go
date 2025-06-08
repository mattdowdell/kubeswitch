package kube

import (
	"fmt"

	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

// ...
func UpdateConfig(access clientcmd.ConfigAccess, conf *api.Config) error {
	if err := clientcmd.ModifyConfig(access, *conf, true /*relativizePaths*/); err != nil {
		return fmt.Errorf("failed to persist update: %w", err)
	}

	return nil
}
