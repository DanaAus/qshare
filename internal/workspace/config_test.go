package workspace

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestCreateDefaultConfig(t *testing.T) {
	tmpdir, err := os.MkdirTemp("", "config_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpdir)

	configPath := filepath.Join(tmpdir, "config.json")
	
	err = CreateDefaultConfig(configPath)
	if err != nil {
		t.Fatalf("CreateDefaultConfig returned error: %v", err)
	}

	// Verify file exists
	if !FileExists(configPath) {
		t.Error("Config file was not created")
	}

	// Verify content
	data, err := os.ReadFile(configPath)
	if err != nil {
		t.Fatal(err)
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		t.Fatalf("Failed to unmarshal config: %v", err)
	}

	// Check some default values
	if config.Port == 0 {
		t.Error("Default Port should not be 0")
	}
}
