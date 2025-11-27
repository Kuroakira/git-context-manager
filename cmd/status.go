package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func Status(args []string) error {
	name, err := getGitConfig("user.name")
	if err != nil {
		return fmt.Errorf("failed to get user.name: %w", err)
	}

	email, err := getGitConfig("user.email")
	if err != nil {
		return fmt.Errorf("failed to get user.email: %w", err)
	}

	sshCommand, _ := getGitConfig("core.sshCommand") // Ignore error as it might not be set

	fmt.Println("Current Git Configuration (Local):")
	fmt.Printf("  user.name:  %s\n", name)
	fmt.Printf("  user.email: %s\n", email)
	if sshCommand != "" {
		fmt.Printf("  core.sshCommand: %s\n", sshCommand)
	} else {
		fmt.Println("  core.sshCommand: (not set)")
	}

	return nil
}

func getGitConfig(key string) (string, error) {
	cmd := exec.Command("git", "config", "--local", "--get", key)
	cmd.Stderr = os.Stderr
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}
