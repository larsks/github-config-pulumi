package readers

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type (
	MemberRole string

	Member struct {
		Name string     `yaml:"name"`
		Role MemberRole `yaml:"role"`
	}
)

const (
	MemberRoleMember MemberRole = "member"
	MemberRoleAdmin  MemberRole = "admin"
)

func ReadMembers() ([]Member, error) {
	var members []Member
	memberFile := fmt.Sprintf("%s/members.yaml", dataDirectory)

	fd, err := os.Open(memberFile)
	if err != nil {
		return nil, err
	}

	decoder := yaml.NewDecoder(fd)
	if err := decoder.Decode(&members); err != nil {
		return nil, err
	}

	return members, nil
}
