package ui

import (
	"fmt"
	"io"
	"github.com/charmbracelet/lipgloss"
)

// DisplayWelcomeMessage prints a styled welcome message to the provided writer.
func DisplayWelcomeMessage(w io.Writer) {
	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		Padding(0, 1).
		MarginBottom(1)

	descStyle := lipgloss.NewStyle().
		Italic(true).
		Foreground(lipgloss.Color("#888888"))

	welcomeText := titleStyle.Render("Welcome to magshare!")
	description := descStyle.Render("Instant, frictionless file sharing across your local network.")

	fmt.Fprintln(w, welcomeText)
	fmt.Fprintln(w, description)
	fmt.Fprintln(w)
}
