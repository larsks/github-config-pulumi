package readers

import (
	"fmt"
)

type (
	Repository struct {
		Name        string               `yaml:"name"`
		Description string               `yaml:"description"`
		HomepageURL string               `yaml:"homepageURL"`
		Visibility  RepositoryVisibility `yaml:"visibility"`

		AllowAutoMerge      *bool `yaml:"allowAutoMerge"`
		AutoInit            *bool `yaml:"autoInit"`
		HasDiscussions      *bool `yaml:"hasDiscussions"`
		HasDownloads        *bool `yaml:"hasDownloads"`
		HasIssues           *bool `yaml:"hasIssues"`
		HasProjects         *bool `yaml:"hasProjects"`
		HasWiki             *bool `yaml:"hasWiki"`
		IsTemplate          *bool `yaml:"istemplate"`
		VulnerabilityAlerts *bool `yaml:"vulnerabilityAlerts"`
		UseDefaultLabels    *bool `yaml:"useDefaultLabels"`
		UseDefaultTemplate  *bool `yaml:"useDefaultTemplate"`

		RequiredStatusChecks []string `yaml:"requiredStatusChecks"`

		Teams  []RepositoryTeamPermissions `yaml:"teams"`
		Labels []*Label                    `yaml:"labels"`
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

	defaults := map[**bool]bool{
		&r.HasDiscussions:      false,
		&r.HasDownloads:        true,
		&r.HasIssues:           true,
		&r.HasProjects:         false,
		&r.HasWiki:             false,
		&r.AutoInit:            true,
		&r.IsTemplate:          false,
		&r.VulnerabilityAlerts: false,
		&r.UseDefaultLabels:    true,
		&r.UseDefaultTemplate:  true,
	}

	for field, defaultVal := range defaults {
		if *field == nil {
			val := defaultVal
			*field = &val
		}
	}

	if r.AllowAutoMerge == nil || r.Visibility != RepositoryVisibilityPublic {
		f := false
		r.AllowAutoMerge = &f
	}
}

func ReadRepositories() ([]*Repository, error) {
	globPattern := fmt.Sprintf("%s/repositories/*.yaml", dataDirectory)
	return readYAMLFilesWithDefaults[*Repository](globPattern)
}
