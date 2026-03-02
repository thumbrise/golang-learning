package database

import (
	"context"

	"github.com/google/wire"
)

var Bindings = wire.NewSet(
	NewModule,
	NewDB,
	NewConfig,
)

type Module struct {
	db *DB
}

func NewModule() *Module {
	return &Module{}
}

func (m *Module) Name() string {
	return "database"
}

func (m *Module) BeforeStart(ctx context.Context) error {
	return nil
}

func (m *Module) OnStart(ctx context.Context) error {
	return m.db.Connect(ctx)
}

func (m *Module) Shutdown(ctx context.Context) error {
	m.db.Pool().Close()

	return nil
}

func (m *Module) LongRun(ctx context.Context) error {
	return nil
}
