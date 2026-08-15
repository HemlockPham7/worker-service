package bookmark

import (
	"context"

	"github.com/HemlockPham7/worker-service/internal/app/model"
	"gorm.io/gorm"
)

//go:generate mockery --name Repository --filename repo.go --outpkg mock_bookmark
type Repository interface {
	CreateBatchBookmarks(ctx context.Context, bookmarks []*model.Bookmark) error
}

type bookmarkRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &bookmarkRepository{db: db}
}
