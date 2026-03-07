package data

import (
	"fmt"
	"os"
	"path/filepath"
)

// Loader reads JSON files from a configured base directory.
type Loader struct {
	baseDir string
}

// NewLoader creates a JSON loader rooted at baseDir.
func NewLoader(baseDir string) Loader {
	return Loader{baseDir: baseDir}
}

// ReadJSON loads a single JSON file by name from the loader base directory.
func (l Loader) ReadJSON(fileName string) ([]byte, error) {
	path := filepath.Join(l.baseDir, fileName)
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", path, err)
	}
	return data, nil
}
