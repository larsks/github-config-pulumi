package readers

import (
	"fmt"
)

type (
	Organization struct {
		Name string `yaml:"name"`
	}
)

func (o *Organization) SetDefaults() {
	// No defaults needed for organization currently
}

func ReadOrganization() (*Organization, error) {
	filePath := fmt.Sprintf("%s/organization.yaml", dataDirectory)
	return readYAMLFileWithDefaults[*Organization](filePath)
}
