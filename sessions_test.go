package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseXSession(t *testing.T) {
	tests := []struct {
		name    string
		content string
		want    xsession
		wantErr bool
	}{
		{
			name: "valid desktop entry",
			content: `[Desktop Entry]
Name=i3
Exec=i3
Type=Application
`,
			want: xsession{Name: "i3", Exec: "i3"},
		},
		{
			name: "ignores other groups",
			content: `[Other Section]
Name=Wrong
Exec=wrong

[Desktop Entry]
Name=sway
Exec=sway
`,
			want: xsession{Name: "sway", Exec: "sway"},
		},
		{
			name:    "missing exec",
			content: "[Desktop Entry]\nName=broken\n",
			wantErr: true,
		},
		{
			name:    "missing name",
			content: "[Desktop Entry]\nExec=broken\n",
			wantErr: true,
		},
		{
			name:    "empty exec value",
			content: "[Desktop Entry]\nName=test\nExec=\n",
			wantErr: true,
		},
		{
			name:    "empty name value",
			content: "[Desktop Entry]\nName=\nExec=test\n",
			wantErr: true,
		},
		{
			name:    "no desktop entry group",
			content: "[Other]\nName=test\nExec=test\n",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpFile := filepath.Join(t.TempDir(), "test.desktop")
			if err := os.WriteFile(tmpFile, []byte(tt.content), 0644); err != nil {
				t.Fatal(err)
			}
			got, err := parseXSession(tmpFile)
			if tt.wantErr {
				if err == nil {
					t.Error("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got.Name != tt.want.Name || got.Exec != tt.want.Exec {
				t.Errorf("got %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestParseXSessionFileNotFound(t *testing.T) {
	_, err := parseXSession("/nonexistent/path/test.desktop")
	if err == nil {
		t.Error("expected error for nonexistent file, got nil")
	}
}

func TestGetSession(t *testing.T) {
	// .xinitrc is always in the session list
	exec := getSession(".xinitrc")
	if exec != "/bin/bash --login .xinitrc" {
		t.Errorf("expected xinitrc exec, got %q", exec)
	}
}

func TestGetSessionNotFound(t *testing.T) {
	exec := getSession("nonexistent-session")
	if exec != "" {
		t.Errorf("expected empty string for nonexistent session, got %q", exec)
	}
}

func TestGetXDGDirs(t *testing.T) {
	t.Run("default dirs", func(t *testing.T) {
		// Unset env vars to test defaults
		origHome := os.Getenv("XDG_DATA_HOME")
		origDirs := os.Getenv("XDG_DATA_DIRS")
		os.Unsetenv("XDG_DATA_HOME")
		os.Unsetenv("XDG_DATA_DIRS")
		defer func() {
			if origHome != "" {
				os.Setenv("XDG_DATA_HOME", origHome)
			}
			if origDirs != "" {
				os.Setenv("XDG_DATA_DIRS", origDirs)
			}
		}()

		dirs := getXDGDirs()
		if len(dirs) == 0 {
			t.Fatal("expected at least one XDG dir")
		}
		// Should contain .local/share from HOME fallback
		found := false
		for _, dir := range dirs {
			if filepath.Base(dir) == "share" {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("expected a 'share' directory in XDG dirs: %v", dirs)
		}
	})

	t.Run("custom XDG_DATA_HOME", func(t *testing.T) {
		origHome := os.Getenv("XDG_DATA_HOME")
		origDirs := os.Getenv("XDG_DATA_DIRS")
		os.Setenv("XDG_DATA_HOME", "/custom/data")
		os.Unsetenv("XDG_DATA_DIRS")
		defer func() {
			if origHome != "" {
				os.Setenv("XDG_DATA_HOME", origHome)
			} else {
				os.Unsetenv("XDG_DATA_HOME")
			}
			if origDirs != "" {
				os.Setenv("XDG_DATA_DIRS", origDirs)
			}
		}()

		dirs := getXDGDirs()
		if dirs[0] != "/custom/data" {
			t.Errorf("expected first dir to be /custom/data, got %q", dirs[0])
		}
	})
}

func TestLoadSessions(t *testing.T) {
	sessions := loadSessions()
	// Should always have at least .xinitrc
	if len(sessions) == 0 {
		t.Fatal("expected at least one session (.xinitrc)")
	}
	if sessions[0].Name != ".xinitrc" {
		t.Errorf("expected first session to be .xinitrc, got %q", sessions[0].Name)
	}
}

func TestLoadSessionsWithTempDir(t *testing.T) {
	tmpDir := t.TempDir()
	xsessDir := filepath.Join(tmpDir, "xsessions")
	if err := os.MkdirAll(xsessDir, 0755); err != nil {
		t.Fatal(err)
	}

	desktopContent := `[Desktop Entry]
Name=TestWM
Exec=testwm --start
Type=Application
`
	if err := os.WriteFile(filepath.Join(xsessDir, "testwm.desktop"), []byte(desktopContent), 0644); err != nil {
		t.Fatal(err)
	}

	origHome := os.Getenv("XDG_DATA_HOME")
	origDirs := os.Getenv("XDG_DATA_DIRS")
	os.Setenv("XDG_DATA_HOME", tmpDir)
	os.Setenv("XDG_DATA_DIRS", tmpDir)
	defer func() {
		if origHome != "" {
			os.Setenv("XDG_DATA_HOME", origHome)
		} else {
			os.Unsetenv("XDG_DATA_HOME")
		}
		if origDirs != "" {
			os.Setenv("XDG_DATA_DIRS", origDirs)
		} else {
			os.Unsetenv("XDG_DATA_DIRS")
		}
	}()

	sessions := loadSessions()
	found := false
	for _, sess := range sessions {
		if sess.Name == "TestWM" && sess.Exec == "testwm --start" {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("expected to find TestWM session, got: %+v", sessions)
	}
}
