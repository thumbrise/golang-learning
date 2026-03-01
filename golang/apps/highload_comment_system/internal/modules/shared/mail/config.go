package mail

import (
	"github.com/thumbrise/demo/golang-demo/internal/contracts"
)

type Config struct {
	Host string `env:"MAIL_HOST" env-required:"true"`
	Port string `env:"MAIL_PORT" env-required:"true"`
	From string `env:"MAIL_FROM" env-required:"true"`
}

func NewConfig(loader contracts.EnvLoader) Config {
	cfg := Config{}

	loader.MustLoad(&cfg)

	return cfg
}
