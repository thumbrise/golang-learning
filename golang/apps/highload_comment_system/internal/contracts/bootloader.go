package contracts

import (
	"context"

	"go.uber.org/fx"
)

type Bootloader interface {
	Name() string
	Bind() []fx.Option
	BeforeStart() error
	OnStart(ctx context.Context) error
	Shutdown(ctx context.Context) error
}
