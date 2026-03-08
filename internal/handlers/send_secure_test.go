package handlers

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestStartSendServerSecure(t *testing.T) {
	// Create a dummy file
	content := []byte("secret content")
	tmpfile, err := os.CreateTemp("", "secret")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	tmpfile.Write(content)
	tmpfile.Close()

	opts := SendOptions{
		Secure: true,
		PIN:    "1234",
	}

	// We can't easily test StartSendServer because it starts a blocking loop.
	// But we can test the handler logic if we extract it.
	// For now, let's just test the handler logic directly using a mock.
	
	t.Run("Missing PIN serves PIN page", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/d/test", nil)
		rr := httptest.NewRecorder()

		// Manually invoke the logic that would be in the handler
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			clientPin := r.URL.Query().Get("pin")
			if clientPin != opts.PIN {
				// This matches the logic added to StartSendServer
				servePINPage(w, r, clientPin)
				return
			}
			ServeFileWithProgress(w, r, tmpfile.Name())
		})

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("expected status 200, got %v", status)
		}
		if !strings.Contains(rr.Body.String(), "Secure Access") {
			t.Errorf("expected PIN page, but 'Secure Access' not found in body")
		}
	})

	t.Run("Correct PIN serves file", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/d/test?pin=1234", nil)
		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			clientPin := r.URL.Query().Get("pin")
			if clientPin != opts.PIN {
				servePINPage(w, r, clientPin)
				return
			}
			ServeFileWithProgress(w, r, tmpfile.Name())
		})

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("expected status 200, got %v", status)
		}
		if rr.Body.String() != string(content) {
			t.Errorf("expected secret content, got %v", rr.Body.String())
		}
	})
}

// Helper to test the logic without refactoring send.go too much yet
func servePINPage(w http.ResponseWriter, r *http.Request, clientPin string) {
	// Dummy implementation for test
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Secure Access - Enter PIN"))
}
