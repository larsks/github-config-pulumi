package readers

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type (
	Organization struct {
		Name string `yaml:"name"`
	}
)

func ReadOrganization() (Organization, error) {
	var org Organization
	orgFile := fmt.Sprintf("%s/organization.yaml", dataDirectory)

	fd, err := os.Open(orgFile)
	if err != nil {
		return Organization{}, err
	}

	decoder := yaml.NewDecoder(fd)
	if err := decoder.Decode(&org); err != nil {
		return Organization{}, err
	}

	return org, nil
}
