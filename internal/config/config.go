package config

import "os"

type Config struct {
	Port    string
	APIKey  string
	DataDir string
}

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
