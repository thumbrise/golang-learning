package contracts

import (
	"context"

	"go.uber.org/fx"
)

type Bootloader interface {
	Boot(ctx context.Context) error
	Shutdown(ctx context.Context) error
}
type FBootloader interface {
	Name() string
	Bind() []fx.Option
	BeforeStart() interface{}
	OnStart(ctx context.Context) error
	Shutdown(ctx context.Context) error
}
