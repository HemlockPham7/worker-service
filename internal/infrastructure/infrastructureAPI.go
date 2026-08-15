package infrastructure

import (
	"github.com/HemlockPham7/common-libs/pkg/common"
	"github.com/HemlockPham7/common-libs/pkg/logger"
	"github.com/HemlockPham7/worker-service/internal/api"
)

func CreateAPIConfig() *api.Config {
	cfg, err := api.NewConfig()
	common.HandleError(err)
	return cfg
}

func CreateEngine() {
	cfg := CreateAPIConfig()

	logger.SetLogLevel(cfg.LogLevel)
}
