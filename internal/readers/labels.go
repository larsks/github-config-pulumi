package readers

import (
	"fmt"
)

type (
	Label struct {
		Name        string `yaml:"name" validate:"required"`
		Description string `yaml:"description"`
		Color       string `yaml:"color" validate:"required"`
	}

	LabelsFile struct {
		Labels []Label `yaml:"labels" validate:"dive"`
	}
)

func (l *Label) SetDefaults() {
	// No defaults needed for labels currently
}

func (lf *LabelsFile) SetDefaults() {
	for i := range lf.Labels {
		lf.Labels[i].SetDefaults()
	}
}

func ReadLabels() ([]*Label, error) {
	filePath := fmt.Sprintf("%s/labels.yaml", dataDirectory)
	labelsFile, err := readYAMLFileWithDefaults[*LabelsFile](filePath)
	if err != nil {
		return nil, err
	}

	var result []*Label
	for i := range labelsFile.Labels {
		result = append(result, &labelsFile.Labels[i])
	}
	return result, nil
}
