package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Context represents a git configuration context
type Context struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	UserName   string `json:"user_name"`
	SSHKeyPath string `json:"ssh_key_path"`
}

// Config represents the application configuration
type Config struct {
	Contexts map[string]Context `json:"contexts"`
}

// GetConfigPath returns the configuration file path
func GetConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".gctxconfig"), nil
}

// LoadConfig loads the configuration from disk
func LoadConfig() (*Config, error) {
	path, err := GetConfigPath()
	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return &Config{Contexts: make(map[string]Context)}, nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	if cfg.Contexts == nil {
		cfg.Contexts = make(map[string]Context)
	}

	return &cfg, nil
}

// SaveConfig saves the configuration to disk
func SaveConfig(cfg *Config) error {
	path, err := GetConfigPath()
	if err != nil {
		return err
	}

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
