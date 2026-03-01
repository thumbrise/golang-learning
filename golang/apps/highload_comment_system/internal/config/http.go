package config

import (
	"github.com/thumbrise/demo/golang-demo/internal/contracts"
)

type Http struct {
	Port string `env:"HTTP_PORT" env-required:"true"`
}

func NewHttp(loader contracts.EnvLoader) Http {
	cfg := Http{}

	loader.MustLoad(&cfg)

	return cfg
}
