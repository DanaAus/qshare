//go:build windows

package workspace

import (
	"errors"
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

func TestUnregisterContextMenu(t *testing.T) {
	t.Run("unregister when registered", func(t *testing.T) {
		err := RegisterContextMenu()
		if err != nil {
			t.Fatalf("RegisterContextMenu() returned error: %v", err)
		}

		err = UnregisterContextMenu()
		if err != nil {
			t.Errorf("UnregisterContextMenu() returned error: %v", err)
		}

		targets := []string{
			`Software\Classes\*\shell\Magshare`,
			`Software\Classes\Directory\shell\Magshare`,
		}

		for _, target := range targets {
			key, err := registry.OpenKey(registry.CURRENT_USER, target, registry.QUERY_VALUE)
			if err == nil {
				key.Close()
				t.Errorf("registry key %s still exists after unregistration", target)
			}
		}
	})

	t.Run("unregister when already unregistered", func(t *testing.T) {
		_ = UnregisterContextMenu()
		err := UnregisterContextMenu()
		if err != nil {
			t.Errorf("UnregisterContextMenu() failed when keys don't exist: %v", err)
		}
	})

	t.Run("IsContextMenuItemRegistered", func(t *testing.T) {
		_ = UnregisterContextMenu()
		if IsContextMenuItemRegistered() {
			t.Error("expected false when unregistered")
		}
		_ = RegisterContextMenu()
		if !IsContextMenuItemRegistered() {
			t.Error("expected true when registered")
		}
		_ = UnregisterContextMenu()
	})
}

// MockKey implements parts of registry.Key interface if needed, but registry.Key is just a handle.
// We can't easily mock methods on registry.Key directly since it's a uint32.

func TestRegistryMocks(t *testing.T) {
	// Backup original functions
	origCreate := registryCreateKey
	origOpen := registryOpenKey
	origDelete := registryDeleteKey
	defer func() {
		registryCreateKey = origCreate
		registryOpenKey = origOpen
		registryDeleteKey = origDelete
	}()

	t.Run("RegisterContextMenu failure on first CreateKey", func(t *testing.T) {
		registryCreateKey = func(k registry.Key, path string, access uint32) (registry.Key, bool, error) {
			return 0, false, errors.New("mock error")
		}
		err := RegisterContextMenu()
		if err == nil {
			t.Error("expected error, got nil")
		}
	})

	t.Run("RegisterContextMenu failure on second CreateKey (command)", func(t *testing.T) {
		registryCreateKey = func(k registry.Key, path string, access uint32) (registry.Key, bool, error) {
			if path == "command" {
				return 0, false, errors.New("mock error")
			}
			k, err := registry.OpenKey(registry.CURRENT_USER, `Software\Classes`, registry.ALL_ACCESS) // just return a valid handle
			return k, false, err
		}
		err := RegisterContextMenu()
		if err == nil {
			t.Error("expected error, got nil")
		}
	})

	t.Run("RegisterContextMenu failure on second target", func(t *testing.T) {
		count := 0
		registryCreateKey = func(k registry.Key, path string, access uint32) (registry.Key, bool, error) {
			if path != "command" {
				count++
			}
			if count == 2 {
				return 0, false, errors.New("mock error")
			}
			k, err := registry.OpenKey(registry.CURRENT_USER, `Software\Classes`, registry.ALL_ACCESS) // just return a valid handle
			return k, false, err
		}
		err := RegisterContextMenu()
		if err == nil {
			t.Error("expected error, got nil")
		}
	})

	t.Run("UnregisterContextMenu failure on parent OpenKey", func(t *testing.T) {
		registryDeleteKey = func(k registry.Key, path string) error {
			return nil
		}
		registryOpenKey = func(k registry.Key, path string, access uint32) (registry.Key, error) {
			if strings.Contains(path, "shell") {
				return 0, errors.New("mock error")
			}
			return registry.OpenKey(registry.CURRENT_USER, `Software\Classes`, registry.ALL_ACCESS)
		}
		err := UnregisterContextMenu()
		if err != nil {
			t.Errorf("expected nil (silent failure), got %v", err)
		}
	})
}
