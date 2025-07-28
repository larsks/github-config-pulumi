package readers

import (
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

const (
	dataDirectory string = "data"
)

type Defaultable interface {
	SetDefaults()
}

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

func readYAMLFilesWithDefaults[PT Defaultable](globPattern string) ([]PT, error) {
	files, err := filepath.Glob(globPattern)
	if err != nil {
		return nil, err
	}

	var results []PT
	for _, file := range files {
		var item PT
		fd, err := os.Open(file)
		if err != nil {
			return nil, err
		}

		decoder := yaml.NewDecoder(fd)
		if err := decoder.Decode(&item); err != nil {
			fd.Close()
			return nil, err
		}
		fd.Close()

		item.SetDefaults()
		results = append(results, item)
	}

	return results, nil
}

func readYAMLFileWithDefaults[PT Defaultable](filePath string) (PT, error) {
	var item PT
	fd, err := os.Open(filePath)
	if err != nil {
		return item, err
	}
	defer fd.Close()

	decoder := yaml.NewDecoder(fd)
	if err := decoder.Decode(&item); err != nil {
		return item, err
	}

	item.SetDefaults()
	return item, nil
}
