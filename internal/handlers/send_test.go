package handlers

import (
	"net/http"
	"net/http/httptest"
	"os"
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
}
