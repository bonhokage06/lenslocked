package helpers

import (
	"os"
	"path/filepath"
)

// create a function that get all files in a directory

func GetFiles(dir string) ([]string, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var paths []string
	for _, file := range files {
		paths = append(paths, filepath.Join(dir, file.Name()))
	}
	return paths, nil
}
