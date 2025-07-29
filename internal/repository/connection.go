package repository

type Connection interface {
	// LeaderboardRepository returns the connection's instance of the
	// leaderboard repository.
	LeaderboardRepository() LeaderboardRepository
}
