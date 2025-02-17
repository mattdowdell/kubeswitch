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

// ...
var (
	_ fmt.Stringer = (*Version)(nil)
)

// ...
type Version struct {
	info *debug.BuildInfo
}

// ...
func Must() *Version {
	version, err := New()
	if err != nil {
		panic(err)
	}

	return version
}

// ...
func New() (*Version, error) {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return nil, errors.New("failed to read buildinfo")
	}

	return &Version{
		info: info,
	}, nil
}

// ...
func (v *Version) String() string {
	return fmt.Sprintf(
		"%s (%s) (%s)",
		v.info.Main.Version,
		v.GoVersion(),
		v.ClientGoVersion(),
	)
}

// ...
func (v *Version) Version() string {
	return v.info.Main.Version
}

// ...
func (v *Version) GoVersion() string {
	return fmt.Sprintf("%s %s/%s", runtime.Version(), runtime.GOOS, runtime.GOARCH)
}

// ...
func (v *Version) ClientGoVersion() string {
	version := "(unknown)"

	for _, dep := range v.info.Deps {
		if dep.Path == clientGoPath {
			version = dep.Version
		}
	}

	return fmt.Sprintf("%s %s", clientGoPath, version)
}
