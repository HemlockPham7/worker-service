package worker

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/rs/zerolog/log"
)

// pool manages a group of workers that process messages concurrently.
type pool struct { // quan li so luong worker trong mot pool
	handler      Handler
	numberWorker int
	messages     chan []byte
	errChan      chan *worker
	wg           *sync.WaitGroup
	nrClient     *newrelic.Application
}

// newPool creates and initializes a new worker pool.
//
// Parameters:
//   - ctx: the context passed to each worker.
//   - handler: the handler used by workers to process messages.
//   - numberWorker: the number of workers to create in the pool.
//   - nrClient: the New Relic application used to monitor worker operations.
//
// Returns:
//   - An initialized worker pool.
func newPool(ctx context.Context, handler Handler, numberWorker int, nrClient *newrelic.Application) *pool {
	messageChan := make(chan []byte, numberWorker)
	errorChan := make(chan *worker, numberWorker)

	initPool := &pool{
		handler:      handler,
		numberWorker: numberWorker,
		messages:     messageChan,
		errChan:      errorChan,
		wg:           &sync.WaitGroup{},
		nrClient:     nrClient,
	}

	initPool.init(ctx)
	return initPool
}

// init starts all workers in the pool and monitors workers for unexpected errors.
//
// Parameters:
//   - ctx: the context passed to each worker.
func (p *pool) init(ctx context.Context) {
	for i := 0; i < p.numberWorker; i++ {
		w := &worker{
			id:       i + 1,
			handler:  p.handler,
			messages: p.messages,
			errChan:  p.errChan,
			wg:       p.wg,
			nrClient: p.nrClient,
		}
		log.Info().Msgf("Starting worker %d", w.id)
		p.wg.Add(1)
		go w.run(ctx)
	}

	go func() {
		for w := range p.errChan {
			log.Error().Msgf("Worker %d exited with error %v", w.id, w.err)

			time.Sleep(1 * time.Second)
			log.Info().Msgf("Restarting worker %d ...", w.id)
			w.err = nil
			go w.run(ctx)
		}
	}()
}

// Consume submits a message to the worker pool for processing.
//
// Parameters:
//   - message: the message payload to be processed by an available worker.
func (p *pool) Consume(message []byte) {
	p.messages <- message
}

// Close shuts down the worker pool and waits for all workers to exit.
func (p *pool) Close() {
	close(p.messages)
	p.wg.Wait()
	close(p.errChan)
	log.Info().Msg("worker pool closed")
}

// worker processes messages from the worker pool.
type worker struct {
	id       int
	handler  Handler
	messages <-chan []byte
	err      error
	errChan  chan *worker
	wg       *sync.WaitGroup
	nrClient *newrelic.Application
}

// run continuously processes messages from the worker pool.
//
// Each message is processed within a separate New Relic transaction.
// If the worker panics, the panic is captured and the worker is reported
// through the error channel so that the pool can restart it.
//
// Parameters:
//   - ctx: the context used to control the lifetime of worker operations.
func (w *worker) run(ctx context.Context) {
	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(error); ok {
				w.err = err
			} else {
				w.err = fmt.Errorf("panic happened with %v", r)
			}
			w.errChan <- w
		} else {
			w.wg.Done()
		}
	}()

	for {
		msg, ok := <-w.messages
		if !ok {
			log.Info().Msgf("Worker %d is closing", w.id)
			return
		}

		// start a new relic transaction for each message
		txn := w.nrClient.StartTransaction("worker/handle-bookmark-import")
		txnCtx := newrelic.NewContext(ctx, txn)

		log.Debug().Msgf("Worker %d is processing message: %s", w.id, string(msg))
		err := w.handler.Handle(txnCtx, msg)
		if err != nil {
			log.Error().Err(err).Msgf("Worker %d failed to process message", w.id)
		} else {
			log.Info().Msgf("Worker %d processed successfully message: %s", w.id, string(msg))
		}

		txn.End()
	}
}
