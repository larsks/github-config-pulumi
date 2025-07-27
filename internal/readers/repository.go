package readers

import (
	"fmt"
)

type (
	Repository struct {
		Name                 string                      `yaml:"name"`
		Description          string                      `yaml:"description"`
		Visibility           RepositoryVisibility        `yaml:"visibility"`
		HomepageURL          string                      `yaml:"homepageURL"`
		RequiredStatusChecks []string                    `yaml:"requiredStatusChecks"`
		Teams                []RepositoryTeamPermissions `yaml:"teams"`
	}

	RepositoryTeamPermissions struct {
		Name       string               `yaml:"name"`
		Permission RepositoryPermission `yaml:"permission"`
	}

	RepositoryPermission string
	RepositoryVisibility string
)

const (
	RepositoryPermissionPush  RepositoryPermission = "push"
	RepositoryPermissionAdmin RepositoryPermission = "admin"
)

const (
	RepositoryVisibilityPublic   RepositoryVisibility = "public"
	RepositoryVisibilityPrivate  RepositoryVisibility = "private"
	RepositoryVisibilityInternal RepositoryVisibility = "internal"
)

func (r *Repository) SetDefaults() {
	if r.Visibility == "" {
		r.Visibility = RepositoryVisibilityPublic
	}
}

func ReadRepositories() ([]Repository, error) {
	globPattern := fmt.Sprintf("%s/repositories/*.yaml", dataDirectory)
	return readYAMLFilesWithDefaults[Repository, *Repository](globPattern)
}
