package readers

import (
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

const (
	dataDirectory string = "data"
)

func readYAMLFile[T any](filePath string) (T, error) {
	var result T
	fd, err := os.Open(filePath)
	if err != nil {
		return result, err
	}
	defer fd.Close() //nolint:errcheck

	decoder := yaml.NewDecoder(fd)
	if err := decoder.Decode(&result); err != nil {
		return result, err
	}

	return result, nil
}

func readYAMLFiles[T any](globPattern string) ([]T, error) {
	var results []T

	files, err := filepath.Glob(globPattern)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		item, err := readYAMLFile[T](file)
		if err != nil {
			return nil, err
		}
		results = append(results, item)
	}

	return results, nil
}
