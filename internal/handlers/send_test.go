package handlers

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestServeFileWithProgress(t *testing.T) {
	// Create a dummy file
	content := []byte("Hello, world!")
	tmpfile, err := os.CreateTemp("", "example")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(content); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	// Create a request
	req, err := http.NewRequest("GET", "/download", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder
	rr := httptest.NewRecorder()

	// Call the handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ServeFileWithProgress(w, r, tmpfile.Name())
	})

	handler.ServeHTTP(rr, req)

	// Check status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check body
	if rr.Body.String() != string(content) {
	        t.Errorf("handler returned unexpected body: got %v want %v",
	                rr.Body.String(), string(content))
	}

	t.Run("Range Request", func(t *testing.T) {
	req, _ := http.NewRequest("GET", "/download", nil)
	req.Header.Set("Range", "bytes=0-4") // Should get "Hello"
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusPartialContent {
	t.Errorf("expected 206 Partial Content, got %d", rr.Code)
	}
	if rr.Body.String() != "Hello" {
	t.Errorf("expected 'Hello', got %q", rr.Body.String())
	}
	})
	}
func TestServeDirWithProgress(t *testing.T) {
	// Create a dummy dir
	tmpdir, err := os.MkdirTemp("", "example_dir")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpdir)

	// Create a file inside
	content := []byte("Hello, zip!")
	err = os.WriteFile(filepath.Join(tmpdir, "hello.txt"), content, 0644)
	if err != nil {
		t.Fatal(err)
	}

	// Request
	req, err := http.NewRequest("GET", "/download", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ServeDirWithProgress(w, r, tmpdir)
	})

	handler.ServeHTTP(rr, req)

	// Check status
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check content type
	if ctype := rr.Header().Get("Content-Type"); ctype != "application/zip" {
		t.Errorf("content type mismatch: got %v want application/zip", ctype)
	}

	// Check body size > 0
	if rr.Body.Len() == 0 {
		t.Errorf("response body is empty")
	}
}
