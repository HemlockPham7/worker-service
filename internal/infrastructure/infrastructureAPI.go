package infrastructure

import (
	"context"

	"github.com/HemlockPham7/common-libs/pkg/common"
	"github.com/HemlockPham7/common-libs/pkg/logger"
	"github.com/HemlockPham7/common-libs/pkg/utils"
	"github.com/HemlockPham7/worker-service/internal/api"
	bookmarkHdl "github.com/HemlockPham7/worker-service/internal/app/handler/worker"
	bookmarkRepo "github.com/HemlockPham7/worker-service/internal/app/repository/bookmark"
	cacheRepo "github.com/HemlockPham7/worker-service/internal/app/repository/cache"
	queueRepo "github.com/HemlockPham7/worker-service/internal/app/repository/queue"
	bookmarkSvc "github.com/HemlockPham7/worker-service/internal/app/service/bookmark"
	"github.com/HemlockPham7/worker-service/internal/worker"
)

// CreateAPIConfig creates the API configuration from environment variables.
//
// Returns:
//   - A configured API configuration.
//
// Panics:
//   - If the API configuration cannot be loaded.
func CreateAPIConfig() *api.Config {
	cfg, err := api.NewConfig()
	common.HandleError(err)
	return cfg
}

// CreateEngine initializes the worker dependencies and starts the worker engine.
func CreateEngine() {
	cfg := CreateAPIConfig()

	logger.SetLogLevel(cfg.LogLevel)

	redisClient := CreateRedisClient("worker")

	dbClient := CreateDB("worker")

	nrClient := CreateNRClient("worker")

	queueRepository := queueRepo.NewRedisQueue(redisClient, cfg.QueueName)
	cacheRepository := cacheRepo.NewRedisDB(redisClient)
	codeGenerator := utils.NewGenCode()

	bookmarkRepository := bookmarkRepo.NewRepository(dbClient)
	bookmarkService := bookmarkSvc.NewService(bookmarkRepository, cacheRepository, codeGenerator)

	bookmarkHandler := bookmarkHdl.NewHandler(bookmarkService)

	workerEngine := worker.NewEngine(queueRepository, bookmarkHandler, nrClient)
	workerEngine.Start(context.Background())
}
