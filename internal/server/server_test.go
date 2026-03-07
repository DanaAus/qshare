package server

import (
	"net/http"
	"testing"
	"time"
)

func TestEphemeralServer(t *testing.T) {
	s := NewEphemeralServer(0) // dynamic port
	
	s.Handle("/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Run in goroutine
	errChan := make(chan error)
	go func() {
		errChan <- s.Start(100 * time.Millisecond)
	}()

	// Wait for server to start (simple sleep for test)
	time.Sleep(10 * time.Millisecond)

	// Trigger shutdown manually
	s.TriggerShutdown()

	err := <-errChan
	if err != nil {
		t.Errorf("Server start/shutdown failed: %v", err)
	}
}
