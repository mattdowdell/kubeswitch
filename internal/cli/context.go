package cli

import (
	"errors"
	"fmt"
	"slices"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"k8s.io/client-go/tools/clientcmd/api"

	"github.com/mattdowdell/kubeswitch/internal/chooser"
	"github.com/mattdowdell/kubeswitch/internal/kube"
	"github.com/mattdowdell/kubeswitch/internal/logging"
)

// ...
type Context struct {
	Config  string `short:"k" name:"kubeconfig" env:"KUBECONFIG" default:"${kubeconfig}" help:"The kubeconfig file to use (env: ${env})."`
	Context string `arg:"" optional:"" help:"The context to switch to."`
	Verbose int    `short:"v" type:"counter" help:"Increase the log verbosity."`
}

// ...
func (*Context) Help() string {
	return ""
}

// ...
func (c *Context) Run() error {
	logger := logging.New(c.Verbose)
	access := kube.ConfigAccess(c.Config)

	conf, err := access.GetStartingConfig()
	if err != nil {
		return err
	}

	updated, err := updateContext(logger, conf, c.Context)
	if err != nil {
		return err
	}

	if !updated {
		logger.Info("skipped context update")
		return nil
	}

	if err := kube.UpdateConfig(access, conf); err != nil {
		return err
	}

	logger.Info("updated with new context", "name", conf.CurrentContext)
	return nil
}

func updateContext(logger *log.Logger, conf *api.Config, val string) (bool, error) {
	names := make([]string, 0, len(conf.Contexts))
	for name := range conf.Contexts {
		names = append(names, name)
	}

	slices.Sort(names)

	name := val
	if val == "" {
		n, err := chooseContext(logger, names, conf.CurrentContext)
		if err != nil {
			return false, err
		}

		name = n
	}

	if _, ok := conf.Contexts[name]; !ok {
		return false, fmt.Errorf("unknown context: %q", name)
	}

	if conf.CurrentContext == name {
		// nothing to do
		return false, nil
	}

	conf.CurrentContext = name
	return true, nil
}

func chooseContext(logger *log.Logger, names []string, current string) (string, error) {
	switch len(names) {
	case 0:
		return "", errors.New("no contexts to choose from")

	case 1:
		logger.Info("defaulting to the one available context", "name", names[0])
		return names[0], nil

	default:
		c := chooser.New("Select a context:", names, current)

		if _, err := tea.NewProgram(c, tea.WithAltScreen()).Run(); err != nil {
			return "", fmt.Errorf("error running program: %w", err)
		}

		choice, ok := c.Choice()
		if !ok {
			return "", errors.New("no context chosen")
		}

		return choice, nil
	}
}
