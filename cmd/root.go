package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "k8s-backup-restore",
	Short: "A CLI tool for backing up and restoring Kubernetes resources",
	Long:  `A simple CLI tool to back up and restore Kubernetes resources using Golang and Cobra.`,
}

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("kubeconfig", "k", "", "Path to the kubeconfig file")
}
