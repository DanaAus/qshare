package handlers

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestServeFileRangeSupport(t *testing.T) {
	// Create a dummy file
	content := []byte("0123456789") // 10 bytes
	tmpfile, err := os.CreateTemp("", "range_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write(content); err != nil {
		t.Fatal(err)
	}
	tmpfile.Close()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ServeFileWithProgress(w, r, tmpfile.Name())
	})

	cases := []struct {
		name           string
		rangeHeader    string
		expectedStatus int
		expectedBody   string
	}{
		{"First 5 bytes", "bytes=0-4", http.StatusPartialContent, "01234"},
		{"Last 5 bytes", "bytes=5-9", http.StatusPartialContent, "56789"},
		{"Middle 2 bytes", "bytes=4-5", http.StatusPartialContent, "45"},
		{"Suffix 3 bytes", "bytes=-3", http.StatusPartialContent, "789"},
		{"Full range", "bytes=0-9", http.StatusPartialContent, "0123456789"},
		{"No range", "", http.StatusOK, "0123456789"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/download", nil)
			if tc.rangeHeader != "" {
				req.Header.Set("Range", tc.rangeHeader)
			}
			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)

			if rr.Code != tc.expectedStatus {
				t.Errorf("expected status %d, got %d", tc.expectedStatus, rr.Code)
			}
			if rr.Body.String() != tc.expectedBody {
				t.Errorf("expected body %q, got %q", tc.expectedBody, rr.Body.String())
			}
		})
	}
}
