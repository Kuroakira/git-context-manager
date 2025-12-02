package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/Kuroakira/git-context-manager/config"
)

func Clone(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("usage: gctx clone <context-name> <repository-url> [directory]")
	}

	contextName := args[0]
	repoURL := args[1]

	// Optional: target directory
	var targetDir string
	if len(args) >= 3 {
		targetDir = args[2]
	} else {
		// Extract directory name from URL
		targetDir = extractRepoName(repoURL)
	}

	// Load config and find context
	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	ctx, exists := cfg.Contexts[contextName]
	if !exists {
		return fmt.Errorf("context '%s' not found", contextName)
	}

	// Convert HTTPS URL to SSH format if SSH key is specified
	cloneURL := repoURL
	if ctx.SSHKeyPath != "" {
		convertedURL, converted := convertToSSHURL(repoURL)
		if converted {
			cloneURL = convertedURL
			fmt.Printf("Converted URL to SSH format: %s\n", cloneURL)
		}
	}

	// Build git clone command
	gitArgs := []string{"clone", cloneURL, targetDir}
	cmd := exec.Command("git", gitArgs...)

	// Set GIT_SSH_COMMAND if SSH key is specified
	if ctx.SSHKeyPath != "" {
		sshCommand := fmt.Sprintf("ssh -i %s -o IdentitiesOnly=yes", ctx.SSHKeyPath)
		cmd.Env = append(os.Environ(), "GIT_SSH_COMMAND="+sshCommand)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("Cloning with context '%s'...\n", contextName)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git clone failed: %w", err)
	}

	// After clone, apply the context to the new repository
	originalDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}

	if err := os.Chdir(targetDir); err != nil {
		return fmt.Errorf("failed to change to cloned directory: %w", err)
	}

	// Apply context using existing Use logic (reuse Use function)
	if err := Use([]string{contextName}); err != nil {
		os.Chdir(originalDir)
		return fmt.Errorf("failed to apply context: %w", err)
	}

	os.Chdir(originalDir)
	return nil
}

// extractRepoName extracts the repository name from a git URL
func extractRepoName(url string) string {
	// Remove trailing .git
	url = strings.TrimSuffix(url, ".git")

	// Handle SSH format: git@github.com:user/repo
	if strings.Contains(url, ":") && strings.HasPrefix(url, "git@") {
		parts := strings.Split(url, ":")
		if len(parts) == 2 {
			url = parts[1]
		}
	}

	// Get the last part of the path
	return filepath.Base(url)
}

// convertToSSHURL converts HTTPS git URL to SSH format
// Returns the converted URL and true if conversion was done, or original URL and false if not needed
func convertToSSHURL(url string) (string, bool) {
	// Already SSH format
	if strings.HasPrefix(url, "git@") {
		return url, false
	}

	// Convert https://github.com/user/repo.git -> git@github.com:user/repo.git
	// Convert https://github.com/user/repo -> git@github.com:user/repo.git
	httpsPrefix := "https://"
	if strings.HasPrefix(url, httpsPrefix) {
		// Remove https://
		remainder := strings.TrimPrefix(url, httpsPrefix)

		// Split host and path: github.com/user/repo
		slashIndex := strings.Index(remainder, "/")
		if slashIndex == -1 {
			return url, false
		}

		host := remainder[:slashIndex]
		path := remainder[slashIndex+1:]

		// Ensure .git suffix
		if !strings.HasSuffix(path, ".git") {
			path = path + ".git"
		}

		sshURL := fmt.Sprintf("git@%s:%s", host, path)
		return sshURL, true
	}

	return url, false
}
