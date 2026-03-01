package modules

import (
	"github.com/thumbrise/demo/golang-demo/internal"
	"go.uber.org/fx"
)

func Build() []fx.Option {
	return internal.Bootloaders()
}
