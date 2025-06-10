package kube

import (
	"os"
	"path/filepath"
	"slices"

	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

// Non-allocating compile-time check for interface compliance.
var _ clientcmd.ConfigAccess = (*Access)(nil)

// Access provides access to the kube config files that should be used to configure the client.
//
// This is loosely based on [clientcmd.PathOptions] with the following differences:
//
//   - Supports multiple files for both environment variables and CLI options.
//   - Does not support an explicit file.
//
// Notably, clientcmd.PathOptions recomputes the loading precedence from environment variables each
// time it is desired instead of respecting any explicitly set order. This means that multiple files
// are only supported via an environment variable and never via an explicit file profvided by a CLI
// option.
//
// [clientcmd.PathOptions]: https://pkg.go.dev/k8s.io/client-go/tools/clientcmd#PathOptions
type Access struct {
	paths []string
}

// NewAccess creates a new access instance. The given path is split using [os.PathListSeparator] and
// any duplicates are removed.
func NewAccess(path string) *Access {
	return &Access{
		paths: dedupe(filepath.SplitList(path)),
	}
}

// GetDefaultFilename returns the first path that exists for the given kube config files.
func (a *Access) GetDefaultFilename() string {
	for _, path := range a.paths {
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}

	return ""
}

// GetLoadingPrecedence returns the paths to kube config files in the order given.
func (a *Access) GetLoadingPrecedence() []string {
	return slices.Clone(a.paths)
}

// GetStartingConfig returns an kube API config from a merge of the kube config files.
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

// IsExplicitFile tests whether an explicit file is provided. It always returns false.
func (a *Access) IsExplicitFile() bool {
	return false
}

// GetExplicitFile returns the explicit file. It always returns an empty string.
func (a *Access) GetExplicitFile() string {
	return ""
}

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
