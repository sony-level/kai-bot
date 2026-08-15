// Package config tests the configuration loader and validation.
package config

import "testing"

func TestLoadValid(t *testing.T) {
	t.Setenv("DISCORD_TOKEN", "test-token")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if cfg.Token != "test-token" {
		t.Errorf("Token = %q, want %q", cfg.Token, "test-token")
	}
}

func TestLoadMissingToken(t *testing.T) {
	t.Setenv("DISCORD_TOKEN", "")

	if _, err := Load(); err == nil {
		t.Fatal("expected an error for missing DISCORD_TOKEN")
	}
}
