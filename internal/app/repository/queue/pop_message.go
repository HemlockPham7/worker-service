package queue

import (
	"context"
	"errors"

	"github.com/redis/go-redis/v9"
)

var NoMessageError = errors.New("no message")

func (r *redisQueue) PopMessage(ctx context.Context) ([]byte, error) {
	msg, err := r.client.RPop(ctx, r.queueName).Bytes()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, NoMessageError
		}
		return nil, err
	}
	return msg, nil
}
