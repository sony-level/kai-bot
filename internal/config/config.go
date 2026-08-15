// Package config loads and validates Kai's environment configuration.
package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// Config holds the values required to run Kai.
type Config struct {
	Token string
}

// Load reads environment variables and, in development, a .env file.
// It returns an error if a required value is missing.
func Load() (*Config, error) {
	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}

	envFile := ".env.dev"
	if env == "production" {
		envFile = ".env.prod"
	}

	// Missing files are ignored so that production deployments
	// can provide values via the host environment.
	_ = godotenv.Load(filepath.Join(".", envFile))

	c := &Config{
		Token: os.Getenv("DISCORD_TOKEN"),
	}

	if c.Token == "" {
		return nil, fmt.Errorf("configuration invalide : DISCORD_TOKEN est absent")
	}

	return c, nil
}
