package worker

import (
	"context"

	"github.com/HemlockPham7/worker-service/internal/app/service/bookmark"
)

// Handler defines the interface for processing worker messages.
type Handler interface {
	Handle(ctx context.Context, message []byte) error
}

type handler struct {
	bookmarkService bookmark.Service
}

// NewHandler creates a new worker message handler.
//
// Parameters:
//   - bookmarkService: the service used to process bookmark-related jobs.
//
// Returns:
//   - A configured worker message handler.
func NewHandler(bookmarkService bookmark.Service) Handler {
	return &handler{
		bookmarkService: bookmarkService,
	}
}
