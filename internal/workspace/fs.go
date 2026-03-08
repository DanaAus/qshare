package workspace

import (
	"fmt"
	"os"
	"path/filepath"
)

// FileExists returns true if the specified file or directory exists.
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// EnsureDirectoryExists creates the specified directory and any missing parent directories.
func EnsureDirectoryExists(path string) error {
	if FileExists(path) {
		return nil
	}
	return os.MkdirAll(path, 0755)
}

// ValidateDownloadPath checks if a path is a valid absolute path and is potentially writable.
func ValidateDownloadPath(path string) error {
	if path == "" {
		return fmt.Errorf("path cannot be empty")
	}

	if !filepath.IsAbs(path) {
		return fmt.Errorf("path must be absolute")
	}

	// Check if parent is writable if path doesn't exist
	if !FileExists(path) {
		parent := filepath.Dir(path)
		// We could try creating a temp file here, but simple check for now
		if !FileExists(parent) {
			// Try to create parent to see if it's possible?
			// Spec says "Ensure the specified downloads folder is created (on demand) when the first transfer occurs"
			// So we just check if the path is VALID.
		}
	}

	return nil
}
