package queue

import (
	"context"
	"errors"

	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/redis/go-redis/v9"
)

var NoMessageError = errors.New("no message")

// PopMessage retrieves and removes the next message from the Redis queue.
//
// Parameters:
//   - ctx: the context used to control the lifetime of the operation.
//
// Returns:
//   - The message payload retrieved from the queue.
//   - NoMessageError if the queue is empty.
//   - An error if the message cannot be retrieved from Redis.
func (r *redisQueue) PopMessage(ctx context.Context) ([]byte, error) {
	txn := newrelic.FromContext(ctx)
	span := txn.StartSegment("PopMessage_QueueRepository")
	defer span.End()

	msg, err := r.client.RPop(ctx, r.queueName).Bytes()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, NoMessageError
		}
		return nil, err
	}
	return msg, nil
}
