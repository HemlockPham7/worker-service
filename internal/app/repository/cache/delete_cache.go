package cache

import "context"

func (r *redisDB) DeleteCache(ctx context.Context, key string) error {
	return r.c.Del(ctx, key).Err()
}
