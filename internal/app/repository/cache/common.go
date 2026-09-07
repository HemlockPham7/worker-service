package cache

import (
	"context"
)

// DB defines the operations for managing cached data.
//
//go:generate mockery --name DB --filename common.go --outpkg mock_cache
type DB interface {
	DeleteCache(ctx context.Context, key string) error
}
