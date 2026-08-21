package bookmark

import (
	"context"
	"fmt"

	"github.com/HemlockPham7/worker-service/internal/app/model"
	"github.com/HemlockPham7/worker-service/internal/app/service/queue"
	"github.com/newrelic/go-agent/v3/newrelic"
)

const (
	getBookmarksCacheGroupKeyFormat = "get_bookmarks_%s"
)

func (s *bookmarkService) CreateBatchBookmarks(ctx context.Context, userId string, bookmarkList []*queue.ImportBookmarkInput) error {
	txn := newrelic.FromContext(ctx)
	span := txn.StartSegment("CreateBatchBookmarks_BookmarkService")
	defer span.End()

	err := s.cacheRepository.DeleteCache(ctx, fmt.Sprintf(getBookmarksCacheGroupKeyFormat, userId))
	if err != nil {
		return err
	}

	bookmarks := make([]*model.Bookmark, len(bookmarkList))

	for i, input := range bookmarkList {
		code, err := s.codeGenerator.GenerateCode(codeLength)
		if err != nil {
			return err
		}

		bookmarks[i] = &model.Bookmark{
			Description: input.Description,
			URL:         input.URL,
			UserID:      userId,
			Code:        code,
		}
	}

	err = s.bookmarkRepository.CreateBatchBookmarks(ctx, bookmarks)
	if err != nil {
		return err
	}
	return nil
}
