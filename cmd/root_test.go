package cmd

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDemoFlag(t *testing.T) {
	flag := rootCmd.PersistentFlags().Lookup("demo")
	if flag == nil {
		t.Fatal("demo flag should be defined")
	}
	if flag.Name != "demo" {
		t.Errorf("expected flag name 'demo', got '%s'", flag.Name)
	}
	if flag.Value.Type() != "bool" {
		t.Errorf("expected flag type 'bool', got '%s'", flag.Value.Type())
	}
}

func TestInitConfig(t *testing.T) {
	// Create a temporary directory for the workspace
	tmpDir, err := os.MkdirTemp("", "magshare_cmd_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// Mock LOCALAPPDATA to point to our temp dir
	originalAppData := os.Getenv("LOCALAPPDATA")
	os.Setenv("LOCALAPPDATA", tmpDir)
	defer os.Setenv("LOCALAPPDATA", originalAppData)

	// Create magshare folder and config.json
	workspaceDir := filepath.Join(tmpDir, "magshare")
	os.MkdirAll(workspaceDir, 0755)
	configPath := filepath.Join(workspaceDir, "config.json")
	
	configContent := `{"port": 9999, "download_dir": ".", "secure_mode": true}`
	os.WriteFile(configPath, []byte(configContent), 0644)

	// Run initConfig
	initConfig()

	// Verify appConfig was populated
	if appConfig.Port != 9999 {
		t.Errorf("expected port 9999, got %d", appConfig.Port)
	}
	if !appConfig.SecureMode {
		t.Error("expected secure mode to be true")
	}
}
