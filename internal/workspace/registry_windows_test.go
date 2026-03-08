//go:build windows

package workspace

import "testing"

func TestRegisterContextMenu(t *testing.T) {
	// For Phase 1, we just verify the function exists and can be called.
	// Actual TDD will happen in Phase 2.
	err := RegisterContextMenu()
	if err != nil {
		t.Errorf("RegisterContextMenu() returned unexpected error: %v", err)
	}
}

func TestUnregisterContextMenu(t *testing.T) {
	// For Phase 1, we just verify the function exists and can be called.
	// Actual TDD will happen in Phase 3.
	err := UnregisterContextMenu()
	if err != nil {
		t.Errorf("UnregisterContextMenu() returned unexpected error: %v", err)
	}
}
