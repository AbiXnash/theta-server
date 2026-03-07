package config

import "os"

// Config stores runtime configuration loaded from environment variables.
type Config struct {
	Port    string
	APIKey  string
	DataDir string
}

// Load reads configuration from environment variables and applies defaults.
func Load() Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	apiKey := os.Getenv("THETA_API_KEY")
	if apiKey == "" {
		apiKey = "somethings"
	}

	dataDir := os.Getenv("THETA_DATA_DIR")
	if dataDir == "" {
		dataDir = "data"
	}

	return Config{
		Port:    port,
		APIKey:  apiKey,
		DataDir: dataDir,
	}
}
