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
		log.Fatalf("Error initialize container: %s", err.Error())
	}

	err = c.Bootstrapper.Bootstrap(ctx, c.Modules)
	if err != nil {
		log.Fatalf("Error bootstrap modules: %s", err.Error())
	}

	err = c.CmdKernel.Execute(ctx)
	if err != nil {
		msg := "main CmdKernel.Execute " + err.Error()
		slog.Error(msg)
	}

	err = c.Bootstrapper.Shutdown(ctx, c.Modules)
	if err != nil {
		log.Fatalf("Error shutdown modules: %s", err.Error())
	}
}
