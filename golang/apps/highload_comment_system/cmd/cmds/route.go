package cmds

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"github.com/thumbrise/demo/golang-demo/internal/infrastructure/kernels/http"
)

type Route struct {
	*cobra.Command
}

func NewRoute() *Route {
	c := &cobra.Command{
		Use:   "route",
		Short: "Route commands",
	}

	return &Route{c}
}

type RouteList struct {
	*cobra.Command
}

func NewRouteList(httpKernel *http.Kernel) *RouteList {
	c := &cobra.Command{
		Use:   "list",
		Short: "List all app routes",
		RunE: func(cmd *cobra.Command, args []string) error {
			engine := httpKernel.Gin()

			w := table.NewWriter()
			w.AppendHeader(table.Row{"#", "method", "path", "handler"})

			for i, info := range engine.Routes() {
				w.AppendRow(table.Row{i + 1, info.Method, info.Path, info.Handler})
			}

			cmd.Println(w.Render())

			return nil
		},
	}

	return &RouteList{c}
}
