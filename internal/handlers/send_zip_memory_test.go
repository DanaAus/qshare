package handlers

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"time"
)

func TestServeDirZipMemoryUsage(t *testing.T) {
	// 1. Create a directory with many large files
	tmpdir, err := os.MkdirTemp("", "zip_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpdir)

	// Create 5 files of 100MB each
	chunk := make([]byte, 1024*1024) // 1MB chunk
	for i := 0; i < 5; i++ {
		fileName := filepath.Join(tmpdir, fmt.Sprintf("file_%d.bin", i))
		file, _ := os.Create(fileName)
		for j := 0; j < 100; j++ {
			file.Write(chunk)
		}
		file.Close()
	}

	// 2. Setup Test Server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := ServeDirWithProgress(w, r, tmpdir)
		if err != nil {
			t.Errorf("ServeDirWithProgress failed: %v", err)
		}
	}))
	defer ts.Close()

	// 3. Track Memory Usage
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
		t.Fatalf("Failed to GET: %v", err)
	}
	defer resp.Body.Close()
	_, _ = io.Copy(io.Discard, resp.Body)

	done <- true

	// 5. Assert Memory Usage
	threshold := uint64(50 * 1024 * 1024) // 50MB
	fmt.Printf("Peak Memory Allocation (ZIP): %d MB\n", maxAlloc/1024/1024)

	if maxAlloc > threshold {
		t.Errorf("Peak memory allocation exceeded threshold: got %d MB, want < %d MB", maxAlloc/1024/1024, threshold/1024/1024)
	}
}
