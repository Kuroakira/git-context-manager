package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/Kuroakira/git-context-manager/config"
)

func Use(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("context name is required")
	}
	contextName := args[0]

	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	ctx, ok := cfg.Contexts[contextName]
	if !ok {
		return fmt.Errorf("context '%s' not found", contextName)
	}

	// Apply git config
	if err := setGitConfig("user.name", ctx.UserName); err != nil {
		return err
	}
	if err := setGitConfig("user.email", ctx.Email); err != nil {
		return err
	}

	if ctx.SSHKeyPath != "" {
		// Use core.sshCommand to specify the key
		sshCommand := fmt.Sprintf("ssh -i %s -o IdentitiesOnly=yes", ctx.SSHKeyPath)
		if err := setGitConfig("core.sshCommand", sshCommand); err != nil {
			return err
		}
	} else {
		// Unset core.sshCommand if no key is specified to fall back to default
		// We ignore error here as it might not be set
		_ = unsetGitConfig("core.sshCommand")
	}

	fmt.Printf("Switched to context '%s'\n", contextName)
	return nil
}

func setGitConfig(key, value string) error {
	cmd := exec.Command("git", "config", "--local", key, value)
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to set git config %s: %w", key, err)
	}
	return nil
}

func unsetGitConfig(key string) error {
	cmd := exec.Command("git", "config", "--local", "--unset", key)
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to unset git config %s: %w", key, err)
	}
	return nil
}
