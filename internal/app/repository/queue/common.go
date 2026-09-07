package queue

import "context"

// Repository defines the data access operations for the message queue.
//
//go:generate mockery --name Repository --filename common.go --outpkg mock_queue
type Repository interface {
	PopMessage(ctx context.Context) ([]byte, error)
}
