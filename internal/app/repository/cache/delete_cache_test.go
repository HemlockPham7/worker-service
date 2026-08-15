package cache

import (
	"context"
	"testing"
	"time"

	redisPkg "github.com/HemlockPham7/common-libs/pkg/redis"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func TestRedisDB_DeleteCache(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string

		setupMock func(ctx context.Context) *redis.Client

		expectedError error

		verifyFunc func(ctx context.Context, redisClient *redis.Client)
	}{
		{
			name: "successful cache deletion",

			setupMock: func(ctx context.Context) *redis.Client {
				redisClient := redisPkg.InitMockRedis(t)
				redisClient.Set(ctx, "cache_key", "cache_value", time.Hour)
				return redisClient
			},

			expectedError: nil,

			verifyFunc: func(ctx context.Context, redisClient *redis.Client) {
				cacheValue, err := redisClient.Get(ctx, "cache_key").Result()
				assert.Equal(t, redis.Nil, err)
				assert.Equal(t, cacheValue, "")
			},
		},
		{
			name: "failed due to redis client close",

			setupMock: func(ctx context.Context) *redis.Client {
				redisClient := redisPkg.InitMockRedis(t)
				_ = redisClient.Close()
				return redisClient
			},

			expectedError: redis.ErrClosed,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ctx := t.Context()

			redisClient := tc.setupMock(ctx)
			storage := NewRedisDB(redisClient)

			err := storage.DeleteCache(ctx, "cache_key")
			assert.Equal(t, tc.expectedError, err)

			if err == nil {
				tc.verifyFunc(ctx, redisClient)
			}
		})
	}
}
