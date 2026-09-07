package infrastructure

import (
	"github.com/HemlockPham7/common-libs/pkg/common"
	redisPkg "github.com/HemlockPham7/common-libs/pkg/redis"
	"github.com/HemlockPham7/common-libs/pkg/sqldb"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// CreateDB creates a GORM database client.
//
// Parameters:
//   - envPrefix: the environment variable prefix used to load database configuration.
//
// Returns:
//   - A configured GORM database client.
//
// Panics:
//   - If the database client cannot be created.
func CreateDB(envPrefix string) *gorm.DB {
	dbClient, err := sqldb.NewClient(envPrefix)
	common.HandleError(err)

	return dbClient
}

// CreateRedisClient creates a Redis client using the specified environment configuration prefix.
//
// Parameters:
//   - envPrefix: the environment variable prefix used to load Redis configuration.
//
// Returns:
//   - A configured Redis client.
//
// Panics:
//   - If the Redis client cannot be created.
func CreateRedisClient(envPrefix string) *redis.Client {
	redisClient, err := redisPkg.NewClient(envPrefix)
	common.HandleError(err)

	return redisClient
}
