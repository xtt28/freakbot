package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/xtt28/freakbot/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Require interface compliance.
var _ Connection = &gormConnection{}
var _ LeaderboardRepository = &gormLeaderboardRepository{}

func createCacheKey(boardID uint, userID string) string {
	return fmt.Sprint(boardID) + ":" + userID
}

type gormConnection struct {
	db *gorm.DB
	leaderboardRepository LeaderboardRepository
}

func (c *gormConnection) LeaderboardRepository() LeaderboardRepository {
	return c.leaderboardRepository
}

func NewGORMSQLiteConnection(dsn string) (*gormConnection, error) {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&model.Leaderboard{})
	db.AutoMigrate(&model.LeaderboardEntry{})

	gormConn := &gormConnection{db: db}
	gormConn.leaderboardRepository = &gormLeaderboardRepository{db: db, cache: map[string]uint{}}
	
	return gormConn, db.Error
}

type gormLeaderboardRepository struct {
	db *gorm.DB
	// Redis? Memcached? What's that?
	// The keys for the cache are formatted as leaderboard_id:user_id
	// We first try to read the cache before reading from the DB
	cache map[string]uint
}

func (r *gormLeaderboardRepository) CreateLeaderboard(guildID string) error {
	ctx := context.Background()

	return gorm.G[model.Leaderboard](r.db).Create(ctx, &model.Leaderboard{
		GuildID: guildID,
	})
}

func (r *gormLeaderboardRepository) GetLeaderboardID(guildID string) (uint, error) {
	ctx := context.Background()

	leaderboard, err := gorm.G[model.Leaderboard](r.db).Where("guild_id = ?", guildID).First(ctx)
	return leaderboard.ID, err
}

func (r *gormLeaderboardRepository) GetEntries(id uint, count uint, offset uint) ([]model.LeaderboardEntry, error) {
	ctx := context.Background()

	entries, err := gorm.G[model.LeaderboardEntry](r.db).
		Where("leaderboard_id = ?", id).
		Order("flagged_message_count DESC").
		Limit(int(count)).
		Offset(int(offset)).
		Find(ctx)
	return entries, err
}

func (r *gormLeaderboardRepository) GetEntryByUser(boardID uint, userID string) (model.LeaderboardEntry, error) {
	cacheKey := createCacheKey(boardID, userID)
	cachedCount, ok := r.cache[cacheKey]
	if ok {
		return model.LeaderboardEntry{
			LeaderboardID: boardID,
			UserID: userID,
			FlaggedMessageCount: cachedCount,
		}, nil
	}

	ctx := context.Background()

	entry, err := gorm.G[model.LeaderboardEntry](r.db).
		Where("leaderboard_id = ?", boardID).
		Where("user_id = ?", userID).
		First(ctx)

	if err == nil {
		r.cache[cacheKey] = entry.FlaggedMessageCount
	}
	
	return entry, err
}

func (r *gormLeaderboardRepository) IncrementUserFlaggedMessages(boardID uint, userID string) error {
	entry, err := r.GetEntryByUser(boardID, userID)
	isNotFound := errors.Is(err, gorm.ErrRecordNotFound)
	if err != nil && !isNotFound {
		return err
	}
	
	ctx := context.Background()

	newCount := entry.FlaggedMessageCount + 1
	cacheKey := createCacheKey(boardID, userID)

	r.cache[cacheKey] = newCount
	if isNotFound {
		return gorm.G[model.LeaderboardEntry](r.db).Create(ctx, &model.LeaderboardEntry{
			LeaderboardID: boardID,
			UserID: userID,
			FlaggedMessageCount: newCount,
		})
	} else {
		_, err := gorm.G[model.LeaderboardEntry](r.db).
			Where("user_id = ?", userID).
			Update(ctx, "flagged_message_count", newCount)
		return err
	}
}