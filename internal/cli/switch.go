package cli

import (
	"context"

	"github.com/mattdowdell/kubeswitch/internal/kube"
	"github.com/mattdowdell/kubeswitch/internal/logging"
)

// Switch provides the ability to switch the current context and namespace within a kube config
// file.
type Switch struct {
	Common

	Context   string `short:"c" help:"The context to switch to."`
	Namespace string `short:"n" help:"The namespace to switch to."`
}

// Help outputs the extended help for the command.
func (*Switch) Help() string {
	return `A new context and namespace can either be selected interactively from the available
choices, or using pre-selected values.

This command lists namespaces from the Kubernetes API for interactive selection and validation. As a
result, a valid kubeconfig with access to namespaces is required.

Examples:
	# Pre-select all values
	kubeswitch sw -c CONTEXT -n NAMESPACE
	kubeswitch switch --context CONTEXT --namespace NAMESPACE

	# Interactively select all values
	kubeswitch sw
	kubeswitch switch

	# Mixture of pre-selected and interactive selected values
	kubeswitch switch --context CONTEXT
	kubeswitch switch --namespace NAMESPACE`
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
