package queue

import "github.com/redis/go-redis/v9"

// redisQueue provides Redis-based message queue operations.
type redisQueue struct {
	client    *redis.Client
	queueName string
}

// NewRedisQueue creates a new Redis-based message queue repository.
//
// Parameters:
//   - c: the Redis client used to access the Redis queue.
//   - queueName: the name of the Redis queue.
//
// Returns:
//   - A Redis-based message queue repository.
func NewRedisQueue(c *redis.Client, queueName string) Repository {
	return &redisQueue{
		client:    c,
		queueName: queueName,
	}
}
