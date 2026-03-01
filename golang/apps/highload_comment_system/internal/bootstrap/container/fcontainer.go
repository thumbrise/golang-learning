package container

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/thumbrise/demo/golang-demo/cmd"
	"github.com/thumbrise/demo/golang-demo/internal/bootstrap"
	"github.com/thumbrise/demo/golang-demo/internal/contracts"
	"github.com/thumbrise/demo/golang-demo/internal/infrastructure/kernels/http"
	"go.uber.org/fx"
)

type FContainer struct {
	bootloaders            []contracts.Bootloader
	bootedBootloadersTypes map[string]bool
	logger                 *slog.Logger
	Runner                 *bootstrap.Runner
	HttpKernel             *http.Kernel
	CmdKernel              *cmd.Kernel
}

func NewFContainer(
	bootloaders []contracts.Bootloader,
	logger *slog.Logger,
	runner *bootstrap.Runner,
	httpKernel *http.Kernel,
	cmdKernel *cmd.Kernel,
	lc fx.Lifecycle,
) *FContainer {
	c := &FContainer{
		bootloaders: bootloaders,
		logger:      logger,
		Runner:      runner,
		HttpKernel:  httpKernel,
		CmdKernel:   cmdKernel,
	}
	c.bootedBootloadersTypes = make(map[string]bool)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return c.Boot(ctx)
		},
		OnStop: func(ctx context.Context) error {
			return c.Shutdown(ctx)
		},
	})

	return c
}

// Boot boots all underlying loaders
//
// Does not stops booting if error occurred. Trying boot all.
//
// Don't forget to call FContainer.Shutdown in the end
func (c *FContainer) Boot(ctx context.Context) error {
	err := error(nil)

	for _, bootloader := range c.bootloaders {
		c.logger.Info(fmt.Sprintf("booting: %T", bootloader))

		errLoader := bootloader.Boot(ctx)
		if errLoader != nil {
			err = errors.Join(err, errLoader)
		} else {
			bootloaderType := fmt.Sprintf("%T", bootloader)
			c.bootedBootloadersTypes[bootloaderType] = true
		}
	}

	if err != nil {
		err = fmt.Errorf("%w: %w", ErrBootloaderBoot, err)
		c.logger.Error("container boot error", "err", err)

		return err
	}

	return nil
}

func (c *FContainer) Shutdown(ctx context.Context) error {
	err := error(nil)

	for i := len(c.bootloaders) - 1; i >= 0; i-- {
		bootloaderType := fmt.Sprintf("%T", c.bootloaders[i])
		if !c.bootedBootloadersTypes[bootloaderType] {
			continue
		}

		bootloader := c.bootloaders[i]
		c.logger.Info(fmt.Sprintf("shutdown-ing: %T", bootloader))

		errLoader := bootloader.Shutdown(ctx)
		if errLoader != nil {
			err = errors.Join(err, errLoader)
		}
	}

	if err != nil {
		err = fmt.Errorf("%w: %w", ErrBootloaderShutdown, err)
		c.logger.Error("container shutdown error", "err", err)

		return err
	}

	return nil
}
