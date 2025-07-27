package readers

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
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

func readTeam(teamFile string) (Team, error) {
	var team = Team{}
	fd, err := os.Open(teamFile)
	if err != nil {
		return Team{}, err
	}

	decoder := yaml.NewDecoder(fd)
	if err := decoder.Decode(&team); err != nil {
		return Team{}, err
	}

	return team, nil
}

func ReadTeams() ([]Team, error) {
	var teams []Team

	files, err := filepath.Glob(fmt.Sprintf("%s/teams/*.yaml", dataDirectory))
	if err != nil {
		return nil, err
	}

	for _, teamFile := range files {
		team, err := readTeam(teamFile)
		if err != nil {
			return nil, err
		}
		teams = append(teams, team)
	}

	return teams, nil
}
