package cli

import (
	"fmt"

	"github.com/mattdowdell/kubeswitch/internal/kube"
)

// Show outputs the current context and namespace within a kubeconfig file.
type Show struct {
	Common
}

// Help outputs the extended help for the command.
func (*Show) Help() string {
	return ""
}

// ...
func (s *Show) Run() error {
	access := kube.NewAccess(s.Config)

	conf, err := access.GetStartingConfig()
	if err != nil {
		return err
	}

	var ns string
	if ctx, ok := conf.Contexts[conf.CurrentContext]; ok {
		ns = ctx.Namespace
	}

	fmt.Println("Context:  ", conf.CurrentContext)
	fmt.Println("Namespace:", ns)

	return nil
}
