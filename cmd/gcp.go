package cmd

import (
    "fmt"
    "multi-cloud-resource-manager/pkg/gcp"
    "github.com/spf13/cobra"
)

var gcpCmd = &cobra.Command{
    Use:   "gcp",
    Short: "Manage GCP resources",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Managing GCP resources...")
        // Implement your logic here
    },
}

var listInstancesCmd = &cobra.Command{
    Use:   "list-instances",
    Short: "List GCP instances",
    Run: func(cmd *cobra.Command, args []string) {
        gcp.ListInstances()
    },
}

func init() {
    rootCmd.AddCommand(gcpCmd)
    gcpCmd.AddCommand(listInstancesCmd)
}
