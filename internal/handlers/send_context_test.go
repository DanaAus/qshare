package handlers

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestServeFileContextCancellation(t *testing.T) {
	// Create a temporary file
	content := make([]byte, 1024*1024) // 1MB
	tmpfile, err := os.CreateTemp("", "context_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	tmpfile.Write(content)
	tmpfile.Close()

	// Handler that uses ServeFileWithProgress
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := ServeFileWithProgress(w, r, tmpfile.Name())
		if err != nil {
			t.Logf("Handler finished with error: %v", err)
		}
	})

	ts := httptest.NewServer(handler)
	defer ts.Close()

	// Create a request with a short timeout/cancellation
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	req, _ := http.NewRequestWithContext(ctx, "GET", ts.URL, nil)
	client := &http.Client{}
	
	resp, err := client.Do(req)
	if err == nil {
		// If we got a response, try to read it until context is cancelled
		_, _ = io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
	}
	
	// We want to verify that the server handler stopped. 
	// This test is mostly a sanity check for now.
}
