package cache

import (
	"context"

	"github.com/newrelic/go-agent/v3/newrelic"
)

// DeleteCache deletes the specified cache group from Redis.
//
// Parameters:
//   - ctx: the context used to control the lifetime of the operation.
//   - key: the key of the cache group to delete.
//
// Returns:
//   - An error if the cache cannot be deleted.
func (r *redisDB) DeleteCache(ctx context.Context, key string) error {
	txn := newrelic.FromContext(ctx)
	span := txn.StartSegment("DeleteCache_CacheRepository")
	defer span.End()

	return r.c.Del(ctx, key).Err()
}
