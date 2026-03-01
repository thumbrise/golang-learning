package bootstrap

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"

	"golang.org/x/sync/errgroup"
)

type Runner struct {
	logger *slog.Logger
}

func NewRunner(
	logger *slog.Logger,
) *Runner {
	return &Runner{
		logger: logger,
	}
}

type Kernel interface {
	Start(ctx context.Context) error
	Shutdown(ctx context.Context) error
	Name() string
}

func (h *Runner) Run(ctx context.Context, kernel Kernel) error {
	msg := fmt.Sprintf("starting %s kernel", kernel.Name())
	h.logger.Info(msg)

	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, os.Kill)
	defer cancel()

	grp, ctx := errgroup.WithContext(ctx)

	grp.Go(func() error {
		return kernel.Start(ctx)
	})

	grp.Go(func() error {
		h.logger.Info("waiting for signal")
		<-ctx.Done()
		h.logger.Info("received signal to exit")

		return kernel.Shutdown(ctx)
	})

	return grp.Wait()
}
