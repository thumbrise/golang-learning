package errorsmap

import (
	"context"

	"github.com/thumbrise/demo/golang-demo/internal/modules/shared/errorsmap/endpoints/http"
	"go.uber.org/fx"
)

type Bootloader struct {
	router *http.ErrorsMapRouter
}

func NewBootloader(
	router *http.ErrorsMapRouter,
) *Bootloader {
	return &Bootloader{
		router: router,
	}
}

func (b *Bootloader) Name() string {
	return "errorsmap"
}

func (b *Bootloader) Bind() []fx.Option {
	return []fx.Option{
		fx.Provide(NewBootloader),
		fx.Provide(http.NewErrorsMapRouter),
	}
}

func (b *Bootloader) BeforeStart() error {
	b.router.Register()
	return nil
}

func (b *Bootloader) OnStart(ctx context.Context) error {
	return nil
}

func (b *Bootloader) Shutdown(context.Context) error {
	return nil
}
