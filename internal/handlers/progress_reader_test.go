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

	func TestProgressReadSeeker(t *testing.T) {
	data := []byte("0123456789")
	src := bytes.NewReader(data)
	bar := progressbar.New(len(data))

	prs := NewProgressReadSeeker(context.Background(), src, bar)

	// Read 2 bytes
	buf := make([]byte, 2)
	n, _ := prs.Read(buf)
	if n != 2 || string(buf) != "01" {
	t.Errorf("Expected '01', got %q", string(buf))
	}

	// Seek to end
	pos, err := prs.Seek(0, io.SeekEnd)
	if err != nil || pos != 10 {
	t.Errorf("SeekEnd failed: pos=%d, err=%v", pos, err)
	}

	// Seek to start
	pos, err = prs.Seek(0, io.SeekStart)
	if err != nil || pos != 0 {
	t.Errorf("SeekStart failed: pos=%d, err=%v", pos, err)
	}

	// Read again
	n, _ = prs.Read(buf)
	if n != 2 || string(buf) != "01" {
	t.Errorf("Read after seek failed: %q", string(buf))
	}
	}
