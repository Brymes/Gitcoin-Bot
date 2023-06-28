package gitcoin

import (
	"Brymes-Gitcoin-Bot/config"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"math"
	"strconv"
	"time"
)

func (bounty BountiesResponse) createBountyEmbed() *discordgo.MessageEmbed {
	embed := &discordgo.MessageEmbed{
		Color:       0x5f3267,
		Title:       fmt.Sprintf("$%f Bounty; %s", roundValue(bounty.ValueTrueUsd), bounty.Keywords),
		Description: bounty.Title,
		URL:         bounty.Url,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Token Name",
				Value:  bounty.TokenName,
				Inline: false,
			},
			{
				Name:   "Token Value",
				Value:  bounty.ValueInToken,
				Inline: false,
			},
			{
				Name:   "Github URL",
				Value:  bounty.GithubUrl,
				Inline: false,
			},
			{
				Name:   "Experience Level",
				Value:  bounty.ExperienceLevel,
				Inline: true,
			},
			{
				Name:   "Number of Applicants",
				Value:  fmt.Sprintf("%d", bounty.InterestedCount),
				Inline: true,
			},
			{
				Name:   "Deadline",
				Value:  bounty.ExpiresDate.Format(time.RFC1123),
				Inline: true,
			},
		},
		Timestamp: time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
		Footer:    &config.MessageFooter,
	}

	return embed
}

func (grant GrantResponse) createGrantEmbed() *discordgo.MessageEmbed {
	embed := &discordgo.MessageEmbed{
		Color:       0x5f3267,
		Title:       fmt.Sprintf("$%f Grant @", roundValue(grant.AmountReceived)),
		Description: grant.Title,
		URL:         fmt.Sprintf("https://bounties.gitcoin.co%s", grant.Url),
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Is Verified?",
				Value:  fmt.Sprintf("%v", grant.Verified),
				Inline: true,
			},
			{
				Name:   "Is Active?",
				Value:  fmt.Sprintf("%v", grant.Active),
				Inline: true,
			},
			{
				Name:   "Contributor Count",
				Value:  fmt.Sprintf("%d", grant.PositiveRoundContributorCount),
				Inline: false,
			},
			{
				Name:   "Github URL",
				Value:  grant.GithubProjectUrl,
				Inline: false,
			},
			{
				Name:   "Reference URL",
				Value:  grant.ReferenceUrl,
				Inline: false,
			},
			{
				Name:   "Twitter",
				Value:  grant.TwitterHandle1,
				Inline: false,
			},
		},
		Timestamp: time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
		Footer:    &config.MessageFooter,
	}

	return embed
}

func roundValue(figure string) float64 {
	val, _ := strconv.ParseFloat(figure, 64)

	return math.Round(val*10) / 10
}
