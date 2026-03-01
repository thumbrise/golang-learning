package redis

import (
	"github.com/thumbrise/demo/golang-demo/internal/contracts"
)

type Config struct {
	URL      string `env:"REDIS_URL" env-required:"true"`
	DB       int    `env:"REDIS_DB"  env-required:"true"`
	Password string `default:"0"     env:"REDIS_PASSWORD" env-required:"false"`
}

func NewConfig(loader contracts.EnvLoader) Config {
	cfg := Config{}

	loader.MustLoad(&cfg)

	return cfg
}
