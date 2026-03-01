package swagger

import (
	"context"

	"github.com/thumbrise/demo/golang-demo/internal/modules/swagger/endpoints/http"
	"go.uber.org/fx"
)

type Bootloader struct {
	swaggerRouter *http.SwaggerRouter
}

func NewBootloader(
	swaggerRouter *http.SwaggerRouter,
) *Bootloader {
	return &Bootloader{
		swaggerRouter: swaggerRouter,
	}
}

func (b *Bootloader) Name() string {
	return "swagger"
}

func (b *Bootloader) Bind() []fx.Option {
	return []fx.Option{
		fx.Provide(NewBootloader),
		fx.Provide(http.NewSwaggerRouter),
	}
}

func (b *Bootloader) BeforeStart() error {
	b.swaggerRouter.Register()

	return nil
}

func (b *Bootloader) OnStart(ctx context.Context) error {
	return nil
}

func (b *Bootloader) Shutdown(context.Context) error {
	return nil
}
