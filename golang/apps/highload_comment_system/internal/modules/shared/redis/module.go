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
)

type Module struct {
	redisClient *redis.Client
}

func NewModule(redisClient *redis.Client) *Module {
	return &Module{redisClient: redisClient}
}

func (m *Module) Name() string {
	return "redis"
}

func (m *Module) BeforeStart(ctx context.Context) error {
	return nil
}

func (m *Module) OnStart(ctx context.Context) error {
	return nil
}

func (m *Module) Shutdown(ctx context.Context) error {
	return m.redisClient.Close()
}
