package workspace

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestCreateDefaultConfig(t *testing.T) {
	tmpdir, err := os.MkdirTemp("", "config_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpdir)

	configPath := filepath.Join(tmpdir, "config.json")

	expectedConfig := Config{
		Port:        9090,
		SecureMode:  true,
		DownloadDir: tmpdir, // Use existing dir for validation
	}

	err = CreateDefaultConfig(configPath, expectedConfig)
	if err != nil {
		t.Fatalf("CreateDefaultConfig returned error: %v", err)
	}

	// Verify file exists
	if !FileExists(configPath) {
		t.Error("Config file was not created")
	}

	// Verify content
	data, err := os.ReadFile(configPath)
	if err != nil {
		t.Fatal(err)
	}

	var actualConfig Config
	if err := json.Unmarshal(data, &actualConfig); err != nil {
		t.Fatalf("Failed to unmarshal config: %v", err)
	}

	// Check values
	if actualConfig.Port != expectedConfig.Port {
		t.Errorf("Port = %d, want %d", actualConfig.Port, expectedConfig.Port)
	}
	if actualConfig.SecureMode != expectedConfig.SecureMode {
		t.Errorf("SecureMode = %v, want %v", actualConfig.SecureMode, expectedConfig.SecureMode)
	}
	if actualConfig.DownloadDir != expectedConfig.DownloadDir {
		t.Errorf("DownloadDir = %q, want %q", actualConfig.DownloadDir, expectedConfig.DownloadDir)
	}
}

func TestLoadConfig(t *testing.T) {
	tmpdir, err := os.MkdirTemp("", "load_config_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpdir)

	t.Run("Valid Config", func(t *testing.T) {
		configPath := filepath.Join(tmpdir, "valid_config.json")
		expected := Config{
			Port:        8080,
			SecureMode:  false,
			DownloadDir: tmpdir,
		}
		data, _ := json.Marshal(expected)
		os.WriteFile(configPath, data, 0644)

		cfg, err := LoadConfig(configPath)
		if err != nil {
			t.Fatalf("LoadConfig failed: %v", err)
		}
		if cfg.Port != expected.Port {
			t.Errorf("got port %d, want %d", cfg.Port, expected.Port)
		}
	})

	t.Run("Missing File", func(t *testing.T) {
		_, err := LoadConfig(filepath.Join(tmpdir, "missing.json"))
		if err == nil {
			t.Error("expected error for missing file, got nil")
		}
	})

	t.Run("Corrupted JSON", func(t *testing.T) {
		configPath := filepath.Join(tmpdir, "corrupted.json")
		os.WriteFile(configPath, []byte("{invalid json"), 0644)
		_, err := LoadConfig(configPath)
		if err == nil {
			t.Error("expected error for corrupted JSON, got nil")
		}
	})

	t.Run("Invalid Port", func(t *testing.T) {
		configPath := filepath.Join(tmpdir, "invalid_port.json")
		cfg := Config{Port: 70000, DownloadDir: tmpdir}
		data, _ := json.Marshal(cfg)
		os.WriteFile(configPath, data, 0644)
		_, err := LoadConfig(configPath)
		if err == nil {
			t.Error("expected error for invalid port, got nil")
		}
	})

	t.Run("Missing DownloadDir", func(t *testing.T) {
		configPath := filepath.Join(tmpdir, "missing_dir.json")
		cfg := Config{Port: 8080, DownloadDir: filepath.Join(tmpdir, "nonexistent")}
		data, _ := json.Marshal(cfg)
		os.WriteFile(configPath, data, 0644)
		_, err := LoadConfig(configPath)
		if err == nil {
			t.Error("expected error for nonexistent download dir, got nil")
		}
	})
}
