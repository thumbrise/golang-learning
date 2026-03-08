package cmd

import (
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"github.com/thumbrise/demo/golang-demo/internal/contracts"
	"github.com/thumbrise/demo/golang-demo/internal/modules/comments/application/usecases"
)

type Comments struct {
	*cobra.Command
}

func NewComments(r contracts.CMDAdder) *Comments {
	c := &cobra.Command{
		Use: "comments",
	}
	r.Add(c)

	return &Comments{c}
}

type CommentsProduce struct {
	*cobra.Command
}

func NewCommentsProduce(r *Comments, produce *usecases.CommentsCommandPublish) *CommentsProduce {
	c := &cobra.Command{
		Use: "produce",
		RunE: func(cmd *cobra.Command, args []string) error {
			in := usecases.CommentsCommandPublishInput{
				UserUUID: uuid.New().String(),
				PostUUID: uuid.New().String(),
				Content:  faker.Paragraph(),
			}

			out, err := produce.Handle(cmd.Context(), in)
			if err != nil {
				return err
			}

			cmd.Printf("out:\n\t%#v\n", out)

			return nil
		},
	}
	r.AddCommand(c)

	return &CommentsProduce{c}
}
