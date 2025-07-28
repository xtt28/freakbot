package handler

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func (h *Handler) MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	content := m.Content
	if len(content) < 3 {
		return
	}

	flag, err := h.classifierService.IsFlagged(content)
	if err != nil {
		log.Println("could not classify message", err)
		return
	}

	log.Println(m.Content)
	
	if flag {
		log.Printf("user %s/guild %s :: flagged :: %s\n", m.Author.ID, m.GuildID, m.Content)
		lbid, err := h.dbConn.LeaderboardRepository().GetLeaderboardID(m.GuildID)
		if err != nil {
			log.Println("could not get guild leaderboard id", err)
			return
		}
		err = h.dbConn.LeaderboardRepository().IncrementUserFlaggedMessages(lbid, m.Author.ID)
		if err != nil {
			log.Println("could not increment user flags", err)
			return
		}
	}
}