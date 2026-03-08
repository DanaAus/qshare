package handlers

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestReceiveFileWithProgress(t *testing.T) {
	// Create a temp dir for upload destination
	tmpdir, err := os.MkdirTemp("", "upload_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpdir)

	// Create a buffer to hold the multipart data
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", "test_upload.txt")
	if err != nil {
		t.Fatal(err)
	}
	part.Write([]byte("Content of the uploaded file"))
	writer.Close()

	req, err := http.NewRequest("POST", "/upload", body)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ReceiveFileWithProgress(w, r, tmpdir, false, "")
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	// Verify file exists
	if _, err := os.Stat(filepath.Join(tmpdir, "test_upload.txt")); os.IsNotExist(err) {
		t.Errorf("Uploaded file was not saved")
	}
}
