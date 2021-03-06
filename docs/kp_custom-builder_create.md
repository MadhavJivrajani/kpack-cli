## kp custom-builder create

Create a custom builder

### Synopsis

Create a custom builder by providing command line arguments.
This custom builder will be created only if it does not exist in the provided namespace.

namespace defaults to the kubernetes current-context namespace.

```
kp custom-builder create <name> --tag <tag> [flags]
```

### Examples

```
kp cb create my-builder --tag my-registry.com/my-builder-tag --order /path/to/order.yaml --stack tiny --store my-store
kp cb create my-builder --tag my-registry.com/my-builder-tag --order /path/to/order.yaml
```

### Options

```
  -h, --help               help for create
  -n, --namespace string   kubernetes namespace
  -o, --order string       path to buildpack order yaml
  -s, --stack string       stack resource to use (default "default")
      --store string       buildpack store to use (default "default")
  -t, --tag string         registry location where the builder will be created
```

### SEE ALSO

* [kp custom-builder](kp_custom-builder.md)	 - Custom Builder Commands

###### Auto generated by spf13/cobra on 30-Jul-2020
