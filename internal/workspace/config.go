package workspace

import (
	"encoding/json"
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
