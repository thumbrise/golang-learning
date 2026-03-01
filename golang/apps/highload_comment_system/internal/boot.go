package internal

import (
	"github.com/thumbrise/demo/golang-demo/cmd"
	"github.com/thumbrise/demo/golang-demo/internal/app"
	"github.com/thumbrise/demo/golang-demo/internal/modules/auth"
	"github.com/thumbrise/demo/golang-demo/internal/modules/homepage"
	"github.com/thumbrise/demo/golang-demo/internal/modules/observability"
	"github.com/thumbrise/demo/golang-demo/internal/modules/shared/database"
	"github.com/thumbrise/demo/golang-demo/internal/modules/shared/errorsmap"
	"github.com/thumbrise/demo/golang-demo/internal/modules/shared/http"
	"github.com/thumbrise/demo/golang-demo/internal/modules/shared/mail"
	"github.com/thumbrise/demo/golang-demo/internal/modules/shared/redis"
	"github.com/thumbrise/demo/golang-demo/internal/modules/swagger"
	"go.uber.org/fx"
)

func Bootloaders() []fx.Option {
	return []fx.Option{
		app.Module,
		cmd.Module,
		http.Module,
		database.Module,
		mail.Module,
		redis.Module,
		errorsmap.Module,
		swagger.Module,
		observability.Module,
		auth.Module,
		homepage.Module,
	}
}
