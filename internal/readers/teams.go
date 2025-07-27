package readers

import (
	"fmt"
)

type (
	TeamPrivacy string
	TeamRole    string

	Team struct {
		Name        string       `yaml:"name"`
		Description string       `yaml:"description"`
		Privacy     TeamPrivacy  `yaml:"privacy"`
		Members     []TeamMember `yaml:"members"`
	}

	TeamMember struct {
		Name string   `yaml:"name"`
		Role TeamRole `yaml:"role"`
	}
)

const (
	TeamPrivacyClosed TeamPrivacy = "closed"
)

const (
	TeamRoleMember     TeamRole = "member"
	TeamRoleMaintainer TeamRole = "maintainer"
)

func (t *Team) SetDefaults() {
	if t.Privacy == "" {
		t.Privacy = TeamPrivacyClosed
	}
	for i := range t.Members {
		if t.Members[i].Role == "" {
			t.Members[i].Role = TeamRoleMember
		}
	}
}

func ReadTeams() ([]Team, error) {
	globPattern := fmt.Sprintf("%s/teams/*.yaml", dataDirectory)
	return readYAMLFilesWithDefaults[Team, *Team](globPattern)
}
