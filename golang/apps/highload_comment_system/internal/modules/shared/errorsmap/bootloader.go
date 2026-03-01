package errorsmap

import (
	"github.com/thumbrise/demo/golang-demo/internal/modules/shared/errorsmap/endpoints/http"
	"go.uber.org/fx"
)

var Module = fx.Module("errorsmap",
	fx.Provide(
		http.NewErrorsMapRouter,
		fx.Invoke(func(router *http.ErrorsMapRouter) {
			router.Register()
		}),
	),
)
