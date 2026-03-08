//go:build !windows

package workspace

import "testing"

func TestRegisterContextMenu_Stub(t *testing.T) {
	err := RegisterContextMenu()
	if err == nil {
		t.Error("RegisterContextMenu() should return an error on non-Windows platforms")
	}
}

func TestUnregisterContextMenu_Stub(t *testing.T) {
	err := UnregisterContextMenu()
	if err == nil {
		t.Error("UnregisterContextMenu() should return an error on non-Windows platforms")
	}
}
