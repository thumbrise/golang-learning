package cmds

import (
	"github.com/spf13/cobra"
	"gitlab.com/thumbrise-task-manager/task-manager-backend/internal/bootstrap"
	"gitlab.com/thumbrise-task-manager/task-manager-backend/internal/infrastructure/kernels/http"
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
