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
		HasDownloads        *bool `yaml:"hasDownloads"`
		HasIssues           *bool `yaml:"hasIssues"`
		HasProjects         *bool `yaml:"hasProjects"`
		HasWiki             *bool `yaml:"hasWiki"`
		IsTemplate          *bool `yaml:"istemplate"`
		VulnerabilityAlerts *bool `yaml:"vulnerabilityAlerts"`
		UseCommonLabels     *bool `yaml:"useCommonLabels"`

		RequiredStatusChecks []string `yaml:"requiredStatusChecks"`

		Teams  []RepositoryTeamPermissions `yaml:"teams"`
		Labels []Label                     `yaml:"labels"`
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
	t := true
	f := false

	if r.Visibility == "" {
		r.Visibility = RepositoryVisibilityPublic
	}

	if r.HasDownloads == nil {
		r.HasDownloads = &t
	}

	if r.HasIssues == nil {
		r.HasIssues = &t
	}

	if r.HasProjects == nil {
		r.HasProjects = &f
	}

	if r.HasWiki == nil {
		r.HasWiki = &f
	}

	if r.HasIssues == nil {
		r.HasIssues = &t
	}

	if r.AllowAutoMerge == nil || r.Visibility == RepositoryVisibilityPrivate {
		r.AllowAutoMerge = &f
	}

	if r.AutoInit == nil {
		r.AutoInit = &t
	}

	if r.IsTemplate == nil {
		r.IsTemplate = &f
	}

	if r.VulnerabilityAlerts == nil {
		r.VulnerabilityAlerts = &f
	}

	if r.UseCommonLabels == nil {
		r.UseCommonLabels = &t
	}
}

func ReadRepositories() ([]*Repository, error) {
	globPattern := fmt.Sprintf("%s/repositories/*.yaml", dataDirectory)
	return readYAMLFilesWithDefaults[*Repository](globPattern)
}
