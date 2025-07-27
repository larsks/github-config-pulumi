package readers

import (
	"fmt"
)

type (
	Label struct {
		Name        string `yaml:"name"`
		Description string `yaml:"description"`
		Color       string `yaml:"color"`
	}
)

func ReadLabels() ([]Label, error) {
	filePath := fmt.Sprintf("%s/labels.yaml", dataDirectory)
	return readYAMLFile[[]Label](filePath)
}
