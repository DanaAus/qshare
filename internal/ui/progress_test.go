package ui_test

import (
	"testing"
	"magshare/internal/ui"
)

func TestNewProgressBar(t *testing.T) {
	bar := ui.NewProgressBar(100, "Sending file")
	if bar == nil {
		t.Error("NewProgressBar returned nil")
	}
}

func TestNewIndeterminateProgressBar(t *testing.T) {
	bar := ui.NewIndeterminateProgressBar("Compressing folder")
	if bar == nil {
		t.Error("NewIndeterminateProgressBar returned nil")
	}
}
