package config

import (
	"gitlab.com/thumbrise-task-manager/task-manager-backend/internal/contracts"
)

type Http struct {
	Port string `env:"HTTP_PORT" env-required:"true"`
}

func NewHttp(loader contracts.EnvLoader) Http {
	cfg := Http{}

	loader.MustLoad(&cfg)

	return cfg
}
