package cmd

import (
	"github.com/google/wire"
	"github.com/spf13/cobra"
	"github.com/thumbrise/demo/golang-demo/cmd/cmds"
)

var Bindings = wire.NewSet(
	Commands,
	cmds.NewServe,
	cmds.NewRoute,
	cmds.NewRouteList,
)

func Commands(
	*cmds.Serve,
	*cmds.Route,
	*cmds.RouteList,
) []*cobra.Command {
	return []*cobra.Command{}
}
