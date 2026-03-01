package user

import (
	"context"
	"log/slog"

	"gitlab.com/thumbrise-task-manager/task-manager-backend/internal/modules/user/endpoints/http/routers"
)

type Bootloader struct {
	logger *slog.Logger
	router *routers.UsersRouter
}

func (b *Bootloader) Shutdown(context.Context) error {
	return nil
}

func NewBootloader(logger *slog.Logger, router *routers.UsersRouter) *Bootloader {
	return &Bootloader{logger: logger, router: router}
}

func (b *Bootloader) Boot(ctx context.Context) error {
	b.router.Register()

	return nil
}
