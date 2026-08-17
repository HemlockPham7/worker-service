package worker

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/HemlockPham7/worker-service/internal/app/service/queue"
)

var ErrUnmarshalMessage = errors.New("failed to unmarshal message")

func (h *handler) Handle(ctx context.Context, message []byte) error {
	input := &queue.ImportMessage{}
	err := json.Unmarshal(message, input)
	if err != nil {
		return ErrUnmarshalMessage
	}

	err = h.bookmarkService.CreateBatchBookmarks(ctx, input.UID, input.Bookmarks)
	if err != nil {
		return err
	}
	return nil
}
