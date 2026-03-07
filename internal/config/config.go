package config

import (
	"fmt"
	"os"
)

// Config stores runtime configuration loaded from environment variables.
type Config struct {
	Port    string
	APIKey  string
	DataDir string
}

// Load reads configuration from environment variables and applies defaults.
func Load() (Config, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	apiKey := os.Getenv("THETA_API_KEY")
	if apiKey == "" {
		return Config{}, fmt.Errorf("THETA_API_KEY is required")
	}

	dataDir := os.Getenv("THETA_DATA_DIR")
	if dataDir == "" {
		dataDir = "data"
	}

	return Config{
		Port:    port,
		APIKey:  apiKey,
		DataDir: dataDir,
	}, nil
}
