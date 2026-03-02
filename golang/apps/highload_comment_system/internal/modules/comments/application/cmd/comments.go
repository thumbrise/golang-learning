package cmd

import (
	"github.com/spf13/cobra"
	"github.com/thumbrise/demo/golang-demo/internal/contracts"
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

func NewCommentsProduce(r *Comments) *CommentsProduce {
	c := &cobra.Command{
		Use:   "produce",
		Short: "produce",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Println("HELLO FROM COMMENTS")

			return nil
		},
	}
	r.AddCommand(c)

	return &CommentsProduce{c}
}
