//go:build windows

package workspace

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/sys/windows/registry"
)

// RegisterContextMenu adds 'Share via Magshare' to the Windows right-click menu for files and directories.
func RegisterContextMenu() error {
	// Implementation will be added in Phase 2
	return nil
}

// UnregisterContextMenu removes Magshare from the Windows right-click menu.
func UnregisterContextMenu() error {
	// Implementation will be added in Phase 3
	return nil
}
