package handler

import "github.com/bwmarrin/discordgo"

func (h *Handler) InteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	h.commandRegistry.HandleCommand(i.ApplicationCommandData().Name, s, i.Interaction)
}