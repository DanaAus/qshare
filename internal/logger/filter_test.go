package logger

import (
	"bytes"
	"strings"
	"testing"
)

func TestFilteredWriter(t *testing.T) {
	var buf bytes.Buffer
	fw := &FilteredWriter{
		Writer:    &buf,
		Threshold: INFO,
	}

	// Should not write DEBUG
	fw.WriteLevel(DEBUG, []byte("debug message\n"))
	if buf.Len() > 0 {
		t.Errorf("Expected nothing to be written for DEBUG, got %q", buf.String())
	}

	// Should write INFO
	msg := "info message\n"
	fw.WriteLevel(INFO, []byte(msg))
	if !strings.Contains(buf.String(), msg) {
		t.Errorf("Expected %q to be written for INFO, got %q", msg, buf.String())
	}

	// Should write ERROR
	buf.Reset()
	msgErr := "error message\n"
	fw.WriteLevel(ERROR, []byte(msgErr))
	if !strings.Contains(buf.String(), msgErr) {
		t.Errorf("Expected %q to be written for ERROR, got %q", msgErr, buf.String())
	}
}

func TestMultiLeveledWriter(t *testing.T) {
	var buf1, buf2 bytes.Buffer
	fw1 := &FilteredWriter{Writer: &buf1, Threshold: INFO}
	fw2 := &FilteredWriter{Writer: &buf2, Threshold: DEBUG}

	mlw := &MultiLeveledWriter{
		Writers: []LeveledWriter{fw1, fw2},
	}

	msg := "debug message\n"
	mlw.WriteLevel(DEBUG, []byte(msg))

	if buf1.Len() > 0 {
		t.Errorf("buf1 (Threshold INFO) should be empty for DEBUG, got %q", buf1.String())
	}
	if !strings.Contains(buf2.String(), msg) {
		t.Errorf("buf2 (Threshold DEBUG) should contain %q, got %q", msg, buf2.String())
	}
}
