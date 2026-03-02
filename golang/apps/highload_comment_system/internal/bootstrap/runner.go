package bootstrap

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"

	"github.com/thumbrise/demo/golang-demo/internal/contracts"
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

func (h *Runner) Run(ctx context.Context, modules []contracts.Module) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, os.Kill)
	defer cancel()

	for _, mm := range modules {
		m := mm
		h.logHook("Got", m.Name(), nil)
	}

	for _, mm := range modules {
		m := mm

		err := m.BeforeStart(ctx)
		h.logHook("BeforeStart", m.Name(), err)
	}

	for _, mm := range modules {
		m := mm

		err := m.OnStart(ctx)
		h.logHook("OnStart", m.Name(), err)
	}

	grp, ctx := errgroup.WithContext(ctx)
	for _, mm := range modules {
		m := mm

		h.logHook("LongRun", m.Name(), nil)

		grp.Go(func() error {
			err := m.LongRun(ctx)
			if err != nil {
				h.logHook("LongRun", m.Name(), err)
			}

			return err
		})
	}

	grp.Go(func() error {
		h.logger.Info("waiting for signal")
		<-ctx.Done()
		h.logger.Info("received signal to exit")

		grpShutdown, ctxShutdown := errgroup.WithContext(ctx)
		for _, mm := range modules {
			m := mm

			grpShutdown.Go(func() error {
				err := m.Shutdown(ctxShutdown)
				h.logHook("Shutdown", m.Name(), err)
				return err
			})
		}

		return grpShutdown.Wait()
	})

	return grp.Wait()
}

func (h *Runner) logHook(hook, moduleName string, err error) {
	msg := fmt.Sprintf("module %s: hook %s", moduleName, hook)
	if err != nil {
		msg = fmt.Sprintf("%s ERROR: %s", msg, err)
		h.logger.Error(msg)
	} else {
		h.logger.Info(msg)
	}
}
