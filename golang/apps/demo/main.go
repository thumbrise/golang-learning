package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/thumbrise/demo/golang-demo/internal/bootstrap/container/wire"
)

var (
	ErrInitializeContainer = errors.New("initilaize container")
	ErrModulesBootstrap    = errors.New("modules bootstrap")
	ErrModulesShutdown     = errors.New("modules shutdown")
	ErrCMDExecute          = errors.New("cmd execute")
)

func main() {
	ctx := context.Background()

	c, err := wire.InitializeContainer(ctx)
	if err != nil {
		slog.ErrorContext(ctx, ErrInitializeContainer.Error(), slog.String("error", err.Error()))
		log.Fatalf("%s: %s", ErrInitializeContainer, err.Error())
	}

	err = c.Bootstrapper.Bootstrap(ctx, c.Modules)
	if err != nil {
		slog.ErrorContext(ctx, ErrModulesBootstrap.Error(), slog.String("error", err.Error()))
		log.Fatalf("%s: %s", ErrModulesBootstrap, err.Error())
	}

	bufout := bytes.NewBuffer(make([]byte, 0))
	buferr := bytes.NewBuffer(make([]byte, 0))

	exitCode := 0
	err = c.CmdKernel.Execute(ctx, bufout, buferr)

	errShutdown := c.Bootstrapper.Shutdown(ctx, c.Modules)
	if errShutdown != nil {
		slog.ErrorContext(ctx, ErrModulesShutdown.Error(), slog.String("error", errShutdown.Error()))

		exitCode = 1
	}

	if err != nil {
		slog.ErrorContext(ctx, ErrCMDExecute.Error(), slog.String("error", err.Error()))
		bufout.WriteString(err.Error())

		exitCode = 1
	}

	if bufout.Len() > 0 {
		_, _ = bufout.WriteTo(os.Stdout)

		fmt.Println()
	}

	if buferr.Len() > 0 {
		_, _ = buferr.WriteTo(os.Stderr)

		fmt.Println()
	}

	os.Exit(exitCode)
}
