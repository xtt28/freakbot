package handler

import "github.com/bwmarrin/discordgo"

func (h *Handler) GuildCreate(s *discordgo.Session, g *discordgo.GuildCreate) {
	_, err := h.dbConn.LeaderboardRepository().GetLeaderboardID(g.ID)
	if err != nil {
		h.dbConn.LeaderboardRepository().CreateLeaderboard(g.ID)
	}
}