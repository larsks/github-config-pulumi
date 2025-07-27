package readers

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
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

func readRepository(repoFile string) (Repository, error) {
	var repo = Repository{}
	fd, err := os.Open(repoFile)
	if err != nil {
		return Repository{}, err
	}

	decoder := yaml.NewDecoder(fd)
	if err := decoder.Decode(&repo); err != nil {
		return Repository{}, err
	}

	return repo, nil

}

func ReadRepositories() ([]Repository, error) {
	var repos []Repository

	files, err := filepath.Glob(fmt.Sprintf("%s/repositories/*.yaml", dataDirectory))
	if err != nil {
		return nil, err
	}

	for _, repoFile := range files {
		repo, err := readRepository(repoFile)
		if err != nil {
			return nil, err
		}
		repos = append(repos, repo)
	}

	return repos, nil
}
