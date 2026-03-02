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
	Modules      []contracts.Module
	Runner       *bootstrap.Runner
	Bootstrapper *bootstrap.Bootstrapper
	HttpKernel   *http.Kernel
	CmdKernel    *cmd.Kernel
}

func NewContainer(bootstrapper *bootstrap.Bootstrapper, cmdKernel *cmd.Kernel, httpKernel *http.Kernel, modules []contracts.Module, runner *bootstrap.Runner) *Container {
	return &Container{Bootstrapper: bootstrapper, CmdKernel: cmdKernel, HttpKernel: httpKernel, Modules: modules, Runner: runner}
}

func InitializeContainer(ctx context.Context) (*Container, error) {
	wire.Build(
		NewContainer,
		internal.Bindings,
	)

	return &Container{}, nil
}
