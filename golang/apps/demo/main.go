package main

import (
	"bytes"
	"context"
	"fmt"
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

	buf := bytes.NewBuffer(make([]byte, 0))

	err = c.CmdKernel.Execute(ctx, buf)
	if err != nil {
		slog.Error("main CmdKernel.Execute " + err.Error())
	}

	err = c.Bootstrapper.Shutdown(ctx, c.Modules)
	if err != nil {
		slog.Error("main CmdKernel.Execute " + err.Error())
	}

	fmt.Print(buf.String())
}
