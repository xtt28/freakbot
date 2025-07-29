package commands

import (
	"errors"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/xtt28/freakbot/internal/manifest"
	"gorm.io/gorm"
)

func (r *CommandRegistry) Leaderboard(s *discordgo.Session, i *discordgo.Interaction) {
	lbID, err := r.dbConn.LeaderboardRepository().GetLeaderboardID(i.GuildID)
	if err != nil {
		log.Println("could not get leaderboard ID for guild "+i.GuildID, err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err := r.dbConn.LeaderboardRepository().CreateLeaderboard(i.GuildID)
			if err != nil {
				log.Println("could not create leaderboard for guild "+i.GuildID, err)
				s.InteractionRespond(i, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "There's no leaderboard for this guild and we couldn't create one. This shouldn't be happening... make an issue on GitHub " + manifest.GitHubURL,
					},
				})
			}
			s.InteractionRespond(i, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "We just created the leaderboard for this guild. Please try again.",
				},
			})
		} else {
			s.InteractionRespond(i, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Couldn't find this guild's leaderboard. This shouldn't be happening... make an issue on GitHub " + manifest.GitHubURL,
				},
			})
		}
		return
	}
	lb, err := r.dbConn.LeaderboardRepository().GetEntries(lbID, 10, 0)
	if err != nil {
		log.Println("could not get leaderboard entries for guild "+i.GuildID, err)
		s.InteractionRespond(i, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Couldn't find the leaderboard entries for this guild. This shouldn't be happening... try again or contact a bot administrator for help.",
			},
		})
		return
	}

	responseBody := "Top 10 freakiest members\n\n"
	for i, v := range lb {
		responseBody += fmt.Sprintf("**%d.** <@%s> (%d freaky messages)\n", i+1, v.UserID, v.FlaggedMessageCount)
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
