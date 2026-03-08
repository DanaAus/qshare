package logger

import (
	"bytes"
	"strings"
	"testing"
)

func TestGlobalLogger(t *testing.T) {
	var buf bytes.Buffer
	l := &StructuredLogger{
		Writer:    &buf,
		Component: "global-test",
		PID:       1234,
	}
	SetGlobalLogger(l)

	Info("testing global info")
	if !strings.Contains(buf.String(), "[INFO]") || !strings.Contains(buf.String(), "testing global info") {
		t.Errorf("Global Info failed, got %q", buf.String())
	}

	buf.Reset()
	Debug("testing global debug")
	if !strings.Contains(buf.String(), "[DEBUG]") || !strings.Contains(buf.String(), "testing global debug") {
		t.Errorf("Global Debug failed, got %q", buf.String())
	}
}

func TestWithComponent(t *testing.T) {
	var buf bytes.Buffer
	l := &StructuredLogger{
		Writer:    &buf,
		Component: "base",
		PID:       1234,
	}
	SetGlobalLogger(l)

	compLogger := WithComponent("child")
	compLogger.Info("child message")

	if !strings.Contains(buf.String(), "[child]") {
		t.Errorf("Component logger failed, expected [child] in %q", buf.String())
	}
}
