package main

import (
	"context"
	"log"
	"log/slog"

	"github.com/thumbrise/demo/golang-demo/internal/bootstrap/container/wire"
)

func main() {
	ctx := context.Background()

	c, err := wire.InitializeContainer(ctx)
	if err != nil {
		log.Fatalf("error initialize container: %s", err.Error())
	}

	err = c.Bootstrapper.Bootstrap(ctx, c.Modules)
	if err != nil {
		log.Fatalf("error bootstrap modules: %s", err.Error())
	}

	err = c.CmdKernel.Execute(ctx)
	if err != nil {
		slog.Error("main CmdKernel.Execute " + err.Error())
	}

	err = c.Bootstrapper.Shutdown(ctx, c.Modules)
	if err != nil {
		slog.Error("main CmdKernel.Execute " + err.Error())
	}
}
