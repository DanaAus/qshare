package workspace

import (
	"path/filepath"
)

// InitializeWorkspace sets up the workspace directory structure if it doesn't exist.
// Returns isFirstRun=true if the workspace was newly created.
func InitializeWorkspace() (bool, error) {
	root, err := GetWorkspaceRoot()
	if err != nil {
		return false, err
	}
	return InitializeWorkspaceAtPath(root)
}

// InitializeWorkspaceAtPath sets up the workspace directory structure at the specified path.
func InitializeWorkspaceAtPath(root string) (bool, error) {
	isFirstRun := false
	if !FileExists(root) {
		isFirstRun = true
	}

	if err := EnsureDirectoryExists(root); err != nil {
		return false, err
	}

	logsDir := filepath.Join(root, "logs")
	if err := EnsureDirectoryExists(logsDir); err != nil {
		return false, err
	}

	if isFirstRun {
		configPath := filepath.Join(root, "config.json")
		if err := CreateDefaultConfig(configPath); err != nil {
			return false, err
		}
	}

	return isFirstRun, nil
}
