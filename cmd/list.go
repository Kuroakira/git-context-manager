package cmd

import (
	"fmt"
	"sort"

	"github.com/Kuroakira/git-context-manager/config"
)

func List(args []string) error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	if len(cfg.Contexts) == 0 {
		fmt.Println("No contexts found.")
		return nil
	}

	var names []string
	for name := range cfg.Contexts {
		names = append(names, name)
	}
	sort.Strings(names)

	fmt.Println("Available contexts:")
	for _, name := range names {
		ctx := cfg.Contexts[name]
		fmt.Printf("- %s (Name: %s, Email: %s", name, ctx.UserName, ctx.Email)
		if ctx.SSHKeyPath != "" {
			fmt.Printf(", SSH Key: %s", ctx.SSHKeyPath)
		}
		fmt.Println(")")
	}

	return nil
}
