package handlers

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"runtime"
	"testing"
	"time"
)

func TestServeFileMemoryUsage(t *testing.T) {
	// 1. Create a large temporary file (500MB)
	fileName := "large_test_file.bin"
	file, err := os.Create(fileName)
	if err != nil {
		t.Fatalf("Failed to create large file: %v", err)
	}
	defer os.Remove(fileName)

	// Fill with 500MB of data
	chunk := make([]byte, 1024*1024) // 1MB chunk
	for i := 0; i < 500; i++ {
		if _, err := file.Write(chunk); err != nil {
			t.Fatalf("Failed to write to large file: %v", err)
		}
	}
	file.Close()

	// 2. Setup Test Server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := ServeFileWithProgress(w, r, fileName)
		if err != nil {
			t.Errorf("ServeFileWithProgress failed: %v", err)
		}
	}))
	defer ts.Close()

	// 3. Track Memory Usage in a Goroutine
	var maxAlloc uint64
	done := make(chan bool)
	go func() {
		ticker := time.NewTicker(10 * time.Millisecond)
		defer ticker.Stop()
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				var m runtime.MemStats
				runtime.ReadMemStats(&m)
				if m.Alloc > maxAlloc {
					maxAlloc = m.Alloc
				}
			}
		}
	}()

	// 4. Perform Request
	resp, err := http.Get(ts.URL)
	if err != nil {
		t.Fatalf("Failed to GET from test server: %v", err)
	}
	defer resp.Body.Close()

	// Use io.Discard to ensure we read the whole body without storing it
	if _, err := io.Copy(io.Discard, resp.Body); err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	done <- true

	// 5. Assert Memory Usage
	// Max allocation should be significantly less than 500MB.
	// We'll set a threshold of 50MB for the test environment overhead.
	threshold := uint64(50 * 1024 * 1024) // 50MB
	fmt.Printf("Peak Memory Allocation: %d MB\n", maxAlloc/1024/1024)

	if maxAlloc > threshold {
		t.Errorf("Peak memory allocation exceeded threshold: got %d MB, want < %d MB", maxAlloc/1024/1024, threshold/1024/1024)
	}
}
