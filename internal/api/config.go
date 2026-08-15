package api

import (
	"github.com/google/uuid"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	LogLevel    string `default:"info" envconfig:"LOG_LEVEL"`
	ServiceName string `default:"worker-service" envconfig:"SERVICE_NAME"`
	InstanceID  string `default:"" envconfig:"INSTANCE_ID"`
	QueueName   string `default:"bookmark-import" envconfig:"QUEUE_NAME"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := envconfig.Process("api", cfg)
	if err != nil {
		return nil, err
	}

	if cfg.InstanceID == "" {
		cfg.InstanceID = uuid.New().String()
	}
	return cfg, err
}
