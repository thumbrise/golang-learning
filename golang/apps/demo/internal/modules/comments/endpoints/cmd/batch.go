package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/thumbrise/demo/golang-demo/internal/bootstrap"
	"github.com/thumbrise/demo/golang-demo/internal/modules/comments/application/workers"
)

type CommentsBatch struct {
	*cobra.Command
}

func NewCommentsBatch(r *Comments, runner *bootstrap.Runner, batcher *workers.CommentsBatcher) *CommentsBatch {
	c := &cobra.Command{
		Use:   "batch",
		Short: "Process batch comments from buffer. Persist data from bus",
		RunE: func(cmd *cobra.Command, args []string) error {
			p := &bootstrap.Process{
				Name: "comments batcher",
				Start: func(ctx context.Context) error {
					return batcher.Run(ctx)
				},
				Shutdown: func(ctx context.Context) error {
					return nil
				},
			}

			return runner.Run(
				cmd.Context(),
				p,
			)
		},
	}
	r.AddCommand(c)

	return &CommentsBatch{c}
}
