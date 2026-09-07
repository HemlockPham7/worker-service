package bookmark

import (
	"context"

	"github.com/HemlockPham7/worker-service/internal/app/model"
	"gorm.io/gorm"
)

// Repository defines the data access operations for bookmarks.
//
//go:generate mockery --name Repository --filename repo.go --outpkg mock_bookmark
type Repository interface {
	CreateBatchBookmarks(ctx context.Context, bookmarks []*model.Bookmark) error
}

type bookmarkRepository struct {
	db *gorm.DB
}

// NewRepository creates a new bookmark repository.
//
// Parameters:
//   - db: the GORM database client used to access bookmark data.
//
// Returns:
//   - A configured bookmark repository.
func NewRepository(db *gorm.DB) Repository {
	return &bookmarkRepository{db: db}
}
