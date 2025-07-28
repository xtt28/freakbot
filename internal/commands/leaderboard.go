package commands

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/xtt28/freakbot/internal/manifest"
)

func (r *CommandRegistry) Leaderboard(s *discordgo.Session, i *discordgo.Interaction) {
	lbID, err := r.dbConn.LeaderboardRepository().GetLeaderboardID(i.GuildID)
	if err != nil {
		log.Println("could not get leaderboard ID for guild " + i.GuildID, err)
		s.InteractionRespond(i, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Couldn't find the leaderboard for this guild. This shouldn't be happening... try again or contact a bot administrator for help.",
			},
		})
	}
	lb, err := r.dbConn.LeaderboardRepository().GetEntries(lbID, 10, 0)
	if err != nil {
		log.Println("could not get leaderboard entries for guild " + i.GuildID, err)
		s.InteractionRespond(i, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Couldn't find the leaderboard entries for this guild. This shouldn't be happening... try again or contact a bot administrator for help.",
			},
		})
	}

	responseBody := "Top 10 freakiest members\n\n"
	for i, v := range lb {
		responseBody += fmt.Sprintf("**%d.** <@%s> (%d freaky messages)\n", i + 1, v.UserID, v.FlaggedMessageCount)
	}

	s.InteractionRespond(i, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       "Freaky leaderboard",
					Description: responseBody,

					Footer: &discordgo.MessageEmbedFooter{
						Text: fmt.Sprintf("Powered by %s (%s)", manifest.Name, manifest.GitHubURL),
					},
				},
			},
		},
	})
}
