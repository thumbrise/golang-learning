package cmd

import (
	"bytes"
	"context"
	"fmt"

	"github.com/thumbrise/demo/golang-demo/cmd/cmds"
	"go.uber.org/fx"
)

type Bootloader struct {
	kernel    *Kernel
	serve     *cmds.Serve
	route     *cmds.Route
	routeList *cmds.RouteList
}

func NewBootloader(kernel *Kernel, route *cmds.Route, routeList *cmds.RouteList, serve *cmds.Serve) *Bootloader {
	return &Bootloader{kernel: kernel, route: route, routeList: routeList, serve: serve}
}

func (b *Bootloader) Name() string {
	return "cmd"
}

func (b *Bootloader) Bind() []fx.Option {
	return []fx.Option{
		fx.Provide(NewBootloader),
		fx.Provide(NewKernel),
		fx.Provide(cmds.NewServe),
		fx.Provide(cmds.NewRoute),
		fx.Provide(cmds.NewRouteList),
	}
}

func (b *Bootloader) BeforeStart() error {
	b.kernel.AddGroup(b.route.Command, b.routeList.Command)
	b.kernel.AddCommand(b.serve.Command)

	return nil
}

func (b *Bootloader) OnStart(ctx context.Context) error {
	buf := bytes.NewBuffer(make([]byte, 0))

	err := b.kernel.Execute(ctx, buf)
	if err != nil {
		return err
	}

	fmt.Print(buf.String())

	return nil
}

func (b *Bootloader) Shutdown(ctx context.Context) error {
	return nil
}
