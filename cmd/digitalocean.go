package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"github.com/spf13/cobra"
)

// DigitalOcean root command
var digitaloceanCmd = &cobra.Command{
	Use:   "digitalocean",
	Short: "Manage DigitalOcean droplets",
}

// DigitalOcean create command
var digitaloceanCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new DigitalOcean droplet",
	Run:   digitaloceanCreate,
}

// DigitalOcean SSH command
var digitaloceanSSHCmd = &cobra.Command{
	Use:   "ssh",
	Short: "SSH into a selected DigitalOcean droplet",
	Run:   digitaloceanSSH,
}

// Check if `doctl` is authenticated
func isDoctlAuthenticated() bool {
	_, err := runCommand("doctl", "account", "get")
	return err == nil
}

// Create DigitalOcean droplet
func digitaloceanCreate(cmd *cobra.Command, args []string) {
	// Check authentication before proceeding
	if !isDoctlAuthenticated() {
		log.Fatal("Error: DigitalOcean CLI (`doctl`) is not authenticated. Run `doctl auth init` first.")
	}

	repo, _ := cmd.Flags().GetString("repo")
	branch, _ := cmd.Flags().GetString("branch")
	playbookPath, _ := cmd.Flags().GetString("playbook-path")
	dropletName, _ := cmd.Flags().GetString("name")
	region, _ := cmd.Flags().GetString("region")
	size, _ := cmd.Flags().GetString("size")
	tags, _ := cmd.Flags().GetString("tags")

	// Validate repository
	if repo == "" {
		log.Fatal("Error: --repo flag is required")
	}
	repoURL := validateAndGetRepoURL(repo)

	// Set default values
	if branch == "" {
		branch = "main"
	}
	if playbookPath == "" {
		playbookPath = "src/pro"
	}
	if dropletName == "" {
		dropletName = "QuickSearch"
	}
	if region == "" {
		region = "sfo3"
	}
	if size == "" {
		size = "s-1vcpu-1gb"
	}

	// Process tags
	allTags := processTags("digitalocean", tags)

	// Generate Cloud Init script
	fmt.Println("Generating Cloud Init script for DigitalOcean...")
	cloudInit := generateCloudInit(repoURL, branch, playbookPath)

	// Execute doctl command to create the droplet
	fmt.Printf("Creating DigitalOcean droplet '%s' in region '%s' with tags: %s...\n", dropletName, region, allTags)
	createOutput, err := runCommand(
		"doctl", "compute", "droplet", "create", dropletName,
		"--region", region,
		"--image", "centos-stream-9-x64",
		"--size", size,
		"--ssh-keys", getSSHKeys(),
		"--user-data", cloudInit,
		"--tag-names", allTags,
	)
	if err != nil {
		log.Fatalf("Error: Failed to create droplet.\nOutput: %s", createOutput)
	}

	fmt.Printf("Droplet '%s' created successfully with tags: %s.\n", dropletName, allTags)
}

// Interactive SSH selection
func digitaloceanSSH(cmd *cobra.Command, args []string) {
	// Check authentication before proceeding
	if !isDoctlAuthenticated() {
		log.Fatal("Error: DigitalOcean CLI (`doctl`) is not authenticated. Run `doctl auth init` first.")
	}

	// Fetch list of droplets
	output, err := runCommand("doctl", "compute", "droplet", "list", "--format", "ID,Name,PublicIPv4", "--no-header")
	if err != nil {
		log.Fatal("Error: Could not retrieve droplets.")
	}

	// Show interactive droplet selection
	selected := fzfSelect(strings.Split(output, "\n")) // Use fzf for selection
	if selected == "" {
		log.Fatal("No droplet selected.")
	}
	ip := strings.Fields(selected)[2] // Extract Public IPv4

	// Connect via SSH with a forced TTY
	fmt.Printf("Connecting to %s...\n", ip)
	sshCommand := exec.Command("ssh", "-t", "root@"+ip) // ✅ Added `-t` for interactive TTY
	sshCommand.Stdin = os.Stdin
	sshCommand.Stdout = os.Stdout
	sshCommand.Stderr = os.Stderr

	err = sshCommand.Run()
	if err != nil {
		log.Fatalf("Error connecting via SSH: %v", err)
	}
}

// Initialize flags and register commands
func init() {
	digitaloceanCreateCmd.Flags().StringP("repo", "r", "", "Git repository to clone (required)")
	digitaloceanCreateCmd.Flags().StringP("branch", "b", "main", "Git branch")
	digitaloceanCreateCmd.Flags().StringP("playbook-path", "p", "src/pro", "Path to the playbook")
	digitaloceanCreateCmd.Flags().StringP("name", "n", "QuickSearch", "Droplet name")
	digitaloceanCreateCmd.Flags().StringP("region", "R", "sfo3", "Droplet region")
	digitaloceanCreateCmd.Flags().StringP("size", "s", "s-1vcpu-1gb", "Droplet size")
	digitaloceanCreateCmd.Flags().String("tags", "", "Comma-separated tags")

	// Register commands under digitaloceanCmd
	digitaloceanCmd.AddCommand(digitaloceanCreateCmd)
	digitaloceanCmd.AddCommand(digitaloceanSSHCmd)

	// Register digitaloceanCmd with root
	RootCmd.AddCommand(digitaloceanCmd)
}
