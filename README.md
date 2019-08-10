Example tool for cleaning up untrusted kubeconfig files.
Implements a simple approach, and while it can be used to check kubeconfig files, it's not particularly practical.

In it current state it mainly serves as an illustration for a blog post about [The dark side of sharing kubeconfig files](https://banzaicloud.com/blog/kubeconfig-security/)

We are happy to discuss the idea and the issue.

## Usage

```bash
go install github.com/banzaicloud/kubeconfiger/cmd/kubeconfiger
```

The tool will either write an error message to stderr, or copy a trustable single-context kubeconfig file to the standard output.
Use input redirection for saving the file (beware that it can't directly be used to filter a file in-place).

```bash
kubeconfiger < untrusted-config.yaml > trusted-config.yaml
```

At the time exec authentication helpers are supported only.
To whitelist a command, symlink it to `~/.kube/bin/`.
