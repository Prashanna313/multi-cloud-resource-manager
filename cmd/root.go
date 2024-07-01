package cmd

import (
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "multi-cloud-resource-manager",
    Short: "A CLI tool to manage resources across AWS and GCP",
    Long:  `A CLI tool to manage resources across AWS and GCP.`,
}

func Execute() error {
    return rootCmd.Execute()
}
