package repository

import "github.com/xtt28/freakbot/internal/model"

type LeaderboardRepository interface {
	// GetLeaderboardEntries returns the given amount of leaderboard entries for
	// the given guild, offset by the given count. The first parameter is the
	// leaderboard ID, the second parameter is the amount of leaderboard entries
	// to return, and the third parameter is the offset. Leaderboard entries are
	// always returned in descending order, from most flagged messages to least.
	GetEntries(uint, uint, uint) ([]model.LeaderboardEntry, error)
	// GetEntryByUser returns the entry in the given leaderboard specific to the
	// user with the given ID.
	GetEntryByUser(uint, string) (model.LeaderboardEntry, error)
	// IncrementUserFlaggedMessages increments the flagged message count for the
	// given user in the given leaderboard.
	IncrementUserFlaggedMessages(uint, string) error
}