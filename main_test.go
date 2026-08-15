package main

import "testing"

func TestGetEnvFile(t *testing.T) {
	tests := []struct {
		name     string
		env      string
		expected string
	}{
		{"default empty uses dev", "", ".env.dev"},
		{"development uses dev", "development", ".env.dev"},
		{"dev alias uses dev", "dev", ".env.dev"},
		{"production uses prod", "production", ".env.prod"},
		{"unknown uses dev", "staging", ".env.dev"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getEnvFile(tt.env)
			if got != tt.expected {
				t.Errorf("getEnvFile(%q) = %q, want %q", tt.env, got, tt.expected)
			}
		})
	}
}
