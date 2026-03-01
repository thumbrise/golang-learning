package config

import (
	"github.com/thumbrise/demo/golang-demo/internal/contracts"
)

type Mail struct {
	Host string `env:"MAIL_HOST" env-required:"true"`
	Port string `env:"MAIL_PORT" env-required:"true"`
	From string `env:"MAIL_FROM" env-required:"true"`
}

func NewMail(loader contracts.EnvLoader) Mail {
	cfg := Mail{}

	loader.MustLoad(&cfg)

	return cfg
}
