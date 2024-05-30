package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup Kubernetes resources",
	Long:  `Backup Kubernetes resources to a file.`,
}

var backupNamespaceCmd = &cobra.Command{
	Use:   "namespace <namespace>",
	Short: "Backup all resources in a specified namespace",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		namespace := args[0]
		kubeconfig, _ := cmd.Flags().GetString("kubeconfig")

		config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			log.Fatalf("Error building kubeconfig: %v", err)
		}

		clientset, err := kubernetes.NewForConfig(config)
		if err != nil {
			log.Fatalf("Error creating Kubernetes client: %v", err)
		}

		resources := []string{"pods", "services", "deployments"}
		allResources := make(map[string]interface{})

		for _, resource := range resources {
			switch resource {
			case "pods":
				pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
				if err != nil {
					log.Fatalf("Error listing pods: %v", err)
				}
				allResources["pods"] = pods
			case "services":
				services, err := clientset.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{})
				if err != nil {
					log.Fatalf("Error listing services: %v", err)
				}
				allResources["services"] = services
			case "deployments":
				deployments, err := clientset.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
				if err != nil {
					log.Fatalf("Error listing deployments: %v", err)
				}
				allResources["deployments"] = deployments
			}
		}

		data, err := json.MarshalIndent(allResources, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling resources: %v", err)
		}

		outputDir, _ := cmd.Flags().GetString("output-dir")
		if outputDir == "" {
			outputDir = "."
		}

		filePath := filepath.Join(outputDir, fmt.Sprintf("backup-%s-%s.json", namespace, metav1.Now().Format("20060102150405")))
		err = os.WriteFile(filePath, data, 0644)
		if err != nil {
			log.Fatalf("Error writing backup file: %v", err)
		}

		fmt.Printf("Backup of namespace %s created at %s\n", namespace, filePath)
	},
}

func init() {
	rootCmd.AddCommand(backupCmd)
	backupCmd.AddCommand(backupNamespaceCmd)
	backupNamespaceCmd.Flags().String("output-dir", "", "Directory to save the backup file")
}
