package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"runtime/debug"

	"github.com/alecthomas/kong"
)

const (
	clientGoPath = "k8s.io/client-go"
)

func main() {
	ctx := kong.Parse(
		&CLI{},
		kong.Description("TODO"),
		kong.Vars{
			"kubeconfig": filepath.Join(os.Getenv("HOME"), ".kube", "config"),
			"version":    version(), // TODO: k8s client version
		},
	)
	ctx.FatalIfErrorf(ctx.Run())
}

func version() string {
	info, _ := debug.ReadBuildInfo()
	return fmt.Sprintf("%s (%s) (%s)", info.Main.Version, goVersion(), clientGoVersion(info))
}

func goVersion() string {
	return fmt.Sprintf("%s %s/%s", runtime.Version(), runtime.GOOS, runtime.GOARCH)
}

func clientGoVersion(info *debug.BuildInfo) string {
	version := "(unknown)"

	for _, dep := range info.Deps {
		if dep.Path == clientGoPath {
			version = dep.Version
		}
	}

	return fmt.Sprintf("%s %s", clientGoPath, version)
}
