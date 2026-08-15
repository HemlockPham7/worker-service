package bookmark

import (
	"context"

	"github.com/HemlockPham7/worker-service/internal/app/model"
	"github.com/HemlockPham7/worker-service/internal/app/service/queue"
)

func (s *bookmarkService) CreateBatchBookmarks(ctx context.Context, userId string, bookmarkList []*queue.ImportBookmarkInput) error {
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

	err := s.bookmarkRepository.CreateBatchBookmarks(ctx, bookmarks)
	if err != nil {
		return err
	}
	return nil
}
