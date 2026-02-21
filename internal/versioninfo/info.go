package versioninfo

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
var _ fmt.Stringer = (*VersionInfo)(nil)

// VersionInfo provides a version string for the CLI.
type VersionInfo struct {
	info *debug.BuildInfo
}

// Must creates a new VersionInfo instance, panicking if any errors occur.
func Must() *VersionInfo {
	version, err := New()
	if err != nil {
		panic(err)
	}

	return version
}

// New creates a new VersionInfo instance.
func New() (*VersionInfo, error) {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return nil, errors.New("failed to read buildinfo")
	}

	return NewFromBuildInfo(info), nil
}

// NewFromBuildInfo creates a new VersionInfo instance from the given [runtime/debug.BuildInfo].
func NewFromBuildInfo(info *debug.BuildInfo) *VersionInfo {
	return &VersionInfo{
		info: info,
	}
}

// String returns the string representation of the version.
func (v *VersionInfo) String() string {
	return fmt.Sprintf(
		"%s (%s) (%s)",
		v.Version(),
		v.GoVersion(),
		v.ClientGoVersion(),
	)
}

// Version returns the CLI version.
func (v *VersionInfo) Version() string {
	return v.info.Main.Version
}

// GoVersion return the current Go version alongside the OS and architecture.
func (v *VersionInfo) GoVersion() string {
	return fmt.Sprintf("%s %s/%s", runtime.Version(), runtime.GOOS, runtime.GOARCH)
}

// ClientGoVersion returns the used version of "k8s.io/client-go".
func (v *VersionInfo) ClientGoVersion() string {
	version := "(unknown)"

	for _, dep := range v.info.Deps {
		if dep.Path == clientGoPath {
			version = dep.Version
		}
	}

	return fmt.Sprintf("%s %s", clientGoPath, version)
}
