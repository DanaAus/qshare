package server

import (
	"context"
	"fmt"
	"magshare/internal/logger"
	"net/http"
	"time"
)

// EphemeralServer represents a web server that shuts down automatically.
type EphemeralServer struct {
	Server   *http.Server
	mux      *http.ServeMux
	quitChan chan struct{}
}

// NewEphemeralServer creates a new server instance.
func NewEphemeralServer(port int) *EphemeralServer {
	mux := http.NewServeMux()

	s := &EphemeralServer{
		Server: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: mux,
		},
		mux:      mux,
		quitChan: make(chan struct{}),
	}
	return s
}

// Handle registers a handler for a given pattern.
func (s *EphemeralServer) Handle(pattern string, handler http.HandlerFunc) {
	s.mux.HandleFunc(pattern, handler)
}

// Start runs the server and handles the timeout logic.
func (s *EphemeralServer) Start(timeout time.Duration) error {
	l := logger.WithComponent("server")
	l.Debug(fmt.Sprintf("Starting server on %s", s.Server.Addr))

	go func() {
		if err := s.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			l.Error(fmt.Sprintf("Server error: %v", err))
		}
	}()

	// Wait for either shutdown signal or timeout
	select {
	case <-s.quitChan:
		// Normal shutdown triggered by a handler
		l.Info("Transfer complete. Shutting down server...")
	case <-time.After(timeout):
		// Timeout reached
		l.Warn("Server timed out (inactivity). Shutting down...")
	}

	return s.Shutdown()
}

// Shutdown stops the server gracefully.
func (s *EphemeralServer) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return s.Server.Shutdown(ctx)
}

// TriggerShutdown can be called by handlers to initiate server shutdown.
func (s *EphemeralServer) TriggerShutdown() {
	// Use non-blocking send in case it's called multiple times
	select {
	case s.quitChan <- struct{}{}:
	default:
	}
}
