package workspace

import (
	"path/filepath"
)

// SetupProvider is a function that returns a Config object, typically from user input.
type SetupProvider func() (Config, error)

// InitializeWorkspace sets up the workspace directory structure if it doesn't exist.
// Returns isFirstRun=true if the setup was performed.
func InitializeWorkspace(setup SetupProvider) (bool, error) {
	root, err := GetWorkspaceRoot()
	if err != nil {
		return false, err
	}
	return InitializeWorkspaceAtPath(root, setup)
}

// InitializeWorkspaceAtPath sets up the workspace directory structure at the specified path.
func InitializeWorkspaceAtPath(root string, setup SetupProvider) (bool, error) {
	if err := EnsureDirectoryExists(root); err != nil {
		return false, err
	}

	logsDir := filepath.Join(root, "logs")
	if err := EnsureDirectoryExists(logsDir); err != nil {
		return false, err
	}

	configPath := filepath.Join(root, "config.json")
	if !FileExists(configPath) {
		if setup == nil {
			// Fallback to default if no setup provider
			defaultCfg := Config{
				Port:       8080,
				SecureMode: false,
			}
			return true, CreateDefaultConfig(configPath, defaultCfg)
		}

		cfg, err := setup()
		if err != nil {
			return false, err
		}

		if err := CreateDefaultConfig(configPath, cfg); err != nil {
			return false, err
		}

		// Create download dir if provided
		if cfg.DownloadDir != "" {
			if err := EnsureDirectoryExists(cfg.DownloadDir); err != nil {
				// We don't fail if download dir creation fails, but maybe we should?
				// Spec says "automatically create the downloads folder"
				return true, err
			}
		}

		return true, nil
	}

	return false, nil
}
