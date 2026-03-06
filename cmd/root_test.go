package cmd

import (
	"testing"
)

func TestDemoFlag(t *testing.T) {
	flag := rootCmd.PersistentFlags().Lookup("demo")
	if flag == nil {
		t.Fatal("demo flag should be defined")
	}
	if flag.Name != "demo" {
		t.Errorf("expected flag name 'demo', got '%s'", flag.Name)
	}
	if flag.Value.Type() != "bool" {
		t.Errorf("expected flag type 'bool', got '%s'", flag.Value.Type())
	}
}
