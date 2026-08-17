package queue

import (
	"context"
	"testing"

	redisPkg "github.com/HemlockPham7/common-libs/pkg/redis"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func TestQueueRepository_PopMessage(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string

		setupMock func(ctx context.Context) *redis.Client

		expectedErr error
		expectedRes []byte
	}{
		{
			name: "success",
			setupMock: func(ctx context.Context) *redis.Client {
				redisClient := redisPkg.InitMockRedis(t)

				err := redisClient.LPush(ctx, "test_queue", "test_message").Err()
				assert.NoError(t, err)

				return redisClient
			},

			expectedErr: nil,
			expectedRes: []byte("test_message"),
		},
		{
			name: "no message in queue",
			setupMock: func(ctx context.Context) *redis.Client {
				return redisPkg.InitMockRedis(t)
			},

			expectedErr: NoMessageError,
		},
		{
			name: "failed pop message due to closed Redis client",

			setupMock: func(ctx context.Context) *redis.Client {
				redisClient := redisPkg.InitMockRedis(t)
				redisClient.Close()
				return redisClient
			},

			expectedErr: redis.ErrClosed,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			ctx := t.Context()

			redisClient := tc.setupMock(ctx)
			repo := NewRedisQueue(redisClient, "test_queue")

			res, err := repo.PopMessage(ctx)
			assert.Equal(t, tc.expectedErr, err)
			assert.Equal(t, tc.expectedRes, res)
		})
	}
}
