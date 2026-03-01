package config

import (
	"gitlab.com/thumbrise-task-manager/task-manager-backend/internal/contracts"
)

type Redis struct {
	URL      string `env:"REDIS_URL" env-required:"true"`
	DB       int    `env:"REDIS_DB"  env-required:"true"`
	Password string `default:"0"     env:"REDIS_PASSWORD" env-required:"false"`
}

func NewRedis(loader contracts.EnvLoader) Redis {
	cfg := Redis{}

	loader.MustLoad(&cfg)

	return cfg
}
