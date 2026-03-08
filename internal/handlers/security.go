package handlers

import (
	"fmt"
	"path/filepath"
	"strings"
)

// SanitizePath ensures that the target path is within the base directory.
// It returns the absolute path to the target if valid, or an error if not.
func SanitizePath(base, target string) (string, error) {
	absBase, err := filepath.Abs(base)
	if err != nil {
		return "", fmt.Errorf("could not get absolute base path: %w", err)
	}

	absTarget, err := filepath.Abs(target)
	if err != nil {
		return "", fmt.Errorf("could not get absolute target path: %w", err)
	}

	// On Windows, filepath.Abs might return paths with different casing or separators.
	// Clean them to be sure.
	absBase = filepath.Clean(absBase)
	absTarget = filepath.Clean(absTarget)

	// Check if target is inside base.
	// We use a trailing separator to ensure /tmp/test doesn't match /tmp/testing.
	if absTarget != absBase && !strings.HasPrefix(absTarget, absBase+string(filepath.Separator)) {
		return "", fmt.Errorf("security error: path %q is outside base directory %q", target, base)
	}

	return absTarget, nil
}
