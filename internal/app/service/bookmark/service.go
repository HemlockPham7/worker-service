package bookmark

import (
	"context"

	"github.com/HemlockPham7/common-libs/pkg/utils"
	"github.com/HemlockPham7/worker-service/internal/app/repository/bookmark"
	"github.com/HemlockPham7/worker-service/internal/app/repository/cache"
	"github.com/HemlockPham7/worker-service/internal/app/service/queue"
)

const codeLength = 8

// Service defines the business logic for batch bookmark operations.
//
//go:generate mockery --name Service --filename service.go --outpkg mock_bookmark
type Service interface {
	CreateBatchBookmarks(ctx context.Context, userId string, bookmarkList []*queue.ImportBookmarkInput) error
}

type bookmarkService struct {
	bookmarkRepository bookmark.Repository
	cacheRepository    cache.DB
	codeGenerator      utils.GenCode
}

// NewService creates a new bookmark service.
//
// Parameters:
//   - bookmarkRepository: the repository used to persist bookmarks.
//   - cacheRepository: the cache database used to invalidate bookmark caches.
//   - codeGenerator: the code generator used to generate bookmark codes.
//
// Returns:
//   - A configured bookmark service.
func NewService(bookmarkRepository bookmark.Repository, cacheRepository cache.DB, codeGenerator utils.GenCode) Service {
	return &bookmarkService{
		bookmarkRepository: bookmarkRepository,
		cacheRepository:    cacheRepository,
		codeGenerator:      codeGenerator,
	}
}
