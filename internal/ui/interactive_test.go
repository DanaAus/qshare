package ui

import (
	"testing"
)

func TestMainActionForm(t *testing.T) {
	var action string
	form := MainActionForm(&action)

	if form == nil {
		t.Fatal("Expected MainActionForm to return a form, got nil")
	}

	// We'll perform a basic check that the form has at least one group.
	// Since huh doesn't expose its internal group structure easily,
	// this is just to confirm wiring.
}
