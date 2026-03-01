package redis

import (
	"context"

	"go.uber.org/fx"
)

type Bootloader struct{}

func NewBootloader() *Bootloader {
	return &Bootloader{}
}
func (b *Bootloader) Name() string {
	return "Redis"
}

func (b *Bootloader) Bind() []fx.Option {
	return []fx.Option{
		fx.Provide(NewBootloader),
		fx.Provide(NewConfig),
		fx.Provide(NewClient),
	}
}

func (b *Bootloader) BeforeStart() error {
	return nil
}

func (b *Bootloader) OnStart(ctx context.Context) error {
	return nil
}

func (b *Bootloader) Shutdown(context.Context) error {
	return nil
}
