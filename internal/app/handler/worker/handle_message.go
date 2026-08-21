package worker

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/HemlockPham7/worker-service/internal/app/service/queue"
	"github.com/newrelic/go-agent/v3/newrelic"
)

var ErrUnmarshalMessage = errors.New("failed to unmarshal message")

func (h *handler) Handle(ctx context.Context, message []byte) error {
	txn := newrelic.FromContext(ctx)
	span := txn.StartSegment("Handle_WorkerHandler")
	defer span.End()

	input := &queue.ImportMessage{}
	err := json.Unmarshal(message, input)
	if err != nil {
		return ErrUnmarshalMessage
	}

	txn.AddAttribute("user_id", input.UID)
	txn.AddAttribute("bookmark_count", len(input.Bookmarks))

	err = h.bookmarkService.CreateBatchBookmarks(ctx, input.UID, input.Bookmarks)
	if err != nil {
		return err
	}
	return nil
}
