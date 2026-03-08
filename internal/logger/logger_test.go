package logger

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestSetupLogging(t *testing.T) {
	tmpdir, err := os.MkdirTemp("", "logger_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpdir)

	// Mock os.Stdout/Stderr would be complex here. 
	// Let's at least test that it creates a file with the right name format.
	
	logPath, cleanup, err := SetupLogging(tmpdir)
	if err != nil {
		t.Fatalf("SetupLogging returned error: %v", err)
	}
	defer cleanup()

	if logPath == "" {
		t.Error("SetupLogging returned empty logPath")
	}

	// Verify filename format: magshare-YYYYMMDD-HHMMSS.log
	filename := filepath.Base(logPath)
	if !strings.HasPrefix(filename, "magshare-") || !strings.HasSuffix(filename, ".log") {
		t.Errorf("Unexpected log filename format: %s", filename)
	}

	// Verify file exists
	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		t.Errorf("Log file %q was not created", logPath)
	}
}

func TestCleanupLogs(t *testing.T) {
	tmpdir, err := os.MkdirTemp("", "cleanup_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpdir)

	testLog := filepath.Join(tmpdir, "test.log")
	if err := os.WriteFile(testLog, []byte("test"), 0644); err != nil {
		t.Fatal(err)
	}

	if _, err := os.Stat(testLog); os.IsNotExist(err) {
		t.Fatal("Test log file not created")
	}

	if err := CleanupLogs(testLog); err != nil {
		t.Fatalf("CleanupLogs returned error: %v", err)
	}

	if _, err := os.Stat(testLog); !os.IsNotExist(err) {
		t.Error("CleanupLogs did not delete the log file")
	}
}
