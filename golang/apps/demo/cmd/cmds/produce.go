package cmds

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/thumbrise/demo/golang-demo/internal/bootstrap"
	"github.com/thumbrise/demo/golang-demo/internal/contracts"
	"github.com/thumbrise/demo/golang-demo/internal/modules/plugins/http/components"
)

type Produce struct {
	*cobra.Command
}

func NewProduce(r contracts.CMDAdder, runner *bootstrap.Runner, httpKernel *components.Kernel) *Produce {
	c := &cobra.Command{
		Use:   "produce",
		Short: "Produce messages",
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
	r.Add(c)

	return &Produce{c}
}
