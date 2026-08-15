package worker

import (
	"context"

	"github.com/HemlockPham7/worker-service/internal/app/service/bookmark"
)

type Handler interface {
	Handle(ctx context.Context, message []byte) error
}

type handler struct {
	bookmarkService bookmark.Service
}

func NewHandler(bookmarkService bookmark.Service) Handler {
	return &handler{
		bookmarkService: bookmarkService,
	}
}
