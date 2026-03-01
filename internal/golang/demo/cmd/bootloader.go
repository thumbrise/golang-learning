package cmd

import (
	"context"

	"gitlab.com/thumbrise-task-manager/task-manager-backend/cmd/cmds"
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

func (b *Bootloader) Boot(ctx context.Context) error {
	b.kernel.AddGroup(b.route.Command, b.routeList.Command)
	b.kernel.AddCommand(b.serve.Command)

	return nil
}

func (b *Bootloader) Shutdown(ctx context.Context) error {
	return nil
}
