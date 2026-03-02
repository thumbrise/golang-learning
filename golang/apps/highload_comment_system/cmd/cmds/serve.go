package cmds

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/thumbrise/demo/golang-demo/internal/bootstrap"
	"github.com/thumbrise/demo/golang-demo/internal/modules/shared/http"
)

type Serve struct {
	*cobra.Command
}

func NewServe(runner *bootstrap.Runner, httpKernel *http.Kernel) *Serve {
	c := &cobra.Command{
		Use:   "serve",
		Short: "Start http server",
		RunE: func(cmd *cobra.Command, args []string) error {
			processes := []*bootstrap.Process{
				{
					Name: "http kernel",
					Start: func(ctx context.Context) error {
						return httpKernel.Start(ctx)
					},
					Shutdown: func(ctx context.Context) error {
						return httpKernel.Shutdown(ctx)
					},
				},
			}

			return runner.Run(
				cmd.Context(),
				processes,
			)
		},
	}

	return &Serve{c}
}
