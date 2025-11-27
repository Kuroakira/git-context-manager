package cmd

import (
	"fmt"

	"github.com/Kuroakira/git-context-manager/config"
)

func Remove(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("context name is required")
	}
	contextName := args[0]

	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	if _, ok := cfg.Contexts[contextName]; !ok {
		return fmt.Errorf("context '%s' not found", contextName)
	}

	delete(cfg.Contexts, contextName)

	if err := config.SaveConfig(cfg); err != nil {
		return fmt.Errorf("failed to save config: %w", err)
	}

	fmt.Printf("Context '%s' removed successfully.\n", contextName)
	return nil
}
