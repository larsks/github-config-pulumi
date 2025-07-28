package readers

import (
	"fmt"
)

type (
	MemberRole string

	Member struct {
		Name string     `yaml:"name" validate:"required"`
		Role MemberRole `yaml:"role" validate:"required,oneof=member admin"`
	}

	MembersFile struct {
		Members []Member `yaml:"members" validate:"dive"`
	}
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

func (mf *MembersFile) SetDefaults() {
	for i := range mf.Members {
		mf.Members[i].SetDefaults()
	}
}

func ReadMembers() ([]*Member, error) {
	filePath := fmt.Sprintf("%s/members.yaml", dataDirectory)
	membersFile, err := readYAMLFileWithDefaults[*MembersFile](filePath)
	if err != nil {
		return nil, err
	}

	var result []*Member
	for i := range membersFile.Members {
		result = append(result, &membersFile.Members[i])
	}
	return result, nil
}
