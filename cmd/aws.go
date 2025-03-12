package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var awsCmd = &cobra.Command{
	Use:   "aws",
	Short: "Manage AWS EC2 instances",
}

var awsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new AWS EC2 instance",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Creating an AWS EC2 instance...")
		// Add API call logic
	},
}

func init() {
	awsCmd.AddCommand(awsCreateCmd)
}
