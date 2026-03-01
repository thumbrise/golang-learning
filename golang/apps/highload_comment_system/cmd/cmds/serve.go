package cmds

import (
	"github.com/spf13/cobra"
	"github.com/thumbrise/demo/golang-demo/internal/bootstrap"
	"github.com/thumbrise/demo/golang-demo/internal/modules/shared/http"
	"go.uber.org/fx"
)

type Serve struct {
	*cobra.Command
}

func NewServe(runner *bootstrap.Runner, httpKernel *http.Kernel, lc fx.Lifecycle) *Serve {
	c := &cobra.Command{
		Use:   "serve",
		Short: "Start http server",
		RunE: func(cmd *cobra.Command, args []string) error {
			//lc.Append(fx.Hook{
			//	OnStart: func(ctx context.Context) error {
			//		return httpKernel.Start(ctx)
			//	},
			//	OnStop: func(ctx context.Context) error {
			//		return httpKernel.Shutdown(ctx)
			//	},
			//})
			return runner.Run(cmd.Context(), httpKernel)
			//return nil
		},
	}

	return &Serve{c}
}
