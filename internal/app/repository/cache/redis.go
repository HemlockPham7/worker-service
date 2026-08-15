package cache

import "github.com/redis/go-redis/v9"

type redisDB struct {
	c *redis.Client
}

func NewRedisDB(c *redis.Client) DB {
	return &redisDB{c: c}
}
