package comments

import (
	"context"

	"github.com/google/wire"
	"github.com/thumbrise/demo/golang-demo/internal/modules/comments/application/cmd"
	"github.com/thumbrise/demo/golang-demo/internal/modules/comments/application/usecases"
)

var Bindings = wire.NewSet(
	NewModule,
	cmd.NewComments,
	cmd.NewCommentsProduce,
	usecases.NewCommentsCommandProduce,
)

type Module struct{}

func NewModule(*cmd.Comments, *cmd.CommentsProduce) *Module {
	return &Module{}
}

func (m *Module) Name() string {
	return "comments"
}

func (m *Module) BeforeStart(ctx context.Context) error {
	return nil
}

func (m *Module) OnStart(ctx context.Context) error {
	return nil
}

func (m *Module) Shutdown(ctx context.Context) error {
	return nil
}
