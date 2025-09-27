# kubeswitch

_TODO: introduction._

## Install

Kubeswitch can be installed via `go install`:

```sh
go install github.com/mattdowdell/kubeswitch@latest
```

<!--
Pre-built binaries for various OS/arch combinations can be downloaded from
[Releases].

[releases]: https://github.com/mattdowdell/kubeswitch/releases
-->

## Usage

A context and namespace can be selected interactively. Contexts are taken a
kubeconfig file, e.g. `~/.kube/config.yaml`, while namespace are queried from
the Kubernetes API using the selected context.

```sh
# print current context and namespace
kubeswitch
kubeswitch sh
kubeswitch show

# interactively select a context and namespace
kubeswitch sw
kubeswitch switch

# interactively select a context
kubeswitch ctx
kubeswitch context

# interactively select a namespace for the current context
kubeswitch ns
kubeswitch namespace
```

Alternatively, a known value can be selected via options/arguments:

```sh
# select a known context and namespace
# if either are missing, it will be prompted for interactively
kubeswitch sw -c CONTEXT -n NAMESPACE
kubeswitch switch --context CONTEXT --namespace NAMESPACE

# select a known context
kubeswitch ctx CONTEXT
kubeswitch context CONTEXT

# select a known namespace
kubeswitch ns NAMESPACE
kubeswitch namespace NAMESPACE
```


