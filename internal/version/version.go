package version

import (
	"errors"
	"fmt"
	"runtime"
	"runtime/debug"
)

const (
	clientGoPath = "k8s.io/client-go"
)

// Non-allocating compile-time check for interface compliance.
var (
	_ fmt.Stringer = (*Version)(nil)
)

// Version provides a version string for the CLI.
type Version struct {
	info *debug.BuildInfo
}

// Must creates a new Version instance, panicking if any errors occur.
func Must() *Version {
	version, err := New()
	if err != nil {
		panic(err)
	}

	return version
}

// New creates a new Version instance.
func New() (*Version, error) {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return nil, errors.New("failed to read buildinfo")
	}

	return &Version{
		info: info,
	}, nil
}

// String returns the string representation of the version.
func (v *Version) String() string {
	return fmt.Sprintf(
		"%s (%s) (%s)",
		v.info.Main.Version,
		v.GoVersion(),
		v.ClientGoVersion(),
	)
}

// Version returns the CLI version.
func (v *Version) Version() string {
	return v.info.Main.Version
}

// GoVersion return the current Go version alongside the OS and architecture.
func (v *Version) GoVersion() string {
	return fmt.Sprintf("%s %s/%s", runtime.Version(), runtime.GOOS, runtime.GOARCH)
}

// ClientGoVersion returns the used version of "k8s.io/client-go".
func (v *Version) ClientGoVersion() string {
	version := "(unknown)"

	for _, dep := range v.info.Deps {
		if dep.Path == clientGoPath {
			version = dep.Version
		}
	}

	return fmt.Sprintf("%s %s", clientGoPath, version)
}
