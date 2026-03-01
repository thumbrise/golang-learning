package main

import (
	"bytes"
	"context"
	"fmt"
	"log"

	"github.com/thumbrise/demo/golang-demo/internal/bootstrap/container"
)

func main() {
	ctx := context.Background()

	var (
		c   *container.Container
		err error
	)
	if c, err = container.InitializeContainer(ctx); err != nil {
		log.Fatal(err)
	}

	if err := c.Boot(ctx); err != nil {
		log.Fatal(err)
	}

	buf := bytes.NewBuffer(make([]byte, 0))

	if err := c.CmdKernel.Execute(ctx, buf); err != nil {
		log.Fatal(err)
	}

	if err := c.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	fmt.Print(buf.String())
}
