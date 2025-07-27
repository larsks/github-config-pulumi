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
)

const (
	MemberRoleMember MemberRole = "member"
	MemberRoleAdmin  MemberRole = "admin"
)

func ReadMembers() ([]Member, error) {
	filePath := fmt.Sprintf("%s/members.yaml", dataDirectory)
	return readYAMLFile[[]Member](filePath)
}
