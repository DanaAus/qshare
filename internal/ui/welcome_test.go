package ui

import (
	"bytes"
	"strings"
	"testing"
)

func TestDisplayWelcomeMessage(t *testing.T) {
	var buf bytes.Buffer
	DisplayWelcomeMessage(&buf)

	output := buf.String()
	if output == "" {
		t.Error("DisplayWelcomeMessage() produced no output")
	}

	if !strings.Contains(output, "Welcome") {
		t.Errorf("DisplayWelcomeMessage() output %q does not contain 'Welcome'", output)
	}
	
	if !strings.Contains(output, "magshare") {
		t.Errorf("DisplayWelcomeMessage() output %q does not contain 'magshare'", output)
	}
}
