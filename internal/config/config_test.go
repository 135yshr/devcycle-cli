package config

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestConfigDirPath(t *testing.T) {
	dir, err := ConfigDirPath()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !strings.HasSuffix(dir, ConfigDir) {
		t.Errorf("expected dir to end with %s, got %s", ConfigDir, dir)
	}
}

func TestConfigFilePath(t *testing.T) {
	path, err := ConfigFilePath()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !strings.HasSuffix(path, filepath.Join(ConfigDir, ConfigFile)) {
		t.Errorf("expected path to end with %s, got %s", filepath.Join(ConfigDir, ConfigFile), path)
	}
}

func TestTokenFilePath(t *testing.T) {
	path, err := TokenFilePath()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !strings.HasSuffix(path, filepath.Join(ConfigDir, TokenFile)) {
		t.Errorf("expected path to end with %s, got %s", filepath.Join(ConfigDir, TokenFile), path)
	}
}

func TestEnsureConfigDir(t *testing.T) {
	// Save current directory
	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get current dir: %v", err)
	}

	// Create temp directory and change to it
	tempDir := t.TempDir()
	if err := os.Chdir(tempDir); err != nil {
		t.Fatalf("failed to change dir: %v", err)
	}
	defer os.Chdir(originalDir)

	err = EnsureConfigDir()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	configDir := filepath.Join(tempDir, ConfigDir)
	info, err := os.Stat(configDir)
	if os.IsNotExist(err) {
		t.Fatal("config directory was not created")
	}
	if !info.IsDir() {
		t.Fatal("config path is not a directory")
	}
}

func TestSetAndGetProject(t *testing.T) {
	projectKey := "test-project"

	SetProject(projectKey)

	result := Project()
	if result != projectKey {
		t.Errorf("expected %s, got %s", projectKey, result)
	}
}

func TestSetAndGetEnvironment(t *testing.T) {
	envKey := "test-environment"

	SetEnvironment(envKey)

	result := Environment()
	if result != envKey {
		t.Errorf("expected %s, got %s", envKey, result)
	}
}
