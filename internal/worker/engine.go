package worker

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/HemlockPham7/worker-service/internal/app/repository/queue"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/rs/zerolog/log"
)

// Engine defines the interface for starting the worker engine.
type Engine interface {
	Start(ctx context.Context)
}

// Handler defines the interface for processing worker messages.
type Handler interface {
	Handle(ctx context.Context, message []byte) error
}

type engine struct {
	queue    queue.Repository
	handler  Handler
	run      bool
	sigChan  chan os.Signal
	nrClient *newrelic.Application
}

// NewEngine creates a new worker engine.
//
// Parameters:
//   - queue: the repository used to retrieve messages from the queue.
//   - handler: the handler used to process queued messages.
//   - nrClient: the New Relic application used to monitor worker operations.
//
// Returns:
//   - A configured worker engine.
func NewEngine(queue queue.Repository, handler Handler, nrClient *newrelic.Application) Engine {
	return &engine{
		queue:    queue,
		handler:  handler,
		run:      false,
		sigChan:  make(chan os.Signal, 1),
		nrClient: nrClient,
	}
}

const (
	// intervalDelay defines the delay between queue polling attempts when no message is available.
	intervalDelay = 1 * time.Second

	// numberOfWorker defines the number of workers processing queued messages concurrently.
	numberOfWorker = 4
)

// Start starts the worker engine and continuously consumes messages from the queue.
//
// The engine polls the queue at a fixed interval when no message is available,
// processes messages using a worker pool, and stops gracefully when it receives
// SIGINT or SIGTERM.
//
// Parameters:
//   - ctx: the context used to control the lifetime of worker operations.
func (e *engine) Start(ctx context.Context) {
	log.Info().Msg("Starting worker engine")

	workerPool := newPool(ctx, e.handler, numberOfWorker, e.nrClient)
	signal.Notify(e.sigChan, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	e.run = true
	for e.run {
		select {
		case sig := <-e.sigChan:
			log.Info().Msgf("Received signal: %s", sig.String())
			e.run = false
		default:
			// pop message
			msg, err := e.queue.PopMessage(ctx)
			if err != nil {
				if errors.Is(err, queue.NoMessageError) {
					time.Sleep(intervalDelay)
					continue
				}

				log.Error().Err(err).Msg("Failed to pop message")
				time.Sleep(intervalDelay)
				continue
			}

			// handle message
			workerPool.Consume(msg)
		}
	}
	workerPool.Close()
}
