package cache

import (
	"context"
)

//go:generate mockery --name DB --filename common.go --outpkg mock_cache
type DB interface {
	DeleteCache(ctx context.Context, key string) error
}
