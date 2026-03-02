package contracts

import (
	"context"
)

type Module interface {
	Name() string
	BeforeStart(ctx context.Context) error
	OnStart(ctx context.Context) error
	Shutdown(ctx context.Context) error
	LongRun(ctx context.Context) error
}
