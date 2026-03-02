package wire

import (
	"context"

	"github.com/google/wire"
	"github.com/thumbrise/demo/golang-demo/internal"
	"github.com/thumbrise/demo/golang-demo/internal/bootstrap/container"
)

func InitializeContainer(ctx context.Context) (*container.Container, error) {
	wire.Build(
		container.NewContainer,
		internal.Bindings,
	)

	return &container.Container{}, nil
}
