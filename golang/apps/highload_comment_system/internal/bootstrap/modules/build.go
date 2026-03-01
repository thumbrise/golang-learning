package modules

import (
	"github.com/thumbrise/demo/golang-demo/internal"
	"github.com/thumbrise/demo/golang-demo/internal/contracts"
	"go.uber.org/fx"
)

func Build() []fx.Option {
	loaders := internal.Bootloaders()

	opts := make([]fx.Option, 0, len(loaders))
	for _, loader := range loaders {
		l := loader
		opts = append(opts, buildModule(l))
	}

	return opts
}

// buildModule builds a module for a bootloader
//
//nolint:ireturn // spec case
func buildModule(bootloader contracts.Bootloader) fx.Option {
	return fx.Module(
		bootloader.Name(),
		fx.Options(bootloader.Bind()...),
		fx.Invoke(bootloader.BeforeStart),
		fx.Invoke(func(lc fx.Lifecycle) {
			lc.Append(fx.Hook{
				OnStart: bootloader.OnStart,
				OnStop:  bootloader.Shutdown,
			})
		}),
	)
}
