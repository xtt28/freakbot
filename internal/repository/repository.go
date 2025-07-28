package repository

import "github.com/xtt28/freakbot/internal/model"

type LeaderboardRepository interface {
	// CreateLeaderboard creates a leaderboard for the server with given ID.
	CreateLeaderboard(guildID string) error
	// GetLeaderboardID gets the ID of the leaderboard for the given server.
	GetLeaderboardID(guildID string) (uint, error)
	// GetLeaderboardEntries returns the given amount of leaderboard entries for
	// the given guild, offset by the given count. The first parameter is the
	// leaderboard ID, the second parameter is the amount of leaderboard entries
	// to return, and the third parameter is the offset. Leaderboard entries are
	// always returned in descending order, from most flagged messages to least.
	GetEntries(leaderboardID uint, count uint, offset uint) ([]model.LeaderboardEntry, error)
	// GetEntryByUser returns the entry in the given leaderboard specific to the
	// user with the given ID.
	GetEntryByUser(leaderboardID uint, userID string) (model.LeaderboardEntry, error)
	// IncrementUserFlaggedMessages increments the flagged message count for the
	// given user in the given leaderboard.
	IncrementUserFlaggedMessages(leaderboardID uint, userID string) error
}