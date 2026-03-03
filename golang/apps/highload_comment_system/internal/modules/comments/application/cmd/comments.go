package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/thumbrise/demo/golang-demo/internal/bootstrap"
	"github.com/thumbrise/demo/golang-demo/internal/contracts"
	"github.com/thumbrise/demo/golang-demo/internal/modules/comments/application/usecases"
)

type Comments struct {
	*cobra.Command
}

func NewComments(r contracts.CMDAdder) *Comments {
	c := &cobra.Command{
		Use:   "comments",
		Short: "comments",
	}
	r.Add(c)

	return &Comments{c}
}

type CommentsProduce struct {
	*cobra.Command
}

func NewCommentsProduce(r *Comments, runner *bootstrap.Runner, produce *usecases.CommentsCommandProduce) *CommentsProduce {
	c := &cobra.Command{
		Use:   "produce",
		Short: "produce",
		RunE: func(cmd *cobra.Command, args []string) error {
			processes := []*bootstrap.Process{
				{
					Name: "CommentsCommandProduce",
					Start: func(ctx context.Context) error {
						in := usecases.CommentsCommandProduceInput{}
						_, err := produce.Handle(ctx, in)

						return err
					},
					Shutdown: nil,
				},
			}

			return runner.Run(cmd.Context(), processes)
		},
	}
	r.AddCommand(c)

	return &CommentsProduce{c}
}
