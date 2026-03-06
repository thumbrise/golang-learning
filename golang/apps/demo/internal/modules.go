package internal

import (
	"github.com/google/wire"
	"github.com/thumbrise/demo/golang-demo/cmd"
	"github.com/thumbrise/demo/golang-demo/internal/app/core"
	"github.com/thumbrise/demo/golang-demo/internal/contracts"
	"github.com/thumbrise/demo/golang-demo/internal/modules/auth"
	"github.com/thumbrise/demo/golang-demo/internal/modules/comments"
	"github.com/thumbrise/demo/golang-demo/internal/modules/homepage"
	"github.com/thumbrise/demo/golang-demo/internal/modules/observability"
	"github.com/thumbrise/demo/golang-demo/internal/modules/shared/database"
	"github.com/thumbrise/demo/golang-demo/internal/modules/shared/errorsmap"
	"github.com/thumbrise/demo/golang-demo/internal/modules/shared/http"
	"github.com/thumbrise/demo/golang-demo/internal/modules/shared/mail"
	"github.com/thumbrise/demo/golang-demo/internal/modules/shared/redis"
	"github.com/thumbrise/demo/golang-demo/internal/modules/swagger"
)

var Bindings = wire.NewSet(
	Modules,
	core.Bindings,
	cmd.Bindings,
	http.Bindings,
	auth.Bindings,
	database.Bindings,
	mail.Bindings,
	redis.Bindings,
	errorsmap.Bindings,
	swagger.Bindings,
	observability.Bindings,
	homepage.Bindings,
	comments.Bindings,
)

func Modules(
	httpModule *http.Module,
	observabilityModule *observability.Module,
	databaseModule *database.Module,
	mailModule *mail.Module,
	redisModule *redis.Module,
	errorsmapModule *errorsmap.Module,
	swaggerModule *swagger.Module,
	authModule *auth.Module,
	homepageModule *homepage.Module,
	commentsModule *comments.Module,
) []contracts.Module {
	return []contracts.Module{
		httpModule,
		observabilityModule,
		databaseModule,
		mailModule,
		redisModule,
		errorsmapModule,
		swaggerModule,
		authModule,
		homepageModule,
		commentsModule,
	}
}
