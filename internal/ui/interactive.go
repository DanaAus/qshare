package ui

import (
	"fmt"
	"github.com/charmbracelet/huh"
)

// MainActionForm returns a huh.Form for choosing the primary task.
func MainActionForm(action *string) *huh.Form {
	return huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("What would you like to do?").
				Options(
					huh.NewOption("Send", "send"),
					huh.NewOption("Receive", "receive"),
					huh.NewOption("Sync", "sync"),
				).
				Value(action),
		),
	)
}

// RunInteractivePrompts launches the interactive TUI for QShare.
func RunInteractivePrompts() error {
	var action string
	form := MainActionForm(&action)

	err := form.Run()
	if err != nil {
		return fmt.Errorf("form failed: %w", err)
	}

	fmt.Printf("You chose: %s\n", action)
	return nil
}
