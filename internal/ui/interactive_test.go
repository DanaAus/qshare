package ui

import (
	"testing"
)

func TestInteractiveConfigStructure(t *testing.T) {
	cfg := &InteractiveConfig{
		Action: "send",
		Path:   "./test",
		Port:   8080,
		PIN:    "1234",
		Secure: true,
	}

	if cfg.Action != "send" {
		t.Errorf("Expected Action send, got %s", cfg.Action)
	}
}
