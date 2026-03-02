package mail

import (
	"context"

	"github.com/google/wire"
)

var Bindings = wire.NewSet(
	NewModule,
	NewConfig,
)

type Module struct{}

func NewModule() *Module {
	return &Module{}
}

func (m *Module) Name() string {
	return "mail"
}

func (m *Module) BeforeStart(ctx context.Context) error {
	return nil
}

func (m *Module) OnStart(ctx context.Context) error {
	return nil
}

func (m *Module) Shutdown(ctx context.Context) error {
	return nil
}
