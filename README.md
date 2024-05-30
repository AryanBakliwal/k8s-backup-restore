# Backup and Restore CLI tool

Kubernetes backup and restore CLI tool using Golang and spf13/cobra.

### How to build
Install dependencies: Run `go mod tidy` to download the necessary dependencies.
Build the CLI tool: Run 
```
go build -o k8s-backup-restore
``` 
to build the binary.

### Run commands:
1. Backup a namespace: 
```
./k8s-backup-restore backup namespace <namespace> --output-dir=<output-directory> --kubeconfig <path/to/your/kubeconfig>
```
