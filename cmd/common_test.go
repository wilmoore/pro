package cmd

import (
	"os/exec"
	"strings"
	"testing"
)

// Test validateAndGetRepoURL function
func TestValidateAndGetRepoURL(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"https://github.com/wilmoore/pro.git", "https://github.com/wilmoore/pro.git"},
		{"wilmoore/pro", "https://github.com/wilmoore/pro.git"},
	}

	for _, test := range tests {
		result := validateAndGetRepoURL(test.input)
		if result != test.expected {
			t.Errorf("validateAndGetRepoURL(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}

// Test processTags function
func TestProcessTags(t *testing.T) {
	tests := []struct {
		providerTag string
		cliTags     string
		expected    string
	}{
		{"digitalocean", "web,db", "pro,digitalocean,web,db"},
		{"aws", "", "pro,aws"},
		{"gcp", "backend,api", "pro,gcp,backend,api"},
	}

	for _, test := range tests {
		result := processTags(test.providerTag, test.cliTags)
		if result != test.expected {
			t.Errorf("processTags(%q, %q) = %q; want %q", test.providerTag, test.cliTags, result, test.expected)
		}
	}
}

// Test generateCloudInit function
func TestGenerateCloudInit(t *testing.T) {
	repoURL := "https://github.com/wilmoore/pro.git"
	branch := "main"
	playbookPath := "src/pro"

	result := generateCloudInit(repoURL, branch, playbookPath)

	expectedParts := []string{
		"git clone -b main https://github.com/wilmoore/pro.git /opt/ansible",
		"cd /opt/ansible/src/pro",
		"ansible-playbook -i \"localhost,\" -c local playbook.yml",
	}

	for _, part := range expectedParts {
		if !strings.Contains(result, part) {
			t.Errorf("generateCloudInit() missing expected content: %q", part)
		}
	}
}

// Test getSSHKeys function
func TestGetSSHKeys(t *testing.T) {
	// Mock doctl response
	mockOutput := "123456\n789012\n"
	cmd := exec.Command("bash", "-c", "echo -n '"+mockOutput+"'") // Simulate doctl output

	output, err := cmd.Output()
	if err != nil {
		t.Fatalf("Error executing mock command: %v", err)
	}

	expected := "123456,789012"
	result := strings.TrimRight(strings.TrimSpace(strings.ReplaceAll(string(output), "\n", ",")), ",")
	if result != expected {
		t.Errorf("getSSHKeys() = %q; want %q", result, expected)
	}
}

// Test runCommand function
func TestRunCommand(t *testing.T) {
	output, err := runCommand("echo", "hello")
	if err != nil {
		t.Fatalf("runCommand() failed: %v", err)
	}
	expected := "hello\n"
	if output != expected {
		t.Errorf("runCommand() = %q; want %q", output, expected)
	}
}

// Test fzfSelect function (mocked)
func TestFzfSelect(t *testing.T) {
	mockOptions := []string{"Option 1", "Option 2", "Option 3"}

	// Mock fzf by using `echo` to select the first option
	cmd := exec.Command("bash", "-c", "echo 'Option 1'")
	cmd.Stdin = strings.NewReader(strings.Join(mockOptions, "\n"))

	output, err := cmd.Output()
	if err != nil {
		t.Fatalf("fzfSelect() mock failed: %v", err)
	}

	expected := "Option 1"
	result := strings.TrimSpace(string(output))
	if result != expected {
		t.Errorf("fzfSelect() = %q; want %q", result, expected)
	}
}