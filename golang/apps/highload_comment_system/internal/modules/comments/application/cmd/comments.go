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
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Println("HELLO FROM COMMENTS")

			return nil
		},
	}
	r.Add(c)

	return &Comments{c}
}
