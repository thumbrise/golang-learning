package comments

import (
	"context"

	"github.com/google/wire"
	"github.com/thumbrise/demo/golang-demo/internal/modules/comments/application/cmd"
)

var Bindings = wire.NewSet(
	NewModule,
	cmd.NewComments,
)

type Module struct{}

func NewModule(*cmd.Comments) *Module {
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
