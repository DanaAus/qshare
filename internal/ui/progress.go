package ui

import (
	"github.com/schollz/progressbar/v3"
)

// NewProgressBar creates a standard progress bar for file transfers.
// max is the total number of bytes.
// description is the label shown next to the bar.
func NewProgressBar(max int64, description string) *progressbar.ProgressBar {
	return progressbar.DefaultBytes(
		max,
		description,
	)
}

// NewIndeterminateProgressBar creates a spinner for operations with unknown total size.
func NewIndeterminateProgressBar(description string) *progressbar.ProgressBar {
	return progressbar.Default(-1, description)
}
