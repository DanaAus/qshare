package workspace

import (
	"os"
	"path/filepath"
	"runtime"
)

const AppName = "magshare"

// GetWorkspaceRoot returns the absolute path to the magshare workspace directory.
// On Windows, it prefers %LOCALAPPDATA%.
// On other platforms, it uses os.UserConfigDir().
func GetWorkspaceRoot() (string, error) {
	var baseDir string
	var err error

	if runtime.GOOS == "windows" {
		baseDir = os.Getenv("LOCALAPPDATA")
		if baseDir == "" {
			baseDir, err = os.UserConfigDir()
		}
	} else {
		baseDir, err = os.UserConfigDir()
	}

	if err != nil {
		return "", err
	}

	return filepath.Join(baseDir, AppName), nil
}

// GetLogsDir returns the absolute path to the magshare logs directory.
func GetLogsDir() (string, error) {
	root, err := GetWorkspaceRoot()
	if err != nil {
		return "", err
	}
	return filepath.Join(root, "logs"), nil
}
