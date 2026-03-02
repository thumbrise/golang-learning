package container

import (
	"context"

	"github.com/google/wire"
	"github.com/thumbrise/demo/golang-demo/cmd"
	"github.com/thumbrise/demo/golang-demo/internal"
	"github.com/thumbrise/demo/golang-demo/internal/bootstrap"
	"github.com/thumbrise/demo/golang-demo/internal/contracts"
	"github.com/thumbrise/demo/golang-demo/internal/modules/shared/http"
)

type Container struct {
	Modules    []contracts.Module
	Runner     *bootstrap.Runner
	HttpKernel *http.Kernel
	CmdKernel  *cmd.Kernel
}

func NewContainer(
	modules []contracts.Module,
	runner *bootstrap.Runner,
	httpKernel *http.Kernel,
	cmdKernel *cmd.Kernel,
) *Container {
	c := &Container{
		Modules:    modules,
		Runner:     runner,
		HttpKernel: httpKernel,
		CmdKernel:  cmdKernel,
	}

	return c
}

func InitializeContainer(ctx context.Context) *Container {
	wire.Build(
		NewContainer,
		internal.Bindings,
	)

	return &Container{}
}
