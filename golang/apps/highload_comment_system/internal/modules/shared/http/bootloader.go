package http

import (
	"context"
	"log/slog"

	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
	"go.uber.org/fx"
)

type Bootloader struct {
	kernel *Kernel
	engine *gin.Engine
}

func NewBootloader(engine *gin.Engine, kernel *Kernel) *Bootloader {
	return &Bootloader{engine: engine, kernel: kernel}
}
func (b *Bootloader) Name() string {
	return "http"
}

func (b *Bootloader) Bind() []fx.Option {
	return []fx.Option{
		fx.Provide(NewBootloader),
		fx.Provide(func(logger *slog.Logger) *gin.Engine {
			engine := gin.New()
			engine.Use(sloggin.New(logger))
			engine.Use(gin.Recovery())
			return engine
		}),
		fx.Provide(NewKernel),
		fx.Provide(NewConfig),
	}
}

func (b *Bootloader) BeforeStart() error {
	return nil
}

func (b *Bootloader) OnStart(ctx context.Context) error {
	return nil
}

func (b *Bootloader) Shutdown(ctx context.Context) error {
	return nil
}
