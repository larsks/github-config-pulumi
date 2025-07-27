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

func ReadTeams() ([]Team, error) {
	globPattern := fmt.Sprintf("%s/teams/*.yaml", dataDirectory)
	return readYAMLFiles[Team](globPattern)
}
