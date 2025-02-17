package cli

import (
	"context"
	"errors"
	"fmt"
	"slices"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd/api"

	"github.com/mattdowdell/kubeswitch/internal/chooser"
	"github.com/mattdowdell/kubeswitch/internal/kube"
	"github.com/mattdowdell/kubeswitch/internal/logging"
)

// ...
type Namespace struct {
	Config    string `short:"k" name:"kubeconfig" env:"KUBECONFIG" default:"${kubeconfig}" help:"The kubeconfig file to use (env: ${env})."`
	Namespace string `arg:"" optional:""  help:"The namespace to switch to."`
	Verbose   int    `short:"v" type:"counter" help:"Increase the log verbosity."`
}

// ...
func (n *Namespace) Run(ctx context.Context) error {
	logger := logging.New(n.Verbose)

	conf, err := kube.LoadConfig(n.Config)
	if err != nil {
		return err
	}

	updated, err := updateNamespace(ctx, logger, conf, n.Namespace)
	if err != nil {
		return err
	}

	if !updated {
		logger.Info("skipped namespace update")
		return nil
	}

	if err := kube.SaveConfig(conf, n.Config); err != nil {
		return err
	}

	logger.Info("updated with new namespace", "name", conf.Contexts[conf.CurrentContext].Namespace)
	return nil
}

func updateNamespace(ctx context.Context, logger *log.Logger, conf *api.Config, val string) (bool, error) {
	client, err := kube.ClientFromConfig(conf)
	if err != nil {
		return false, err
	}

	namespaces, err := client.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		return false, err
	}

	names := make([]string, 0, len(namespaces.Items))
	for i := range namespaces.Items {
		names = append(names, namespaces.Items[i].Name)
	}
	slices.Sort(names)

	var name string

	if val != "" {
		for _, n := range names {
			if n == val {
				name = n
				break
			}
		}

		if name == "" {
			return false, fmt.Errorf("unknown namespace: %s", val)
		}
	} else {
		name, err = chooseNamespace(logger, names, conf.Contexts[conf.CurrentContext].Namespace)
		if err != nil {
			return false, err
		}
	}

	if name == conf.Contexts[conf.CurrentContext].Namespace {
		return false, nil
	}

	conf.Contexts[conf.CurrentContext].Namespace = name
	return true, nil
}

func chooseNamespace(logger *log.Logger, names []string, current string) (string, error) {
	switch len(names) {
	case 0:
		return "", errors.New("no namespace to choose from")

	case 1:
		logger.Info("defaulting to the one available namespace", "name", names[0])
		return names[0], nil

	default:
		c := chooser.New("Select a namespace:", names, current)

		if _, err := tea.NewProgram(c, tea.WithAltScreen()).Run(); err != nil {
			return "", fmt.Errorf("error running program: %w", err)
		}

		choice, ok := c.Choice()
		if !ok {
			return "", errors.New("no namespace chosen")
		}

		return choice, nil
	}
}
