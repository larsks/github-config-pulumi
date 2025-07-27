package readers

import (
	"fmt"
)

type (
	Organization struct {
		Name string `yaml:"name"`
	}
)

func ReadOrganization() (Organization, error) {
	filePath := fmt.Sprintf("%s/organization.yaml", dataDirectory)
	return readYAMLFile[Organization](filePath)
}
