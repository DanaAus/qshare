package handlers

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSanitizePath(t *testing.T) {
	wd, _ := os.Getwd()
	base := filepath.Join(wd, "testdata")
	
	cases := []struct {
		name    string
		target  string
		wantErr bool
	}{
		{"Valid same dir", filepath.Join(base, "file.txt"), false},
		{"Valid subdir", filepath.Join(base, "subdir", "file.txt"), false},
		{"Relative valid", "testdata/file.txt", false},
		{"Traversal up", filepath.Join(base, "..", "main.go"), true},
		{"Absolute outside", "/etc/passwd", true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := SanitizePath(base, tc.target)
			if (err != nil) != tc.wantErr {
				t.Errorf("SanitizePath() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if !tc.wantErr {
				absBase, _ := filepath.Abs(base)
				absGot, _ := filepath.Abs(got)
				if absGot[:len(absBase)] != absBase {
					t.Errorf("SanitizePath() got %v, outside base %v", absGot, absBase)
				}
			}
		})
	}
}
