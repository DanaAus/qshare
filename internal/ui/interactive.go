package ui

import (
	"fmt"
	"strconv"

	"github.com/charmbracelet/huh"
)

// InteractiveConfig holds the values collected from the interactive prompts.
type InteractiveConfig struct {
	Action string
	Path   string
	Port   int
	PIN    string
	Secure bool
}

// RunInteractivePrompts launches the interactive TUI for QShare.
func RunInteractivePrompts() (*InteractiveConfig, error) {
	cfg := &InteractiveConfig{
		Port: 8080, // Default port
	}

	// 1. Choose Action
	actionForm := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("What would you like to do?").
				Options(
					huh.NewOption("Send File/Folder", "send"),
					huh.NewOption("Receive File", "receive"),
					huh.NewOption("Sync Mode", "sync"),
				).
				Value(&cfg.Action),
		),
	)

	err := actionForm.Run()
	if err != nil {
		return nil, fmt.Errorf("action selection failed: %w", err)
	}

	// 2. Configure Action
	var portStr string = "8080"
	var groups []*huh.Group

	switch cfg.Action {
	case "send", "sync":
		groups = append(groups, huh.NewGroup(
			huh.NewInput().
				Title("File or Directory Path").
				Placeholder("e.g., ./my-file.txt").
				Value(&cfg.Path).
				Validate(func(s string) error {
					if s == "" {
						return fmt.Errorf("path cannot be empty")
					}
					return nil
				}),
			huh.NewInput().
				Title("Port").
				Value(&portStr).
				Validate(func(s string) error {
					_, err := strconv.Atoi(s)
					if err != nil {
						return fmt.Errorf("must be a number")
					}
					return nil
				}),
			huh.NewConfirm().
				Title("Use Secure Mode (PIN)?").
				Value(&cfg.Secure),
		))
	case "receive":
		groups = append(groups, huh.NewGroup(
			huh.NewInput().
				Title("Destination Directory").
				Placeholder("e.g., ./downloads").
				Value(&cfg.Path).
				Validate(func(s string) error {
					if s == "" {
						return fmt.Errorf("destination cannot be empty")
					}
					return nil
				}),
			huh.NewInput().
				Title("Port").
				Value(&portStr).
				Validate(func(s string) error {
					_, err := strconv.Atoi(s)
					if err != nil {
						return fmt.Errorf("must be a number")
					}
					return nil
				}),
		))
	}

	configForm := huh.NewForm(groups...)
	err = configForm.Run()
	if err != nil {
		return nil, fmt.Errorf("configuration failed: %w", err)
	}

	// Optional PIN prompt if Secure is true
	if cfg.Secure && (cfg.Action == "send" || cfg.Action == "sync") {
		pinForm := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Title("Set a 4-digit PIN").
					Placeholder("e.g., 1234").
					Value(&cfg.PIN).
					Validate(func(s string) error {
						if len(s) != 4 {
							return fmt.Errorf("PIN must be 4 digits")
						}
						_, err := strconv.Atoi(s)
						if err != nil {
							return fmt.Errorf("PIN must be numeric")
						}
						return nil
					}),
			),
		)
		err = pinForm.Run()
		if err != nil {
			return nil, fmt.Errorf("PIN entry failed: %w", err)
		}
	}

	cfg.Port, _ = strconv.Atoi(portStr)

	return cfg, nil
}
