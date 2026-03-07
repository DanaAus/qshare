package logger

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

// SetupLogging redirects os.Stdout and os.Stderr to both the terminal and a log file.
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

	// Create pipes
	outR, outW, _ := os.Pipe()
	errR, errW, _ := os.Pipe()

	// Redirect
	os.Stdout = outW
	os.Stderr = errW

	// Start copying in goroutines
	done := make(chan bool)
	go func() {
		mw := io.MultiWriter(originalStdout, logFile)
		io.Copy(mw, outR)
		done <- true
	}()
	go func() {
		mw := io.MultiWriter(originalStderr, logFile)
		io.Copy(mw, errR)
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
