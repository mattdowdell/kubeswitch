package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/alecthomas/kong"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/mattdowdell/kubeswitch/internal/cli"
	"github.com/mattdowdell/kubeswitch/internal/version"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	k := kong.Parse(
		&cli.CLI{},
		kong.Description("Switch between kube contexts and namespaces."),
		kong.Vars{
			"kubeconfig": clientcmd.RecommendedHomeFile,
			"version":    version.Must().String(),
		},
		kong.BindTo(ctx, (*context.Context)(nil)),
	)

	k.FatalIfErrorf(k.Run())
}
