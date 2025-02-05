package main

import (
	"os"
	"path/filepath"

	"github.com/alecthomas/kong"
)

func main() {
	ctx := kong.Parse(
		&CLI{},
		kong.Description("TODO"),
		kong.Vars{
			"kubeconfig": filepath.Join(os.Getenv("HOME"), ".kube", "config"),
			"version": "TODO", // CLI + k8s client version
		},
	)
	ctx.FatalIfErrorf(ctx.Run())
}

func envOrDefault(name, fallback string) string {
	if val, ok := os.LookupEnv(name); ok {
		return val
	}

	return fallback
}
