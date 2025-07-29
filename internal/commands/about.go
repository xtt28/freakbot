package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/xtt28/freakbot/internal/manifest"
)

func (r *CommandRegistry) About(s *discordgo.Session, i *discordgo.Interaction) {
	s.InteractionRespond(i, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       "About this bot",
					Description: "Information as presented in the app manifest",
					Fields: []*discordgo.MessageEmbedField{
						{
							Name:  "Name",
							Value: manifest.Name,
						},
						{
							Name:  "GitHub URL",
							Value: manifest.GitHubURL,
						},
						{
							Name:  "Identifier",
							Value: manifest.Iden,
						},
					},
				},
			},
		},
	})
}
