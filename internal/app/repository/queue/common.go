package queue

import "context"

//go:generate mockery --name Repository --filename common.go --outpkg mock_queue
type Repository interface {
	PopMessage(ctx context.Context) ([]byte, error)
}
