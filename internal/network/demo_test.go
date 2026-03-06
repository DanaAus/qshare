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

	// Test demo mode
	gotDemo := GetDisplayURL(realURL, true)
	expectedDemo := "http://192.168.100.100:8080/u/abc"
	if gotDemo != expectedDemo {
		t.Errorf("expected %s, got %s", expectedDemo, gotDemo)
	}
}
