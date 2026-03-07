package workspace

import (
	"os"
	"path/filepath"
	"testing"
)

func TestInitializeWorkspaceAtPath(t *testing.T) {
	tmpBase, err := os.MkdirTemp("", "magshare_init_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpBase)

	testRoot := filepath.Join(tmpBase, "magshare_test")
	downloadDir := filepath.Join(tmpBase, "downloads")
	
	mockSetup := func() (Config, error) {
		return Config{
			Port:        1234,
			DownloadDir: downloadDir,
			SecureMode:  true,
		}, nil
	}

	// 1. Initial run: should create folders, config, and download dir, return isFirstRun = true
	isFirstRun, err := InitializeWorkspaceAtPath(testRoot, mockSetup)
	if err != nil {
		t.Fatalf("InitializeWorkspaceAtPath returned error: %v", err)
	}
	if !isFirstRun {
		t.Error("InitializeWorkspaceAtPath returned isFirstRun = false, want true")
	}

	// Verify structure
	if !FileExists(testRoot) {
		t.Error("Workspace root not created")
	}
	if !FileExists(filepath.Join(testRoot, "logs")) {
		t.Error("Logs directory not created")
	}
	if !FileExists(filepath.Join(testRoot, "config.json")) {
		t.Error("Config file not created")
	}
	if !FileExists(downloadDir) {
		t.Error("Download directory not created")
	}

	// 2. Second run: should return isFirstRun = false
	isFirstRun, err = InitializeWorkspaceAtPath(testRoot, mockSetup)
	if err != nil {
		t.Fatalf("InitializeWorkspaceAtPath (second run) returned error: %v", err)
	}
	if isFirstRun {
		t.Error("InitializeWorkspaceAtPath (second run) returned isFirstRun = true, want false")
	}
}

func TestInitializeWorkspace(t *testing.T) {
	// Simple smoke test
	_, err := InitializeWorkspace(nil)
	if err != nil {
		t.Errorf("InitializeWorkspace() returned error: %v", err)
	}
}
