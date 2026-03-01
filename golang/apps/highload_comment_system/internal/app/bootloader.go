package app

import (
	"context"

	"go.uber.org/fx"
)

type Bootloader struct {
}

func NewBootloader() *Bootloader {
	return &Bootloader{}
}

func (b Bootloader) Name() string {
	return "app"
}

func (b Bootloader) Bind() []fx.Option {
	return []fx.Option{
		fx.Provide(NewBootloader),
		fx.Provide(NewConfig),
	}
}

func (b Bootloader) BeforeStart() error {
	return nil
}

func (b Bootloader) OnStart(ctx context.Context) error {
	return nil
}

func (b Bootloader) Shutdown(ctx context.Context) error {
	return nil
}
