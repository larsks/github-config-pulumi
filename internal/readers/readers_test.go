package readers

import (
	"os"
	"path/filepath"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestValidateStruct_ValidData(t *testing.T) {
	tests := []struct {
		name string
		data any
	}{
		{
			name: "nil pointer should not error",
			data: nil,
		},
		{
			name: "slice should be skipped",
			data: []string{"test", "data"},
		},
		{
			name: "pointer to slice should be skipped",
			data: &[]string{"test", "data"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateStruct(tt.data)
			if err != nil {
				t.Errorf("validateStruct() error = %v, expected nil", err)
			}
		})
	}
}

func TestOrganizationValidation(t *testing.T) {
	tests := []struct {
		name        string
		org         Organization
		expectError bool
	}{
		{
			name: "valid organization with template",
			org: Organization{
				Name: "test-org",
				DefaultTemplate: &TemplateSpec{
					Owner:      "test-owner",
					Repository: "test-repo",
				},
			},
			expectError: false,
		},
		{
			name: "valid organization without template",
			org: Organization{
				Name: "test-org",
			},
			expectError: false,
		},
		{
			name: "missing organization name",
			org: Organization{
				DefaultTemplate: &TemplateSpec{
					Owner:      "test-owner",
					Repository: "test-repo",
				},
			},
			expectError: true,
		},
		{
			name: "missing repository in template",
			org: Organization{
				Name: "test-org",
				DefaultTemplate: &TemplateSpec{
					Owner: "test-owner",
				},
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Apply defaults like the real code does
			tt.org.SetDefaults()

			err := validateStruct(&tt.org)
			if tt.expectError && err == nil {
				t.Error("validateStruct() expected error but got nil")
			}
			if !tt.expectError && err != nil {
				t.Errorf("validateStruct() unexpected error = %v", err)
			}
		})
	}
}

func TestRepositoryValidation(t *testing.T) {
	tests := []struct {
		name        string
		repo        Repository
		expectError bool
	}{
		{
			name: "valid repository",
			repo: Repository{
				Name:       "test-repo",
				Visibility: RepositoryVisibilityPublic,
			},
			expectError: false,
		},
		{
			name: "missing repository name",
			repo: Repository{
				Visibility: RepositoryVisibilityPublic,
			},
			expectError: true,
		},
		{
			name: "invalid visibility",
			repo: Repository{
				Name:       "test-repo",
				Visibility: "invalid",
			},
			expectError: true,
		},
		{
			name: "valid team permissions",
			repo: Repository{
				Name:       "test-repo",
				Visibility: RepositoryVisibilityPublic,
				Teams: []RepositoryTeamPermissions{
					{
						Name:       "team1",
						Permission: RepositoryPermissionPush,
					},
				},
			},
			expectError: false,
		},
		{
			name: "invalid team permission",
			repo: Repository{
				Name:       "test-repo",
				Visibility: RepositoryVisibilityPublic,
				Teams: []RepositoryTeamPermissions{
					{
						Name:       "team1",
						Permission: "invalid",
					},
				},
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Apply defaults like the real code does
			tt.repo.SetDefaults()

			err := validateStruct(&tt.repo)
			if tt.expectError && err == nil {
				t.Error("validateStruct() expected error but got nil")
			}
			if !tt.expectError && err != nil {
				t.Errorf("validateStruct() unexpected error = %v", err)
			}
		})
	}
}

func TestMemberValidation(t *testing.T) {
	tests := []struct {
		name        string
		member      Member
		expectError bool
	}{
		{
			name: "valid member",
			member: Member{
				Name: "testuser",
				Role: MemberRoleMember,
			},
			expectError: false,
		},
		{
			name: "valid admin",
			member: Member{
				Name: "adminuser",
				Role: MemberRoleAdmin,
			},
			expectError: false,
		},
		{
			name: "missing name",
			member: Member{
				Role: MemberRoleMember,
			},
			expectError: true,
		},
		{
			name: "invalid role",
			member: Member{
				Name: "testuser",
				Role: "invalid",
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Apply defaults like the real code does
			tt.member.SetDefaults()

			err := validateStruct(&tt.member)
			if tt.expectError && err == nil {
				t.Error("validateStruct() expected error but got nil")
			}
			if !tt.expectError && err != nil {
				t.Errorf("validateStruct() unexpected error = %v", err)
			}
		})
	}
}

func TestTeamValidation(t *testing.T) {
	tests := []struct {
		name        string
		team        Team
		expectError bool
	}{
		{
			name: "valid team",
			team: Team{
				Name:    "test-team",
				Privacy: TeamPrivacyClosed,
				Members: []TeamMember{
					{
						Name: "user1",
						Role: TeamRoleMember,
					},
				},
			},
			expectError: false,
		},
		{
			name: "missing team name",
			team: Team{
				Privacy: TeamPrivacyClosed,
			},
			expectError: true,
		},
		{
			name: "invalid team member role",
			team: Team{
				Name:    "test-team",
				Privacy: TeamPrivacyClosed,
				Members: []TeamMember{
					{
						Name: "user1",
						Role: "invalid",
					},
				},
			},
			expectError: true,
		},
		{
			name: "missing team member name",
			team: Team{
				Name:    "test-team",
				Privacy: TeamPrivacyClosed,
				Members: []TeamMember{
					{
						Role: TeamRoleMember,
					},
				},
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Apply defaults like the real code does
			tt.team.SetDefaults()

			err := validateStruct(&tt.team)
			if tt.expectError && err == nil {
				t.Error("validateStruct() expected error but got nil")
			}
			if !tt.expectError && err != nil {
				t.Errorf("validateStruct() unexpected error = %v", err)
			}
		})
	}
}

func TestLabelValidation(t *testing.T) {
	tests := []struct {
		name        string
		label       Label
		expectError bool
	}{
		{
			name: "valid label",
			label: Label{
				Name:        "bug",
				Description: "Something isn't working",
				Color:       "d73a4a",
			},
			expectError: false,
		},
		{
			name: "missing label name",
			label: Label{
				Description: "Something isn't working",
				Color:       "d73a4a",
			},
			expectError: true,
		},
		{
			name: "missing label color",
			label: Label{
				Name:        "bug",
				Description: "Something isn't working",
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Apply defaults like the real code does
			tt.label.SetDefaults()

			err := validateStruct(&tt.label)
			if tt.expectError && err == nil {
				t.Error("validateStruct() expected error but got nil")
			}
			if !tt.expectError && err != nil {
				t.Errorf("validateStruct() unexpected error = %v", err)
			}
		})
	}
}

func TestMembersFileValidation(t *testing.T) {
	tests := []struct {
		name        string
		membersFile MembersFile
		expectError bool
	}{
		{
			name: "valid members file",
			membersFile: MembersFile{
				Members: []Member{
					{Name: "user1", Role: MemberRoleMember},
					{Name: "user2", Role: MemberRoleAdmin},
				},
			},
			expectError: false,
		},
		{
			name: "invalid members file - missing name",
			membersFile: MembersFile{
				Members: []Member{
					{Name: "user1", Role: MemberRoleMember},
					{Role: MemberRoleAdmin}, // Missing name
				},
			},
			expectError: true,
		},
		{
			name: "invalid members file - invalid role",
			membersFile: MembersFile{
				Members: []Member{
					{Name: "user1", Role: MemberRoleMember},
					{Name: "user2", Role: "invalid"}, // Invalid role
				},
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Apply defaults like the real code does
			tt.membersFile.SetDefaults()

			err := validateStruct(&tt.membersFile)
			if tt.expectError && err == nil {
				t.Error("validateStruct() expected error but got nil")
			}
			if !tt.expectError && err != nil {
				t.Errorf("validateStruct() unexpected error = %v", err)
			}
		})
	}
}

func TestLabelsFileValidation(t *testing.T) {
	tests := []struct {
		name        string
		labelsFile  LabelsFile
		expectError bool
	}{
		{
			name: "valid labels file",
			labelsFile: LabelsFile{
				Labels: []Label{
					{Name: "bug", Color: "d73a4a", Description: "Something isn't working"},
					{Name: "feature", Color: "a2eeef", Description: "New feature request"},
				},
			},
			expectError: false,
		},
		{
			name: "invalid labels file - missing name",
			labelsFile: LabelsFile{
				Labels: []Label{
					{Name: "bug", Color: "d73a4a", Description: "Something isn't working"},
					{Color: "a2eeef", Description: "Missing name"}, // Missing name
				},
			},
			expectError: true,
		},
		{
			name: "invalid labels file - missing color",
			labelsFile: LabelsFile{
				Labels: []Label{
					{Name: "bug", Color: "d73a4a", Description: "Something isn't working"},
					{Name: "feature", Description: "Missing color"}, // Missing color
				},
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Apply defaults like the real code does
			tt.labelsFile.SetDefaults()

			err := validateStruct(&tt.labelsFile)
			if tt.expectError && err == nil {
				t.Error("validateStruct() expected error but got nil")
			}
			if !tt.expectError && err != nil {
				t.Errorf("validateStruct() unexpected error = %v", err)
			}
		})
	}
}

func TestYAMLIntegration(t *testing.T) {
	// Create temporary directory for test files
	tempDir, err := os.MkdirTemp("", "readers_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	tests := []struct {
		name        string
		yamlContent string
		structType  string
		expectError bool
	}{
		{
			name: "valid organization YAML with template",
			yamlContent: `name: "test-org"
defaultTemplate:
  owner: "test-owner"
  repository: "test-repo"`,
			structType:  "organization",
			expectError: false,
		},
		{
			name:        "valid organization YAML without template",
			yamlContent: `name: "test-org"`,
			structType:  "organization",
			expectError: false,
		},
		{
			name: "invalid organization YAML - missing name",
			yamlContent: `defaultTemplate:
  owner: "test-owner"
  repository: "test-repo"`,
			structType:  "organization",
			expectError: true,
		},
		{
			name: "valid repository YAML",
			yamlContent: `name: "test-repo"
description: "A test repository"
visibility: "public"`,
			structType:  "repository",
			expectError: false,
		},
		{
			name: "invalid repository YAML - bad visibility",
			yamlContent: `name: "test-repo"
description: "A test repository"
visibility: "invalid"`,
			structType:  "repository",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create test file
			testFile := filepath.Join(tempDir, "test.yaml")
			err := os.WriteFile(testFile, []byte(tt.yamlContent), 0644)
			if err != nil {
				t.Fatal(err)
			}

			// Test reading with validation based on struct type
			var validationErr error
			switch tt.structType {
			case "organization":
				var org Organization
				fd, openErr := os.Open(testFile)
				if openErr != nil {
					t.Fatal(openErr)
				}
				defer fd.Close()

				decoder := yaml.NewDecoder(fd)
				if validationErr = decoder.Decode(&org); validationErr == nil {
					org.SetDefaults()
					validationErr = validateStruct(&org)
				}

			case "repository":
				var repo Repository
				fd, openErr := os.Open(testFile)
				if openErr != nil {
					t.Fatal(openErr)
				}
				defer fd.Close()

				decoder := yaml.NewDecoder(fd)
				if validationErr = decoder.Decode(&repo); validationErr == nil {
					repo.SetDefaults()
					validationErr = validateStruct(&repo)
				}
			}

			if tt.expectError && validationErr == nil {
				t.Error("expected validation error but got nil")
			}
			if !tt.expectError && validationErr != nil {
				t.Errorf("unexpected validation error = %v", validationErr)
			}
		})
	}
}
