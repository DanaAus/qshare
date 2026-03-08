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
                Demo:   true,
        }

        if cfg.Action != "send" {
                t.Errorf("Expected Action send, got %s", cfg.Action)
        }
}

func TestInteractiveSignature(t *testing.T) {
	// Just check if we can call it (won't run in tests as it needs TTY)
	// But we can check if it compiles and the types are correct.
	_ = RunInteractivePrompts
}
