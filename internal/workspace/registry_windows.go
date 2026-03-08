//go:build windows

package workspace

import (
	"fmt"
	"magshare/internal/logger"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/sys/windows/registry"
)

// For testing purposes
var (
	registryCreateKey = registry.CreateKey
	registryOpenKey   = registry.OpenKey
	registryDeleteKey = registry.DeleteKey
)

// RegisterContextMenu adds 'Share via Magshare' to the Windows right-click menu for files and directories.
func RegisterContextMenu() error {
	log := logger.WithComponent("registry")
	log.Info("Registering Windows context menu...")

	exePath, err := os.Executable()
	if err != nil {
		err = fmt.Errorf("failed to get executable path: %w", err)
		log.Error(err.Error())
		return err
	}
	exePath, err = filepath.Abs(exePath)
	if err != nil {
		err = fmt.Errorf("failed to get absolute path: %w", err)
		log.Error(err.Error())
		return err
	}

	// Command: "C:\path\to\magshare.exe" send "%1"
	command := fmt.Sprintf("\"%s\" send \"%%1\"", exePath)
	const menuTitle = "Share via Magshare"

	targets := []string{
		`Software\Classes\*\shell\Magshare`,
		`Software\Classes\Directory\shell\Magshare`,
	}

	for _, target := range targets {
		// Create/Open the shell key
		key, _, err := registryCreateKey(registry.CURRENT_USER, target, registry.ALL_ACCESS)
		if err != nil {
			err = fmt.Errorf("failed to create registry key %s: %w", target, err)
			log.Error(err.Error())
			return err
		}

		if err := key.SetStringValue("", menuTitle); err != nil {
			key.Close()
			err = fmt.Errorf("failed to set menu title for %s: %w", target, err)
			log.Error(err.Error())
			return err
		}

		// Create/Open the command subkey
		cmdKey, _, err := registryCreateKey(key, "command", registry.ALL_ACCESS)
		key.Close() // Close parent
		if err != nil {
			err = fmt.Errorf("failed to create command key for %s: %w", target, err)
			log.Error(err.Error())
			return err
		}

		if err := cmdKey.SetStringValue("", command); err != nil {
			cmdKey.Close()
			err = fmt.Errorf("failed to set command for %s: %w", target, err)
			log.Error(err.Error())
			return err
		}
		cmdKey.Close()
	}

	log.Info("Successfully registered Windows context menu.")
	return nil
}

// UnregisterContextMenu removes Magshare from the Windows right-click menu.
func UnregisterContextMenu() error {
	log := logger.WithComponent("registry")
	log.Info("Unregistering Windows context menu...")

	targets := []string{
		`Software\Classes\*\shell\Magshare`,
		`Software\Classes\Directory\shell\Magshare`,
	}

	for _, target := range targets {
		// First, delete the "command" subkey
		err := registryDeleteKey(registry.CURRENT_USER, target+`\command`)
		if err != nil && err != registry.ErrNotExist {
			log.Warn(fmt.Sprintf("failed to delete command key for %s: %v", target, err))
		}

		// Now delete the Magshare key itself
		lastSlash := strings.LastIndex(target, `\`)
		if lastSlash == -1 {
			continue
		}
		parentPath := target[:lastSlash]
		childName := target[lastSlash+1:]

		parentKey, err := registryOpenKey(registry.CURRENT_USER, parentPath, registry.ALL_ACCESS)
		if err != nil {
			if err != registry.ErrNotExist {
				log.Warn(fmt.Errorf("failed to open parent key %s: %w", parentPath, err).Error())
			}
			continue
		}
		
		err = registryDeleteKey(parentKey, childName)
		parentKey.Close()
		if err != nil && err != registry.ErrNotExist {
			log.Warn(fmt.Sprintf("failed to delete registry key %s: %v", target, err))
		}
	}

	log.Info("Successfully unregistered Windows context menu.")
	return nil
}
