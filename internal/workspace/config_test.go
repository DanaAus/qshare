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

	expectedConfig := Config{
		Port:        9090,
		SecureMode:  true,
		DownloadDir: "/tmp/downloads",
	}

	err = CreateDefaultConfig(configPath, expectedConfig)
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

	var actualConfig Config
	if err := json.Unmarshal(data, &actualConfig); err != nil {
		t.Fatalf("Failed to unmarshal config: %v", err)
	}

	// Check values
	if actualConfig.Port != expectedConfig.Port {
		t.Errorf("Port = %d, want %d", actualConfig.Port, expectedConfig.Port)
	}
	if actualConfig.SecureMode != expectedConfig.SecureMode {
		t.Errorf("SecureMode = %v, want %v", actualConfig.SecureMode, expectedConfig.SecureMode)
	}
	if actualConfig.DownloadDir != expectedConfig.DownloadDir {
		t.Errorf("DownloadDir = %q, want %q", actualConfig.DownloadDir, expectedConfig.DownloadDir)
	}
}
