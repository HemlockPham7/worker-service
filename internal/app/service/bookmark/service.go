package bookmark

import (
	"context"

	"github.com/HemlockPham7/common-libs/pkg/utils"
	"github.com/HemlockPham7/worker-service/internal/app/repository/bookmark"
	"github.com/HemlockPham7/worker-service/internal/app/service/queue"
)

const codeLength = 8

//go:generate mockery --name Service --filename service.go --outpkg mock_bookmark
type Service interface {
	CreateBatchBookmarks(ctx context.Context, userId string, bookmarkList []*queue.ImportBookmarkInput) error
}

type bookmarkService struct {
	bookmarkRepository bookmark.Repository
	codeGenerator utils.GenCode
}

func NewService(bookmarkRepository bookmark.Repository, codeGenerator utils.GenCode) Service {
	return &bookmarkService{
		bookmarkRepository: bookmarkRepository,
		codeGenerator: codeGenerator,
	}
}