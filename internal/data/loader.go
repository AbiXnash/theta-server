package data

import (
	"fmt"
	"os"
	"path/filepath"
)

type Loader struct {
	baseDir string
}

func NewLoader(baseDir string) Loader {
	return Loader{baseDir: baseDir}
}

func (l Loader) ReadJSON(fileName string) ([]byte, error) {
	path := filepath.Join(l.baseDir, fileName)
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", path, err)
	}
	return data, nil
}
