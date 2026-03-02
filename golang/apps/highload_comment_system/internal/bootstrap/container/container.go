package container

import (
	"github.com/spf13/cobra"
	"github.com/thumbrise/demo/golang-demo/internal/app/core"
	"github.com/thumbrise/demo/golang-demo/internal/bootstrap"
	"github.com/thumbrise/demo/golang-demo/internal/contracts"
	"github.com/thumbrise/demo/golang-demo/internal/modules/shared/http"
)

type Container struct {
	Modules      []contracts.Module
	Runner       *bootstrap.Runner
	Bootstrapper *bootstrap.Bootstrapper
	HttpKernel   *http.Kernel
	CmdKernel    *core.Kernel
	Commands     []*cobra.Command
}

func NewContainer(bootstrapper *bootstrap.Bootstrapper, cmdKernel *core.Kernel, commands []*cobra.Command, httpKernel *http.Kernel, modules []contracts.Module, runner *bootstrap.Runner) *Container {
	return &Container{Bootstrapper: bootstrapper, CmdKernel: cmdKernel, Commands: commands, HttpKernel: httpKernel, Modules: modules, Runner: runner}
}
