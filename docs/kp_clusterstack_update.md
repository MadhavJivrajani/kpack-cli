## kp clusterstack update

Update a cluster stack

### Synopsis

Updates the run and build images of a specific cluster-scoped stack.

The run and build images will be uploaded to the the registry configured on your stack.
Therefore, you must have credentials to access the registry on your machine.

```
kp clusterstack update <name> [flags]
```

### Examples

```
kp clusterstack update my-stack --build-image my-registry.com/build --run-image my-registry.com/run
kp clusterstack update my-stack --build-image ../path/to/build.tar --run-image ../path/to/run.tar
```

### Options

```
  -b, --build-image string             build image tag or local tar file path
      --dry-run                        perform validation with no side-effects; no objects are sent to the server.
                                         The --dry-run flag can be used in combination with the --output flag to
                                         view the Kubernetes resource(s) without sending anything to the server.
      --dry-run-with-image-upload      similar to --dry-run, but with container image uploads allowed.
                                         This flag is provided as a convenience for kp commands that can output Kubernetes
                                         resource with generated container image references. A "kubectl apply -f" of the
                                         resource from --output without image uploads will result in a reconcile failure.
  -h, --help                           help for update
      --output string                  print Kubernetes resources in the specified format; supported formats are: yaml, json.
                                         The output can be used with the "kubectl apply -f" command. To allow this, the command 
                                         updates are redirected to stderr and only the Kubernetes resource(s) are written to stdout.
      --registry-ca-cert-path string   add CA certificate for registry API (format: /tmp/ca.crt)
      --registry-verify-certs          set whether to verify server's certificate chain and host name (default true)
  -r, --run-image string               run image tag or local tar file path
```

### SEE ALSO

* [kp clusterstack](kp_clusterstack.md)	 - ClusterStack Commands

