package logger

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

// SetupLogging redirects os.Stdout and os.Stderr to both the terminal and a log file.
// It also configures the global structured logger.
// It returns the path to the log file and a cleanup function to restore originals.
func SetupLogging(logsDir string) (string, func(), error) {
	// Generate filename
	timestamp := time.Now().Format("20060102-150405")
	logPath := filepath.Join(logsDir, fmt.Sprintf("magshare-%s.log", timestamp))

	// Open log file
	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return "", nil, err
	}

	// Save original stdout/stderr
	originalStdout := os.Stdout
	originalStderr := os.Stderr

	// Create filtered writers for terminal (INFO+) and file (DEBUG+)
	termWriter := &FilteredWriter{Writer: originalStdout, Threshold: INFO}
	fileWriter := &FilteredWriter{Writer: logFile, Threshold: DEBUG}

	// For stderr, we might want to always show errors in terminal
	termErrWriter := &FilteredWriter{Writer: originalStderr, Threshold: INFO}

	// Create MultiLeveledWriters for internal use
	outMLW := &MultiLeveledWriter{Writers: []LeveledWriter{termWriter, fileWriter}}
	errMLW := &MultiLeveledWriter{Writers: []LeveledWriter{termErrWriter, fileWriter}}

	// Configure Global Structured Logger
	SetGlobalLogger(&StructuredLogger{
		Writer:    outMLW,
		Component: "main",
		PID:       os.Getpid(),
	})

	// Setup Pipe redirection for raw fmt.Print and third-party libs
	outR, outW, _ := os.Pipe()
	errR, errW, _ := os.Pipe()

	// Redirect
	os.Stdout = outW
	os.Stderr = errW

	// Start copying in goroutines
	done := make(chan bool)
	go func() {
		// Raw writes to stdout are treated as INFO
		io.Copy(outMLW, outR)
		done <- true
	}()
	go func() {
		// Raw writes to stderr are treated as INFO
		io.Copy(errMLW, errR)
		done <- true
	}()

	cleanup := func() {
		// Restore original stdout/stderr
		os.Stdout = originalStdout
		os.Stderr = originalStderr

		// Close write ends of pipes to trigger EOF in goroutines
		outW.Close()
		errW.Close()

		// Wait for goroutines to finish
		<-done
		<-done

		// Close log file
		logFile.Close()
	}

	return logPath, cleanup, nil
}

// CleanupLogs deletes the log file at the specified path.
func CleanupLogs(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil
	}
	return os.Remove(path)
}
