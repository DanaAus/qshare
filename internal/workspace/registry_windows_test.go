//go:build windows

package workspace

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"golang.org/x/sys/windows/registry"
)

func TestRegisterContextMenu(t *testing.T) {
	// Clean up any existing keys before starting (best effort)
	_ = UnregisterContextMenu()

	err := RegisterContextMenu()
	if err != nil {
		t.Fatalf("RegisterContextMenu() returned error: %v", err)
	}

	exePath, err := os.Executable()
	if err != nil {
		t.Fatal(err)
	}
	exePath, err = filepath.Abs(exePath)
	if err != nil {
		t.Fatal(err)
	}

	expectedCommand := "\"" + exePath + "\" send \"%1\""
	_ = expectedCommand
	expectedMenuTitle := "Share via Magshare"

	targets := []string{
		`Software\Classes\*\shell\Magshare`,
		`Software\Classes\Directory\shell\Magshare`,
	}

	for _, target := range targets {
		t.Run(target, func(t *testing.T) {
			// Check shell key
			key, err := registry.OpenKey(registry.CURRENT_USER, target, registry.QUERY_VALUE)
			if err != nil {
				t.Errorf("failed to open registry key %s: %v", target, err)
				return
			}
			defer key.Close()

			title, _, err := key.GetStringValue("")
			if err != nil {
				t.Errorf("failed to get default value for %s: %v", target, err)
			} else if title != expectedMenuTitle {
				t.Errorf("menu title = %q, want %q", title, expectedMenuTitle)
			}

			// Check command subkey
			cmdKeyPath := target + `\command`
			cmdKey, err := registry.OpenKey(registry.CURRENT_USER, cmdKeyPath, registry.QUERY_VALUE)
			if err != nil {
				t.Errorf("failed to open command key %s: %v", cmdKeyPath, err)
				return
			}
			defer cmdKey.Close()

			command, _, err := cmdKey.GetStringValue("")
			if err != nil {
				t.Errorf("failed to get default value for %s: %v", cmdKeyPath, err)
			} else if !strings.Contains(command, exePath) || !strings.Contains(command, "send") {
				t.Errorf("command = %q, want it to contain %q and %q", command, exePath, "send \"%1\"")
			}
		})
	}

	// Clean up after test
	_ = UnregisterContextMenu()
}
