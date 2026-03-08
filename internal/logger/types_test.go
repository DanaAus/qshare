package logger

import (
	"bytes"
	"strings"
	"testing"
)

func TestLogLevelComparison(t *testing.T) {
	if DEBUG >= INFO {
		t.Error("DEBUG should be less than INFO")
	}
	if INFO >= WARN {
		t.Error("INFO should be less than WARN")
	}
	if WARN >= ERROR {
		t.Error("WARN should be less than ERROR")
	}
}

func TestStructuredLoggerCreation(t *testing.T) {
	var buf bytes.Buffer
	l := &StructuredLogger{
		Writer:    &buf,
		Component: "test",
		PID:       1234,
	}

	if l.Component != "test" {
		t.Errorf("Expected component 'test', got %s", l.Component)
	}
	if l.PID != 1234 {
		t.Errorf("Expected PID 1234, got %d", l.PID)
	}
}

func TestStructuredLoggerFormatting(t *testing.T) {
	var buf bytes.Buffer
	l := &StructuredLogger{
		Writer:    &buf,
		Component: "test",
		PID:       1234,
	}

	msg := "Started listening on :8080"
	l.Info(msg)

	output := buf.String()
	// Format: [TIMESTAMP] [LEVEL] [COMPONENT] [PID] MESSAGE
	// Example: [2023-10-27 14:30:05] [INFO] [server] [1234] Started listening on :8080
	
	if !strings.Contains(output, "[INFO]") {
		t.Error("Output missing [INFO] level")
	}
	if !strings.Contains(output, "[test]") {
		t.Error("Output missing [test] component")
	}
	if !strings.Contains(output, "[1234]") {
		t.Error("Output missing [1234] PID")
	}
	if !strings.Contains(output, msg) {
		t.Error("Output missing original message")
	}
}
