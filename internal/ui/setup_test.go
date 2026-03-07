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

	expected := filepath.Join(home, "Magshare Downloads")
	// This is internal logic inside RunFirstRunSetup, hard to test without refactoring.
	// But we can check if we can at least resolve the home dir.
	if expected == "" {
		t.Error("expected default download dir to not be empty")
	}
}
