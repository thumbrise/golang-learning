package auth

import (
	"context"
	"log/slog"

	"gitlab.com/thumbrise-task-manager/task-manager-backend/internal/modules/auth/endpoints/http"
)

type Bootloader struct {
	logger *slog.Logger
	router *http.Router
}

func (b *Bootloader) Shutdown(context.Context) error {
	return nil
}

func NewBootloader(
	logger *slog.Logger,
	router *http.Router,
) *Bootloader {
	return &Bootloader{
		logger: logger,
		router: router,
	}
}

func (b *Bootloader) Boot(ctx context.Context) error {
	b.router.Register()

	return nil
}
