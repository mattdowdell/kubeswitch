package cli

import (
	"context"

	"github.com/mattdowdell/kubeswitch/internal/kube"
	"github.com/mattdowdell/kubeswitch/internal/logging"
)

// ...
type Switch struct {
	Config    string `short:"k" name:"kubeconfig" env:"KUBECONFIG" default:"${kubeconfig}" help:"The kubeconfig file to use (env: ${env})."`
	Context   string `short:"c" help:"The context to switch to."`
	Namespace string `short:"n" help:"The namespace to switch to."`
	Verbose   int    `short:"v" type:"counter" help:"Increase the log verbosity."`
}

// ...
func (*Switch) Help() string {
	return ""
}

// ...
func (s *Switch) Run(ctx context.Context) error {
	logger := logging.New(s.Verbose)
	access := kube.NewAccess(s.Config)

	conf, err := access.GetStartingConfig()
	if err != nil {
		return err
	}

	ctxUpdated, err := updateContext(logger, conf, s.Context)
	if err != nil {
		return err
	}

	nsUpdated, err := updateNamespace(ctx, logger, conf, s.Namespace)
	if err != nil {
		return err
	}

	if !ctxUpdated && !nsUpdated {
		logger.Info("skipped update")
		return nil
	}

	if err := kube.UpdateConfig(access, conf); err != nil {
		return err
	}

	logger.Info(
		"updated with new values",
		"context", conf.CurrentContext,
		"namespace", conf.Contexts[conf.CurrentContext].Namespace,
	)

	return nil
}
