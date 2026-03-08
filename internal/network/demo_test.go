package network

import (
	"testing"
)

func TestGetDisplayIP(t *testing.T) {
	realIP := "192.168.1.5"
	
	// Test normal mode
	gotNormal := GetDisplayIP(realIP, false)
	if gotNormal != realIP {
		t.Errorf("expected %s, got %s", realIP, gotNormal)
	}

	// Test demo mode
	gotDemo := GetDisplayIP(realIP, true)
	expectedDemo := "192.168.100.100"
	if gotDemo != expectedDemo {
		t.Errorf("expected %s, got %s", expectedDemo, gotDemo)
	}
}

func TestGetDisplayURL(t *testing.T) {
	realURL := "http://192.168.1.5:54321/u/abc"
	
	// Test normal mode
	gotNormal := GetDisplayURL(realURL, false)
	if gotNormal != realURL {
		t.Errorf("expected %s, got %s", realURL, gotNormal)
	}

	// Test demo mode (preserves port)
	gotDemo := GetDisplayURL(realURL, true)
	expectedDemo := "http://192.168.100.100:54321/u/abc"
	if gotDemo != expectedDemo {
		t.Errorf("expected %s, got %s", expectedDemo, gotDemo)
	}
}

func TestGetDisplayURL_CustomPort(t *testing.T) {
	realURL := "http://192.168.1.5:9090/u/abc"
	
	// Test demo mode with custom port
	// This is expected to FAIL with current implementation which hardcodes 8080
	gotDemo := GetDisplayURL(realURL, true)
	expectedDemo := "http://192.168.100.100:9090/u/abc"
	if gotDemo != expectedDemo {
		t.Errorf("expected %s, got %s", expectedDemo, gotDemo)
	}
}
