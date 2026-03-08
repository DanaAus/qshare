package handlers

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestServeFileConcurrency(t *testing.T) {
	// 1. Create a large temporary file (100MB)
	fileName := "concurrency_test_file.bin"
	file, err := os.Create(fileName)
	if err != nil {
		t.Fatalf("Failed to create large file: %v", err)
	}
	defer os.Remove(fileName)

	chunk := make([]byte, 1024*1024) // 1MB chunk
	for i := 0; i < 100; i++ {
		file.Write(chunk)
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

	// 3. Start 5 Concurrent Downloads
	var wg sync.WaitGroup
	numConcurrent := 5
	wg.Add(numConcurrent)

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

	for i := 0; i < numConcurrent; i++ {
		go func(id int) {
			defer wg.Done()
			resp, err := http.Get(ts.URL)
			if err != nil {
				t.Errorf("Goroutine %d failed to GET: %v", id, err)
				return
			}
			defer resp.Body.Close()
			_, _ = io.Copy(io.Discard, resp.Body)
		}(i)
	}

	wg.Wait()
	close(done)

	// 4. Assert Memory Usage
	threshold := uint64(50 * 1024 * 1024) // 50MB (buffer for test environment)
	fmt.Printf("Peak Memory Allocation (Concurrent): %d MB\n", maxAlloc/1024/1024)

	if maxAlloc > threshold {
		t.Errorf("Peak memory allocation exceeded threshold: got %d MB, want < %d MB", maxAlloc/1024/1024, threshold/1024/1024)
	}
}
