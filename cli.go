package main

import (
	"context"
	"errors"
	"fmt"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
	"github.com/alecthomas/kong"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ...
type CLI struct {
	Switch Switch `cmd:"" aliases:"sw" default:"1" help:"Switch all values."`
	Context Context `cmd:"" aliases:"ctx" help:"Switch the context only."`
	Namespace Namespace `cmd:"" aliases:"ns" help:"Switch the namespace only."`

	Version kong.VersionFlag `short:"V" help:"Display the version."`
}

type Switch struct {
	Config    string `short:"k" name:"kubeconfig" env:"KUBECONFIG" default:"${kubeconfig}" help:"The kubeconfig file to use (env: ${env})."`
	Context   string `short:"c" help:"The context to switch to."`
	Namespace string `short:"n" help:"The namespace to switch to."`
}

func (s *Switch) Run() error {
	conf, err := parseConfig(s.Config)
	if err != nil {
		return err
	}

	updated, err := updateContext(conf, s.Context)
	if err != nil {
		return err
	}

	if updated {
		// TODO: write changes to file
		fmt.Println("context updated to:", conf.CurrentContext)
	}

	return nil
}

type Context struct {
	Config    string `short:"k" name:"kubeconfig" env:"KUBECONFIG" default:"${kubeconfig}" help:"The kubeconfig file to use (env: ${env})."`
	Context   string `arg:"" optional:"" help:"The context to switch to."`
}

func (c *Context) Run() error {
	conf, err := parseConfig(c.Config)
	if err != nil {
		return err
	}

	updated, err := updateContext(conf, c.Context)
	if err != nil {
		return err
	}

	if updated {
		// TODO: write changes to file
		fmt.Println("context updated to:", conf.CurrentContext)
	}

	return nil
}

type Namespace struct {
	Config    string `short:"k" name:"kubeconfig" env:"KUBECONFIG" default:"${kubeconfig}" help:"The kubeconfig file to use (env: ${env})."`
	Namespace string `arg:"" optional:""  help:"The namespace to switch to."`
}

func (n *Namespace) Run() error {
	conf, err := parseConfig(n.Config)
	if err != nil {
		return err
	}

	clientConf := clientcmd.NewDefaultClientConfig(*conf, &clientcmd.ConfigOverrides{})
	restConf, err := clientConf.ClientConfig()
	if err != nil {
		return err
	}

	clientset, err := kubernetes.NewForConfig(restConf)
	if err != nil {
		return err
	}

	namespaces, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return err
	}

	fmt.Println(namespaces)

	return nil
}

func parseConfig(path string) (*api.Config, error) {
	return clientcmd.LoadFromFile(path)
}

func updateContext(conf *api.Config, val string) (bool, error) {
	names := make([]string, 0, len(conf.Contexts))
	for name := range conf.Contexts {
		names = append(names, name)
	}

	name := val
	if val == "" {
		n, err := chooseContext(names)
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

func chooseContext(names []string) (string, error) {
	switch len(names) {
	case 0:
		return "", errors.New("no contexts to choose from")

	case 1:
		return names[0], nil

	default:
		// TODO: implement choice with bubbletea
		return "", errors.New("not implemented")
	}
}
