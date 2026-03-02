package handlers

import (
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/schollz/progressbar/v3"
)

func TestProgressReader(t *testing.T) {
	data := []byte("hello, world!")
	src := bytes.NewReader(data)
	bar := progressbar.New(len(data))
	
	// ProgressReader is now context-aware
	reader := NewProgressReader(context.Background(), src, bar)
	
	buf := make([]byte, 5)
	n, err := reader.Read(buf)
	if err != nil {
		t.Fatalf("Read failed: %v", err)
	}
	if n != 5 {
		t.Errorf("Expected 5 bytes, got %d", n)
	}
	
	if bar.State().CurrentNum != 5 {
		t.Errorf("Expected progress bar to be at 5, got %d", bar.State().CurrentNum)
	}
	
	// Read the rest
	_, _ = io.ReadAll(reader)
	if bar.State().CurrentNum != int64(len(data)) {
		t.Errorf("Expected progress bar to be at %d, got %d", len(data), bar.State().CurrentNum)
	}
}
