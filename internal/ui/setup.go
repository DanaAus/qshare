package ui

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/charmbracelet/huh"
)

// SetupResult holds the configuration collected during the first-run setup.
type SetupResult struct {
	DownloadDir      string
	SecureMode       bool
	ShellIntegration bool
}

// RunFirstRunSetup guides the user through the initial configuration of Magshare.
func RunFirstRunSetup() (*SetupResult, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("could not detect home directory: %w", err)
	}

	defaultDownloadDir := filepath.Join(home, "Magshare Downloads")
	result := &SetupResult{
		DownloadDir:      defaultDownloadDir,
		SecureMode:       false,
		ShellIntegration: true, // Default to true on Windows
	}

	fields := []huh.Field{
		huh.NewInput().
			Title("Download Location").
			Description("Where should received files be saved?").
			Placeholder(defaultDownloadDir).
			Value(&result.DownloadDir).
			Validate(func(s string) error {
				if s == "" {
					return fmt.Errorf("path cannot be empty")
				}
				if !filepath.IsAbs(s) {
					return fmt.Errorf("path must be absolute")
				}
				return nil
			}),
		huh.NewConfirm().
			Title("Default Security").
			Description("Enable PIN security by default for all future transfers?").
			Value(&result.SecureMode),
	}

	// Add Windows-specific integration prompt
	if runtime.GOOS == "windows" {
		fields = append(fields,
			huh.NewConfirm().
				Title("Explorer Integration").
				Description("Enable Windows Explorer integration? (Adds 'Share via Magshare' to right-click menu)").
				Value(&result.ShellIntegration),
		)
	}

	form := huh.NewForm(huh.NewGroup(fields...)).WithTheme(huh.ThemeCharm())

	err = form.Run()
	if err != nil {
		if errors.Is(err, huh.ErrUserAborted) {
			return nil, fmt.Errorf("setup cancelled by user")
		}
		return nil, fmt.Errorf("setup failed: %w", err)
	}

	return result, nil
}
