package swagger

import (
	"github.com/thumbrise/demo/golang-demo/internal/modules/swagger/endpoints/http"
	"go.uber.org/fx"
)

var Module = fx.Module("swagger",
	fx.Provide(
		http.NewSwaggerRouter,
	),
	fx.Invoke(func(swaggerRouter *http.SwaggerRouter) {
		swaggerRouter.Register()
	}),
)
