package cmd

import (
	"flag"
	"fmt"

	"github.com/Kuroakira/git-context-manager/config"
)

func Add(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("context name is required")
	}

	contextName := args[0]

	// Shift args to parse flags
	flagArgs := args[1:]

	fs := flag.NewFlagSet("add", flag.ExitOnError)
	email := fs.String("email", "", "Git user email")
	name := fs.String("name", "", "Git user name")
	sshKey := fs.String("ssh-key", "", "Path to SSH key")

	if err := fs.Parse(flagArgs); err != nil {
		return err
	}

	if *email == "" || *name == "" {
		return fmt.Errorf("email and name are required")
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	cfg.Contexts[contextName] = config.Context{
		Name:       contextName,
		Email:      *email,
		UserName:   *name,
		SSHKeyPath: *sshKey,
	}

	if err := config.SaveConfig(cfg); err != nil {
		return fmt.Errorf("failed to save config: %w", err)
	}

	fmt.Printf("Context '%s' added successfully.\n", contextName)
	return nil
}
