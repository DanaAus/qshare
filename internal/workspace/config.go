package workspace

import (
	"encoding/json"
	"os"
)

// Config represents the magshare configuration.
type Config struct {
	Port       int    `json:"port"`
	SecureMode bool   `json:"secure_mode"`
	PIN        string `json:"pin,omitempty"`
}

// CreateDefaultConfig creates a default configuration file at the specified path.
func CreateDefaultConfig(path string) error {
	defaultConfig := Config{
		Port:       8080,
		SecureMode: false,
	}

	data, err := json.MarshalIndent(defaultConfig, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
