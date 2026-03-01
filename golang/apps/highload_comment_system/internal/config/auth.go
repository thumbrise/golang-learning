package config

import (
	"github.com/thumbrise/demo/golang-demo/internal/contracts"
)

type Auth struct {
	OTPLength      int    `env:"AUTH_OTP_LENGTH"       env-required:"true"`
	OTPTTLMinutes  int    `env:"AUTH_OTP_TTL_MINUTES"  env-required:"true"`
	OTPForcedValue string `env:"AUTH_OTP_FORCED_VALUE"`

	JWTSecret            string `env:"AUTH_JWT_SECRET"              env-required:"true"`
	JWTIssuer            string `env:"AUTH_JWT_ISSUER"              env-required:"true"`
	JWTAccessTTLMinutes  int    `env:"AUTH_JWT_ACCESS_TTL_MINUTES"  env-required:"true"`
	JWTRefreshTTLMinutes int    `env:"AUTH_JWT_REFRESH_TTL_MINUTES" env-required:"true"`
}

func NewAuth(loader contracts.EnvLoader) Auth {
	cfg := Auth{}

	loader.MustLoad(&cfg)

	return cfg
}
