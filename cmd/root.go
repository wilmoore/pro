package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"sort"
)

// Root command
var RootCmd = &cobra.Command{
	Use:   "pro",
	Short: "PRO is a CLI tool for provisioning, configuring, and managing cloud servers.",
	Long: `PRO is a CLI tool for provisioning, configuring, and managing cloud servers.
Supports multiple cloud providers including DigitalOcean, AWS, Azure, GCP, and OpenStack.`,
}

// Custom help function to separate commands
func customHelpFunc(cmd *cobra.Command, args []string) {
	fmt.Println("🚀 PRO CLI: Cloud Server Provisioning")
	fmt.Println("Usage:")
	fmt.Printf("  %s [command]\n\n", cmd.Name())

	// Sort commands for consistent ordering
	commands := cmd.Commands()
	sort.Slice(commands, func(i, j int) bool {
		return commands[i].Use < commands[j].Use
	})

	// Separate Cloud Providers from General Commands
	var cloudProviders []*cobra.Command
	var generalCommands []*cobra.Command

	for _, c := range commands {
		switch c.Use {
		case "help", "completion":
			generalCommands = append(generalCommands, c) // ✅ Only add `help` here
		default:
			cloudProviders = append(cloudProviders, c)
		}
	}

	// Print Cloud Providers section
	fmt.Println("🌐 Cloud Providers:")
	for _, c := range cloudProviders {
		fmt.Printf("  %-15s %s\n", c.Use, c.Short)
	}

	// Print General Commands section
	fmt.Println("\n⌘ General Commands:")
	for _, c := range generalCommands {
		fmt.Printf("  %-15s %s\n", c.Use, c.Short)
	}

	fmt.Println("\nUse 'pro [command] --help' for more details.")
}

func init() {
	// Add cloud provider subcommands
	RootCmd.AddCommand(digitaloceanCmd)
	RootCmd.AddCommand(awsCmd)
	RootCmd.AddCommand(azureCmd)
	RootCmd.AddCommand(gcpCmd)
	RootCmd.AddCommand(openstackCmd)

	// Override the default help output
	RootCmd.SetHelpFunc(customHelpFunc)
}

// Execute runs the root command
func Execute() error {
	return RootCmd.Execute()
}
