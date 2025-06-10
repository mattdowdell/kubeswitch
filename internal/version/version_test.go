package version_test

import (
	"runtime/debug"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mattdowdell/kubeswitch/internal/version"
)

func Test_Must(t *testing.T) {
	assert.NotPanics(t, func() {
		v := version.Must()
		assert.NotNil(t, v)
	})
}

func Test_New(t *testing.T) {
	// arrange

	// act
	v, err := version.New()

	// assert
	assert.NotNil(t, v)
	assert.NoError(t, err)
}

func Test_NewFromInfo(t *testing.T) {
	// arrange
	info := &debug.BuildInfo{}

	// act
	v := version.NewFromInfo(info)

	// assert
	assert.NotNil(t, v)
}

func Test_Version_String(t *testing.T) {
	// arrange
	info := &debug.BuildInfo{
		Main: debug.Module{
			Version: "v0.0.1",
		},
		Deps: []*debug.Module{
			{
				Path:    "k8s.io/client-go",
				Version: "v0.32.1",
			},
		},
	}

	v := version.NewFromInfo(info)

	// act
	got := v.String()

	// assert
	assert.Regexp(
		t,
		`^v0\.0\.1 \(go1\.\d+.\d+ \w+\/\w+\) \(k8s\.io\/client-go v0\.32\.1\)$`,
		got,
	)
}

func Test_Version_Version(t *testing.T) {
	// arrange
	info := &debug.BuildInfo{
		Main: debug.Module{
			Version: "v0.0.1",
		},
	}

	v := version.NewFromInfo(info)

	// act
	got := v.Version()

	// assert
	assert.Equal(t, "v0.0.1", got)
}

func Test_Version_GoVersion(t *testing.T) {
	// arrange
	info := &debug.BuildInfo{}
	v := version.NewFromInfo(info)

	// act
	got := v.GoVersion()

	// assert
	assert.Regexp(t, `^go1\.\d+\.\d+ \w+\/\w+$`, got)
}

func Test_Version_ClientGoVersion(t *testing.T) {
	tests := map[string]struct {
		have *debug.BuildInfo
		want string
	}{
		"missing": {
			have: &debug.BuildInfo{},
			want: "k8s.io/client-go (unknown)",
		},
		"present": {
			have: &debug.BuildInfo{
				Deps: []*debug.Module{
					{
						Path:    "k8s.io/client-go",
						Version: "v0.32.1",
					},
				},
			},
			want: "k8s.io/client-go v0.32.1",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			// arrange
			v := version.NewFromInfo(tt.have)

			// act
			got := v.ClientGoVersion()

			// assert
			assert.Equal(t, tt.want, got)
		})
	}
}
