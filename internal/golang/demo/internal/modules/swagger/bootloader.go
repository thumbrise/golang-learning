package swagger

import (
	"context"

	"github.com/thumbrise/demo/golang-demo/internal/modules/swagger/endpoints/http"
)

type Bootloader struct {
	swaggerRouter *http.SwaggerRouter
}

func (b *Bootloader) Shutdown(context.Context) error {
	return nil
}

func NewBootloader(
	swaggerRouter *http.SwaggerRouter,
) *Bootloader {
	return &Bootloader{
		swaggerRouter: swaggerRouter,
	}
}

func (b *Bootloader) Boot(ctx context.Context) error {
	b.swaggerRouter.Register()

	return nil
}
