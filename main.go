package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/pulumi/pulumi-github/sdk/v5/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github-config-pulumi/internal/readers"
)

type (
	OrgManager struct {
		Name       string
		ImportMode bool
	}
)

func main() {
	orgSpec, err := readers.ReadOrganization()
	if err != nil {
		log.Fatalf("failed to read organization: %v", err)
	}

	om := &OrgManager{
		Name:       orgSpec.Name,
		ImportMode: strings.ToLower(os.Getenv("PULUMI_IMPORT")) == "true",
	}

	members, err := readers.ReadMembers()
	if err != nil {
		log.Fatalf("failed to read organization members: %v", err)
	}

	teams, err := readers.ReadTeams()
	if err != nil {
		log.Fatalf("failed to read teams: %v", err)
	}

	labels, err := readers.ReadLabels()
	if err != nil {
		log.Fatalf("failed to read labels: %v", err)
	}

	repos, err := readers.ReadRepositories()
	if err != nil {
		log.Fatalf("failed to read repositories: %v", err)
	}

	pulumi.Run(func(ctx *pulumi.Context) error {
		if err := om.realizeMembers(ctx, members); err != nil {
			log.Fatalf("failed to manage members: %v", err)
		}

		if err := om.realizeTeams(ctx, teams); err != nil {
			log.Fatalf("failed to manage teams: %v", err)
		}

		if err := om.realizeRepos(ctx, repos, labels); err != nil {
			log.Fatalf("failed to manage repositories: %v", err)
		}

		return nil
	})
}

func (om *OrgManager) realizeMembers(ctx *pulumi.Context, members []*readers.Member) error {
	for _, memberSpec := range members {
		var options []pulumi.ResourceOption
		if om.ImportMode {
			options = append(options, pulumi.Import(pulumi.ID(fmt.Sprintf("%s:%s", om.Name, memberSpec.Name))))
		}

		_, err := github.NewMembership(ctx, fmt.Sprintf("github-member-%s", memberSpec.Name), &github.MembershipArgs{
			Username: pulumi.String(memberSpec.Name),
			Role:     pulumi.String(memberSpec.Role),
		}, options...)
		if err != nil {
			return err
		}
	}
	return nil
}

func (om *OrgManager) realizeTeams(ctx *pulumi.Context, teams []*readers.Team) error {
	for _, teamSpec := range teams {
		var options []pulumi.ResourceOption
		if om.ImportMode {
			options = append(options, pulumi.Import(pulumi.ID(teamSpec.Name)))
		}

		team, err := github.NewTeam(ctx, fmt.Sprintf("github-team-%s", teamSpec.Name), &github.TeamArgs{
			Name:        pulumi.String(teamSpec.Name),
			Description: pulumi.String(teamSpec.Description),
			Privacy:     pulumi.String(teamSpec.Privacy),
		}, options...)
		if err != nil {
			return err
		}

		var teamMembers github.TeamMembersMemberArray
		for _, member := range teamSpec.Members {
			teamMembers = append(teamMembers, github.TeamMembersMemberArgs{
				Username: pulumi.String(member.Name),
				Role:     pulumi.String(member.Role),
			})
		}

		var teamMembersOptions []pulumi.ResourceOption
		teamMembersOptions = append(teamMembersOptions, pulumi.DependsOn([]pulumi.Resource{team}))
		if om.ImportMode {
			teamMembersOptions = append(teamMembersOptions, pulumi.Import(pulumi.ID(teamSpec.Name)))
		}

		_, err = github.NewTeamMembers(ctx, fmt.Sprintf("github-team-%s-members", teamSpec.Name), &github.TeamMembersArgs{
			TeamId:  team.ID(),
			Members: teamMembers,
		}, teamMembersOptions...)
		if err != nil {
			return err
		}
	}

	return nil
}

func (om *OrgManager) realizeRepos(ctx *pulumi.Context, repos []*readers.Repository, labels []*readers.Label) error {
	for _, repoSpec := range repos {
		var options []pulumi.ResourceOption
		if om.ImportMode {
			options = append(options, pulumi.Import(pulumi.ID(repoSpec.Name)))
		}

		_, err := github.NewRepository(ctx, fmt.Sprintf("github-repo-%s", repoSpec.Name), &github.RepositoryArgs{
			Name: pulumi.String(repoSpec.Name),

			AllowAutoMerge:      pulumi.Bool(*repoSpec.AllowAutoMerge),
			AutoInit:            pulumi.Bool(*repoSpec.AutoInit),
			Description:         pulumi.String(repoSpec.Description),
			HasDownloads:        pulumi.Bool(*repoSpec.HasDownloads),
			HasIssues:           pulumi.Bool(*repoSpec.HasIssues),
			HasProjects:         pulumi.Bool(*repoSpec.HasProjects),
			HasWiki:             pulumi.Bool(*repoSpec.HasWiki),
			IsTemplate:          pulumi.Bool(*repoSpec.IsTemplate),
			Visibility:          pulumi.String(repoSpec.Visibility),
			VulnerabilityAlerts: pulumi.Bool(*repoSpec.VulnerabilityAlerts),
		}, options...)
		if err != nil {
			return err
		}
	}

	return nil
}
