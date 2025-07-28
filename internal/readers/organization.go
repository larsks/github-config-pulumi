package readers

import (
	"fmt"
)

type (
	Organization struct {
		Name                             string                      `yaml:"name" validate:"required"`
		DefaultTemplate                  *TemplateSpec               `yaml:"defaultTemplate" validate:"omitempty"`
		DefaultRepositoryTeamPermissions []RepositoryTeamPermissions `yaml:"defaultRepositoryTeamPermissions"`
	}

	TemplateSpec struct {
		Owner              string `yaml:"owner"`
		Repository         string `yaml:"repository" validate:"required"`
		IncludeAllBranches *bool  `yaml:"includeAllBranches"`
	}
)

func (o *Organization) SetDefaults() {
	defaults := map[**bool]bool{}

	for field, defaultVal := range defaults {
		if *field == nil {
			val := defaultVal
			*field = &val
		}
	}

	if o.DefaultTemplate != nil {
		if o.DefaultTemplate.Owner == "" {
			o.DefaultTemplate.Owner = o.Name
		}

		if o.DefaultTemplate.IncludeAllBranches == nil {
			f := false
			o.DefaultTemplate.IncludeAllBranches = &f
		}
	}
}

func ReadOrganization() (*Organization, error) {
	filePath := fmt.Sprintf("%s/organization.yaml", dataDirectory)
	return readYAMLFileWithDefaults[*Organization](filePath)
}
