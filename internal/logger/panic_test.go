package logger

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestHandlePanic(t *testing.T) {
	tmpdir, err := os.MkdirTemp("", "panic_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpdir)

	logPath := filepath.Join(tmpdir, "crash.log")
	
	// Create an empty log file first
	if err := os.WriteFile(logPath, []byte(""), 0644); err != nil {
		t.Fatal(err)
	}

	// Capture what goes into the global logger
	var buf bytes.Buffer
	originalGlobal := GetGlobalLogger()
	defer SetGlobalLogger(originalGlobal)

	// Set global logger to write to BOTH the buffer and the file
	// HandlePanic will also open the file directly for the stack trace
	f, err := os.OpenFile(logPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		t.Fatal(err)
	}
	
	SetGlobalLogger(&StructuredLogger{
		Writer:    io.MultiWriter(&buf, f),
		Component: "test",
		PID:       1234,
	})
	f.Close() // Close it, HandlePanic opens it independently

	finished := make(chan bool)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				HandlePanic(logPath, r)
				finished <- true
			}
		}()
		panic("test panic")
	}()

	<-finished

	// Verify log file contains the panic message (from global logger Error call)
	// AND the stack trace (from direct file write)
	content, err := os.ReadFile(logPath)
	if err != nil {
		t.Fatal(err)
	}

	contentStr := string(content)
	if !strings.Contains(contentStr, "test panic") {
		t.Errorf("Log file does not contain panic message. Content: %s", contentStr)
	}
	
	if !strings.Contains(contentStr, "STACK TRACE") {
		t.Error("Log file does not contain stack trace")
	}
}
