package cli

import (
	"github.com/alecthomas/kong"
)

// ...
type CLI struct {
	Switch    Switch    `cmd:"" aliases:"sw" default:"1" help:"Switch all values."`
	Context   Context   `cmd:"" aliases:"ctx" help:"Switch the context only."`
	Namespace Namespace `cmd:"" aliases:"ns" help:"Switch the namespace only."`

	Version kong.VersionFlag `short:"V" help:"Display the version."`
}
