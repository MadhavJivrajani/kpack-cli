## kp custom-builder status

Display status of a custom builder

### Synopsis

Prints detailed information about the status of a specific custom builder in the provided namespace.

namespace defaults to the kubernetes current-context namespace.

```
kp custom-builder status <name> [flags]
```

### Examples

```
kp cb status my-builder
kp cb status -n my-namespace other-builder
```

### Options

```
  -h, --help               help for status
  -n, --namespace string   kubernetes namespace
```

### SEE ALSO

* [kp custom-builder](kp_custom-builder.md)	 - Custom Builder Commands

###### Auto generated by spf13/cobra on 30-Jul-2020
