package database

import (
	"context"

	"go.uber.org/fx"
)

type Bootloader struct {
	db *DB
}

func NewBootloader(db *DB) *Bootloader {
	return &Bootloader{db: db}
}

func (b *Bootloader) Name() string {
	return "database"
}

func (b *Bootloader) Bind() []fx.Option {
	return []fx.Option{
		fx.Provide(NewBootloader),
		fx.Provide(NewDB),
		fx.Provide(NewConfig),
	}
}

func (b *Bootloader) BeforeStart() error {
	return nil
}

func (b *Bootloader) OnStart(ctx context.Context) error {
	return b.db.Connect(ctx)
}

func (b *Bootloader) Shutdown(context.Context) error {
	b.db.Pool().Close()

	return nil
}
