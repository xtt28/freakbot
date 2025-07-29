package model

import "gorm.io/gorm"

// Leaderboard represents the leaderboard unique to each guild in which the bot
// is present.
type Leaderboard struct {
	gorm.Model
	GuildID string `gorm:"uniqueIndex"`
}

// LeaderboardEntry represents a single entry in the per-guild leaderboard.
type LeaderboardEntry struct {
	gorm.Model
	Leaderboard         Leaderboard
	LeaderboardID       uint
	UserID              string
	FlaggedMessageCount uint
}
