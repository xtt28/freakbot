package repository

import (
	"github.com/xtt28/freakbot/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Require interface compliance.
var _ Connection = &gormConnection{}
var _ LeaderboardRepository = &gormLeaderboardRepository{}

type gormConnection struct {
	db *gorm.DB
	leaderboardRepository LeaderboardRepository

	// Redis? Memcached? What's that?
	// The keys for the cache are formatted as leaderboard_id:user_id
	// We first try to read the cache before reading from the DB
	cache map[string]uint
}

func (c *gormConnection) LeaderboardRepository() LeaderboardRepository {
	return c.leaderboardRepository
}

func NewGORMSQLiteConnection(dsn string) (*gormConnection, error) {
	db, err := gorm.Open(sqlite.Open(dsn))
	if err != nil {
		return nil, err
	}

	gormConn := &gormConnection{db: db}
	gormConn.leaderboardRepository = &gormLeaderboardRepository{db: db}
	
	return gormConn, db.Error
}

type gormLeaderboardRepository struct {
	db *gorm.DB
}

func (r *gormLeaderboardRepository) CreateLeaderboard(string) error {
	return nil
}

func (r *gormLeaderboardRepository) GetLeaderboardID(string) (uint, error) {
	return 0, nil
}

func (r *gormLeaderboardRepository) GetEntries(uint, uint, uint) ([]model.LeaderboardEntry, error) {
	return []model.LeaderboardEntry{}, nil
}

func (r *gormLeaderboardRepository) GetEntryByUser(uint, string) (model.LeaderboardEntry, error) {
	return model.LeaderboardEntry{}, nil
}

func (r *gormLeaderboardRepository) IncrementUserFlaggedMessages(uint, string) error {
	return nil
}