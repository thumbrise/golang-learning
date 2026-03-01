package errorsmap

import (
	"context"

	"gitlab.com/thumbrise-task-manager/task-manager-backend/internal/modules/shared/errorsmap/endpoints/http"
)

type Bootloader struct {
	router *http.ErrorsMapRouter
}

func (b *Bootloader) Shutdown(context.Context) error {
	return nil
}

func NewBootloader(
	router *http.ErrorsMapRouter,
) *Bootloader {
	return &Bootloader{
		router: router,
	}
}

func (b *Bootloader) Boot(ctx context.Context) error {
	b.router.Register()

	return nil
}
