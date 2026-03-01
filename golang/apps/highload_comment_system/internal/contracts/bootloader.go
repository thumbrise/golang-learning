package contracts

import (
	"go.uber.org/fx"
)

type Bootloader interface {
	Name() string
	Bind() []fx.Option
}
