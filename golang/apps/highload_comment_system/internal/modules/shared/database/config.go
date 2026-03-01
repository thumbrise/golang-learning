package database

import (
	"github.com/thumbrise/demo/golang-demo/internal/contracts"
)

type Config struct {
	Host     string `env:"DB_HOST"     env-required:"true"`
	Port     string `env:"DB_PORT"     env-required:"true"`
	Database string `env:"DB_DATABASE" env-required:"true"`
	Username string `env:"DB_USERNAME" env-required:"true"`
	Password string `env:"DB_PASSWORD" env-required:"true"`
}

func NewConfig(loader contracts.EnvLoader) Config {
	cfg := Config{}

	loader.MustLoad(&cfg)

	return cfg
}
