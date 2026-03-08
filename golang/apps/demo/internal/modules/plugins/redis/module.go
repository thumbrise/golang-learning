package redis

import (
	"context"

	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
)

var Bindings = wire.NewSet(
	NewModule,
	NewConfig,
	NewClient,
	NewOTELRegistrar,
)

type Module struct {
	otelRegistrar *OTELRegistrar
	redisClient   *redis.Client
}

func NewModule(otelRegistrar *OTELRegistrar, redisClient *redis.Client) *Module {
	return &Module{otelRegistrar: otelRegistrar, redisClient: redisClient}
}

func (m *Module) Name() string {
	return "redis"
}

func (m *Module) BeforeStart(ctx context.Context) error {
	return m.otelRegistrar.Register()
}

func (m *Module) OnStart(ctx context.Context) error {
	return nil
}

func (m *Module) Shutdown(ctx context.Context) error {
	return m.redisClient.Close()
}
