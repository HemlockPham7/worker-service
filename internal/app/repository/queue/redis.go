package queue

import "github.com/redis/go-redis/v9"

type redisQueue struct {
	client    *redis.Client
	queueName string
}

func NewRedisQueue(c *redis.Client, queueName string) Repository {
	return &redisQueue{
		client:    c,
		queueName: queueName,
	}
}
