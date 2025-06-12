package functional_test

import (
	"testing"

	"github.com/rogpeppe/go-internal/testscript"
)

func Test_All(t *testing.T) {
	testscript.Run(t, testscript.Params{
		Dir: "testdata",
	})
}
