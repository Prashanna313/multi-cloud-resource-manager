package cmd

import (
	"fmt"
	"multi-cloud-resource-manager/pkg/aws"

	"github.com/spf13/cobra"
)

var awsCmd = &cobra.Command{
	Use:   "aws",
	Short: "Manage AWS resources",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Managing AWS resources...")
		// Implement your logic here
	},
}

var listInstancesCmd = &cobra.Command{
	Use:   "list-instances",
	Short: "List EC2 instances",
	Run: func(cmd *cobra.Command, args []string) {
		aws.ListInstances()
	},
}

func init() {
	rootCmd.AddCommand(awsCmd)
	awsCmd.AddCommand(listInstancesCmd)
}
