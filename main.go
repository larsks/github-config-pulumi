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

func main() {
	importMode := strings.ToLower(os.Getenv("PULUMI_IMPORT")) == "true"

	orgSpec, err := readers.ReadOrganization()
	if err != nil {
		log.Fatalf("failed to read organization: %v", err)
	}

	members, err := readers.ReadMembers()
	if err != nil {
		log.Fatalf("failed to read organization members: %v", err)
	}

	teams, err := readers.ReadTeams()
	if err != nil {
		log.Fatalf("failed to read teams: %v", err)
	}

	_, err = readers.ReadLabels()
	if err != nil {
		log.Fatalf("failed to read labels: %v", err)
	}

	pulumi.Run(func(ctx *pulumi.Context) error {
		for _, memberSpec := range members {
			var options []pulumi.ResourceOption
			if importMode {
				options = append(options, pulumi.Import(pulumi.ID(fmt.Sprintf("%s:%s", orgSpec.Name, memberSpec.Name))))
			}

			_, err := github.NewMembership(ctx, fmt.Sprintf("github-member-%s", memberSpec.Name), &github.MembershipArgs{
				Username: pulumi.String(memberSpec.Name),
				Role:     pulumi.String(memberSpec.Role),
			}, options...)
			if err != nil {
				return err
			}
		}

		for _, teamSpec := range teams {
			var options []pulumi.ResourceOption
			if importMode {
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
			if importMode {
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
	})
}
