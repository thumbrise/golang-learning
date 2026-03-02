package cmd

import (
	"context"

	"github.com/google/wire"
	"github.com/thumbrise/demo/golang-demo/cmd/cmds"
)

var Bindings = wire.NewSet(
	NewWireModule,
	NewKernel,
	cmds.NewServe,
	cmds.NewRoute,
	cmds.NewRouteList,
)

type Module struct {
	kernel    *Kernel
	route     *cmds.Route
	routeList *cmds.RouteList
	serve     *cmds.Serve
}

func NewWireModule(kernel *Kernel, route *cmds.Route, routeList *cmds.RouteList, serve *cmds.Serve) *Module {
	return &Module{kernel: kernel, route: route, routeList: routeList, serve: serve}
}
func (m *Module) Name() string {
	return "cmd"
}

func (m *Module) LongRun(ctx context.Context) error {
	return m.kernel.Execute(ctx)
}

func (m *Module) BeforeStart(context.Context) error {
	m.kernel.AddGroup(m.route.Command, m.routeList.Command)
	m.kernel.AddCommand(m.serve.Command)
	return nil
}
func (m *Module) OnStart(ctx context.Context) error {
	return nil
}
func (m *Module) Shutdown(ctx context.Context) error {
	return nil
}
