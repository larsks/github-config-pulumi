package readers

import (
	"fmt"
)

type (
	Organization struct {
		Name                             string                      `yaml:"name"`
		DefaultTemplate                  TemplateSpec                `yaml:"defaultTemplate"`
		DefaultRepositoryTeamPermissions []RepositoryTeamPermissions `yaml:"defaultRepositoryTeamPermissions"`
	}

	TemplateSpec struct {
		Owner              string `yaml:"owner"`
		Repository         string `yaml:"repository"`
		IncludeAllBranches *bool  `yaml:"includeAllBranches"`
	}
)

func (o *Organization) SetDefaults() {
	defaults := map[**bool]bool{
		&o.DefaultTemplate.IncludeAllBranches: false,
	}

	for field, defaultVal := range defaults {
		if *field == nil {
			val := defaultVal
			*field = &val
		}
	}

	if o.DefaultTemplate.Owner == "" {
		o.DefaultTemplate.Owner = o.Name
	}
}

func ReadOrganization() (*Organization, error) {
	filePath := fmt.Sprintf("%s/organization.yaml", dataDirectory)
	return readYAMLFileWithDefaults[*Organization](filePath)
}
