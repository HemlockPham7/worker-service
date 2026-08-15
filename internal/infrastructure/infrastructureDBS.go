package infrastructure

import (
	"github.com/HemlockPham7/common-libs/pkg/common"
	redisPkg "github.com/HemlockPham7/common-libs/pkg/redis"
	"github.com/HemlockPham7/common-libs/pkg/sqldb"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func CreateDB(envPrefix string) *gorm.DB {
	dbClient, err := sqldb.NewClient(envPrefix)
	common.HandleError(err)

	return dbClient
}

func CreateRedisClient(envPrefix string) *redis.Client {
	redisClient, err := redisPkg.NewClient(envPrefix)
	common.HandleError(err)

	return redisClient
}
