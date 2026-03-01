package swagger

import (
	"context"

	"gitlab.com/thumbrise-task-manager/task-manager-backend/internal/modules/swagger/endpoints/http"
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
