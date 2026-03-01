package cmd

import (
	"bytes"
	"context"
	"fmt"

	"github.com/thumbrise/demo/golang-demo/cmd/cmds"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"cmd",
	fx.Provide(
		NewKernel,
		cmds.NewServe,
		cmds.NewRoute,
		cmds.NewRouteList,
		fx.Invoke(func(lc fx.Lifecycle, kernel *Kernel, route *cmds.Route, routeList *cmds.RouteList, serve *cmds.Serve) {
			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					buf := bytes.NewBuffer(make([]byte, 0))
					fmt.Println("IM IN CMD BOOTLOADER")
					err := kernel.Execute(ctx, buf)
					if err != nil {
						return err
					}

					fmt.Print(buf.String())

					return nil
				},
			})
			kernel.AddGroup(route.Command, routeList.Command)
			kernel.AddCommand(serve.Command)
		}),
	),
)
