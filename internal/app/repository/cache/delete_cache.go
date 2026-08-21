package cache

import (
	"context"

	"github.com/newrelic/go-agent/v3/newrelic"
)

func (r *redisDB) DeleteCache(ctx context.Context, key string) error {
	txn := newrelic.FromContext(ctx)
	span := txn.StartSegment("DeleteCache_CacheRepository")
	defer span.End()

	return r.c.Del(ctx, key).Err()
}
