package cli

import (
	"github.com/alecthomas/kong"
)

// ...
type CLI struct {
	Show      Show      `cmd:"" aliases:"sh" default:"1" help:"Show current values."`
	Switch    Switch    `cmd:"" aliases:"sw"  help:"Switch all values."`
	Context   Context   `cmd:"" aliases:"ctx" help:"Switch the context only."`
	Namespace Namespace `cmd:"" aliases:"ns" help:"Switch the namespace only."`

	Version kong.VersionFlag `short:"V" help:"Display the version."`
}

// Common contains common options for CLI subcommands.
type Common struct {
	Config  string `short:"k" name:"kubeconfig" env:"KUBECONFIG" default:"${kubeconfig}" help:"The kubeconfig file to use (env: ${env})."`
	Verbose int    `short:"v" type:"counter" help:"Increase the log verbosity."`
}
