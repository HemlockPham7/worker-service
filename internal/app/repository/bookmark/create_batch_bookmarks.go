package bookmark

import (
	"context"

	"github.com/HemlockPham7/common-libs/pkg/dbutils"
	"github.com/HemlockPham7/worker-service/internal/app/model"
	"github.com/newrelic/go-agent/v3/newrelic"
	"gorm.io/gorm"
)

// CreateBatchBookmarks creates multiple bookmarks in a single database transaction.
//
// Parameters:
//   - ctx: the context used to control the lifetime of the database operation.
//   - bookmarks: the bookmarks to create.
//
// Returns:
//   - An error if any bookmark cannot be created.
//   - A database error converted by dbutils.CatchDBError if the transaction fails.
func (r *bookmarkRepository) CreateBatchBookmarks(ctx context.Context, bookmarks []*model.Bookmark) error {
	txn := newrelic.FromContext(ctx)
	span := txn.StartSegment("CreateBatchBookmarks_BookmarkRepository")
	defer span.End()

	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, bookmark := range bookmarks {
			if err := tx.Create(bookmark).Error; err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return dbutils.CatchDBError(err)
	}
	return nil
}
