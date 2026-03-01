package cmds

import (
	"github.com/spf13/cobra"
	"github.com/thumbrise/demo/golang-demo/internal/bootstrap"
	"github.com/thumbrise/demo/golang-demo/internal/infrastructure/kernels/http"
)

type Serve struct {
	*cobra.Command
}

func NewServe(runner *bootstrap.Runner, httpKernel *http.Kernel) *Serve {
	c := &cobra.Command{
		Use:   "serve",
		Short: "Start http server",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runner.Run(cmd.Context(), httpKernel)
		},
	}

	return &Serve{c}
}
