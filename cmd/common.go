package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

// runCommand executes a shell command and returns output
func runCommand(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	cmd.Stderr = os.Stderr
	output, err := cmd.Output()
	return string(output), err
}

// validateAndGetRepoURL ensures the repo URL is correctly formatted
func validateAndGetRepoURL(repo string) string {
	if strings.HasPrefix(repo, "http") {
		return repo
	}
	return "https://github.com/" + repo + ".git"
}

// processTags combines default, provider-specific, and CLI tags
func processTags(providerTag string, cliTags string) string {
	defaultTags := "pro"
	allTags := []string{defaultTags, providerTag}
	if cliTags != "" {
		allTags = append(allTags, cliTags)
	}
	return strings.Join(allTags, ",")
}

// generateCloudInit creates a cloud-init script for the instance
func generateCloudInit(repoURL, branch, playbookPath string) string {
	return fmt.Sprintf(`#cloud-config
packages:
  - git
  - python3-pip
runcmd:
  - echo "Updating system packages..."
  - dnf install -y python3-pip
  - pip3 install --upgrade pip ansible
  - echo "Cloning repository..."
  - git clone -b %s %s /opt/ansible || (cd /opt/ansible && git pull)
  - echo "Running playbook..."
  - cd /opt/ansible/%s
  - ansible-playbook -i "localhost," -c local playbook.yml
`, branch, repoURL, playbookPath)
}

// getSSHKeys fetches the SSH key IDs from doctl
func getSSHKeys() string {
	output, err := runCommand("doctl", "compute", "ssh-key", "list", "--format", "ID", "--no-header")
	if err != nil {
		log.Fatalf("Error fetching SSH keys: %v", err)
	}
	return strings.TrimSpace(strings.ReplaceAll(output, "\n", ","))
}

// fzfSelect presents a list of options using fzf and returns the selected line
func fzfSelect(options []string) string {
	cmd := exec.Command("fzf")
	cmd.Stdin = strings.NewReader(strings.Join(options, "\n"))
	out, err := cmd.Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(out))
}
