## kp custom-builder list

List available custom builders

### Synopsis

Prints a table of the most important information about the available custom builders in the provided namespace.

namespace defaults to the kubernetes current-context namespace.

```
kp custom-builder list [flags]
```

### Examples

```
kp cb list
kp cb list -n my-namespace
```

### Options

```
  -h, --help               help for list
  -n, --namespace string   kubernetes namespace
```

### SEE ALSO

* [kp custom-builder](kp_custom-builder.md)	 - Custom Builder Commands

###### Auto generated by spf13/cobra on 30-Jul-2020
