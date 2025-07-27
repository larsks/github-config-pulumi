package readers

import (
	"fmt"
)

type (
	Repository struct {
		Name                 string                      `yaml:"name"`
		Description          string                      `yaml:"description"`
		HomepageURL          string                      `yaml:"homepageURL"`
		RequiredStatusChecks []string                    `yaml:"requiredStatusChecks"`
		Teams                []RepositoryTeamPermissions `yaml:"teams"`
	}

	RepositoryTeamPermissions struct {
		Name       string               `yaml:"name"`
		Permission RepositoryPermission `yaml:"permission"`
	}

	RepositoryPermission string
)

const (
	RepositoryPermissionPush  RepositoryPermission = "push"
	RepositoryPermissionAdmin RepositoryPermission = "admin"
)

func ReadRepositories() ([]Repository, error) {
	globPattern := fmt.Sprintf("%s/repositories/*.yaml", dataDirectory)
	return readYAMLFiles[Repository](globPattern)
}
