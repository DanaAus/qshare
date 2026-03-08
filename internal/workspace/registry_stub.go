//go:build !windows

package workspace

import "errors"

// RegisterContextMenu is a no-op stub for non-Windows platforms.
func RegisterContextMenu() error {
	return errors.New("context menu registration is only supported on Windows")
}

// UnregisterContextMenu is a no-op stub for non-Windows platforms.
func UnregisterContextMenu() error {
	return errors.New("context menu removal is only supported on Windows")
}

// IsContextMenuItemRegistered is a no-op stub for non-Windows platforms.
func IsContextMenuItemRegistered() bool {
	return false
}
