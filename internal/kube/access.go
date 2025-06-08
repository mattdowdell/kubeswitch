package kube

import (
	"os"
	"path/filepath"
	"slices"

	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

// ...
var _ clientcmd.ConfigAccess = (*Access)(nil)

// ...
//
// TODO: discuss why this exists, i.e. to avoid hacks with env vars,
// support multiple files through CLI options, etc.
type Access struct {
	Files []string
}

// ...
func NewAccess(path string) *Access {
	return &Access{
		Files: dedupe(filepath.SplitList(path)),
	}
}

// ...
func (a *Access) GetDefaultFilename() string {
	for _, path := range a.Files {
		if _, err := os.Stat(path); err != nil {
			return path
		}
	}

	return clientcmd.RecommendedHomeFile
}

// ...
func (a *Access) GetLoadingPrecedence() []string {
	return slices.Clone(a.Files)
}

// ...
func (a *Access) GetStartingConfig() (*api.Config, error) {
	rules := &clientcmd.ClientConfigLoadingRules{
		Precedence: a.GetLoadingPrecedence(),
	}

	conf := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		rules,
		&clientcmd.ConfigOverrides{},
	)

	raw, err := conf.RawConfig()
	if err != nil {
		return nil, err
	}

	return &raw, nil
}

// ...
func (a *Access) IsExplicitFile() bool {
	return false
}

// ...
func (a *Access) GetExplicitFile() string {
	return ""
}

// ...
func dedupe(inputs []string) []string {
	seen := map[string]struct{}{}
	var output []string

	for _, input := range inputs {
		if _, ok := seen[input]; ok {
			continue
		}

		output = append(output, input)
		seen[input] = struct{}{}
	}

	return output
}
