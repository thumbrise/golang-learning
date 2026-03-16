package components

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
	"github.com/thumbrise/demo/golang-demo/internal/app"
)

func NewGinEngine(logger *slog.Logger, slogginConfig sloggin.Config, appConfig app.Config) *gin.Engine {
	engine := gin.New()

	engine.Use(sloggin.NewWithConfig(logger, slogginConfig))
	engine.Use(gin.Recovery())

	if appConfig.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	return engine
}
