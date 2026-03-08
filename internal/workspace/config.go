package workspace

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config represents the magshare configuration.
type Config struct {
	Port        int    `json:"port"`
	SecureMode  bool   `json:"secure_mode"`
	PIN         string `json:"pin,omitempty"`
	DownloadDir string `json:"download_dir"`
}

// CreateDefaultConfig creates a configuration file at the specified path with the given config.
func CreateDefaultConfig(path string, cfg Config) error {
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

// LoadConfig reads and validates the configuration from the specified path.
func LoadConfig(path string) (Config, error) {
	var cfg Config

	data, err := os.ReadFile(path)
	if err != nil {
		return cfg, fmt.Errorf("could not read config file: %w", err)
	}

	if err := json.Unmarshal(data, &cfg); err != nil {
		return cfg, fmt.Errorf("could not parse config JSON: %w", err)
	}

	// Validation
	if cfg.Port < 1 || cfg.Port > 65535 {
		return cfg, fmt.Errorf("invalid port: %d (must be between 1 and 65535)", cfg.Port)
	}

	if cfg.DownloadDir != "" {
		if _, err := os.Stat(cfg.DownloadDir); os.IsNotExist(err) {
			return cfg, fmt.Errorf("download directory %q does not exist", cfg.DownloadDir)
		}
	}

	return cfg, nil
}
