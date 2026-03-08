package ui

import (
	"errors"
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
	Demo   bool
}

// RunInteractivePrompts launches the interactive TUI for MagShare.
func RunInteractivePrompts(demo bool, defaultPort int, defaultPath string, defaultSecure bool) (*InteractiveConfig, error) {
	if defaultPort == 0 {
		defaultPort = 8080
	}

	cfg := &InteractiveConfig{
		Port:   defaultPort,
		Secure: defaultSecure,
		Path:   defaultPath,
		Demo:   demo,
	}

	// 1. Choose Action
	actionForm := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("MagShare Interactive").
				Description("Choose what you want to do today.").
				Options(
					huh.NewOption("Send File/Folder", "send"),
					huh.NewOption("Receive File", "receive"),
				).
				Value(&cfg.Action),
		),
	).WithTheme(huh.ThemeCharm())

	err := actionForm.Run()
	if err != nil {
		if errors.Is(err, huh.ErrUserAborted) {
			return nil, fmt.Errorf("user cancelled")
		}
		return nil, fmt.Errorf("action selection failed: %w", err)
	}

	// 2. Configure Action
	var portStr string = strconv.Itoa(cfg.Port)
	var groups []*huh.Group

	switch cfg.Action {
	case "send":
		groups = append(groups, huh.NewGroup(
			huh.NewInput().
				Title("Path").
				Description("Enter the file or directory you want to share.").
				Placeholder("./my-data").
				Value(&cfg.Path).
				Validate(func(s string) error {
					if s == "" {
						return fmt.Errorf("path cannot be empty")
					}
					return nil
				}),
			huh.NewInput().
				Title("Port").
				Description("The port to run the ephemeral server on.").
				Value(&portStr).
				Validate(func(s string) error {
					_, err := strconv.Atoi(s)
					if err != nil {
						return fmt.Errorf("must be a number")
					}
					return nil
				}),
			huh.NewConfirm().
				Title("Secure Mode").
				Description("Require a 4-digit PIN for access?").
				Value(&cfg.Secure),
		))
	case "receive":
		groups = append(groups, huh.NewGroup(
			huh.NewInput().
				Title("Destination").
				Description("Where should received files be saved?").
				Placeholder("./downloads").
				Value(&cfg.Path).
				Validate(func(s string) error {
					if s == "" {
						return fmt.Errorf("destination cannot be empty")
					}
					return nil
				}),
			huh.NewInput().
				Title("Port").
				Description("The port to run the dropzone server on.").
				Value(&portStr).
				Validate(func(s string) error {
					_, err := strconv.Atoi(s)
					if err != nil {
						return fmt.Errorf("must be a number")
					}
					return nil
				}),
			huh.NewConfirm().
				Title("Secure Mode").
				Description("Require a 4-digit PIN for access?").
				Value(&cfg.Secure),
		))
	}

	configForm := huh.NewForm(groups...).WithTheme(huh.ThemeCharm())
	err = configForm.Run()
	if err != nil {
		if errors.Is(err, huh.ErrUserAborted) {
			return nil, fmt.Errorf("user cancelled")
		}
		return nil, fmt.Errorf("configuration failed: %w", err)
	}

	// Optional PIN prompt if Secure is true
	if cfg.Secure {
		pinForm := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Title("Access PIN").
					Description("Set a 4-digit numeric PIN.").
					Placeholder("1234").
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
		).WithTheme(huh.ThemeCharm())
		err = pinForm.Run()
		if err != nil {
			if errors.Is(err, huh.ErrUserAborted) {
				return nil, fmt.Errorf("user cancelled")
			}
			return nil, fmt.Errorf("PIN entry failed: %w", err)
		}
	}

	cfg.Port, _ = strconv.Atoi(portStr)

	return cfg, nil
}
