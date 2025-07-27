package readers

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type (
	Label struct {
		Name        string `yaml:"name"`
		Description string `yaml:"description"`
		Color       string `yaml:"color"`
	}
)

func ReadLabels() ([]Label, error) {
	var labels []Label
	labelFile := fmt.Sprintf("%s/labels.yaml", dataDirectory)

	fd, err := os.Open(labelFile)
	if err != nil {
		return nil, err
	}

	decoder := yaml.NewDecoder(fd)
	if err := decoder.Decode(&labels); err != nil {
		return nil, err
	}

	return labels, nil
}
