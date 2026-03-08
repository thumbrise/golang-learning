package comments

import (
	"context"

	"github.com/google/wire"
	"github.com/thumbrise/demo/golang-demo/internal/modules/comments/application/usecases"
	"github.com/thumbrise/demo/golang-demo/internal/modules/comments/endpoints/cmd"
	"github.com/thumbrise/demo/golang-demo/internal/modules/comments/endpoints/http"
)

var Bindings = wire.NewSet(
	NewModule,
	cmd.NewComments,
	cmd.NewCommentsProduce,
	usecases.NewCommentsCommandPublish,
	http.NewRouter,
)

type Module struct {
	cmd        *cmd.Comments
	cmdProduce *cmd.CommentsProduce
	router     *http.Router
}

func NewModule(cmd *cmd.Comments, cmdProduce *cmd.CommentsProduce, router *http.Router) *Module {
	return &Module{
		router:     router,
		cmd:        cmd,
		cmdProduce: cmdProduce,
	}
}

func (m *Module) Name() string {
	return "comments"
}

func (m *Module) BeforeStart(ctx context.Context) error {
	m.router.Register()

	return nil
}

func (m *Module) OnStart(ctx context.Context) error {
	return nil
}

func (m *Module) Shutdown(ctx context.Context) error {
	return nil
}
