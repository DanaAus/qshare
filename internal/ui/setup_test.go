package ui

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDefaultDownloadDir(t *testing.T) {
	home, err := os.UserHomeDir()
	if err != nil {
		t.Skip("skipping test; home directory not available")
	}

	got := filepath.Join(home, "Magshare Downloads")
	if !filepath.IsAbs(got) {
		t.Errorf("expected absolute path, got %q", got)
	}
}

func TestSetupResultStructure(t *testing.T) {
	result := &SetupResult{
		DownloadDir:      "/tmp/magshare",
		SecureMode:       true,
		ShellIntegration: true,
	}
	if result.DownloadDir != "/tmp/magshare" {
		t.Error("DownloadDir not set correctly")
	}
	if !result.SecureMode {
		t.Error("SecureMode not set correctly")
	}
	if !result.ShellIntegration {
		t.Error("ShellIntegration not set correctly")
	}
}
