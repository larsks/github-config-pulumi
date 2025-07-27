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

	Labels []Label
)

func (l *Label) SetDefaults() {
	// No defaults needed for labels currently
}

func (labels *Labels) SetDefaults() {
	for i := range *labels {
		(*labels)[i].SetDefaults()
	}
}

func ReadLabels() ([]Label, error) {
	filePath := fmt.Sprintf("%s/labels.yaml", dataDirectory)
	labels, err := readYAMLFileWithDefaults[Labels, *Labels](filePath)
	return []Label(labels), err
}
