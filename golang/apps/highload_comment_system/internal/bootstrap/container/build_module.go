package container

import (
	"github.com/thumbrise/demo/golang-demo/internal/contracts"
	"go.uber.org/fx"
)

func BuildModule(bootloader contracts.FBootloader) fx.Option {
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
