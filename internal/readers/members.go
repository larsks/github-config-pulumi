package readers

import (
	"fmt"
)

type (
	MemberRole string

	Member struct {
		Name string     `yaml:"name"`
		Role MemberRole `yaml:"role"`
	}

	Members []Member
)

const (
	MemberRoleMember MemberRole = "member"
	MemberRoleAdmin  MemberRole = "admin"
)

func (m *Member) SetDefaults() {
	if m.Role == "" {
		m.Role = MemberRoleMember
	}
}

func (members *Members) SetDefaults() {
	for i := range *members {
		(*members)[i].SetDefaults()
	}
}

func ReadMembers() ([]*Member, error) {
	filePath := fmt.Sprintf("%s/members.yaml", dataDirectory)
	members, err := readYAMLFileWithDefaults[*Members](filePath)
	if err != nil {
		return nil, err
	}

	var result []*Member
	for i := range *members {
		result = append(result, &(*members)[i])
	}
	return result, nil
}
