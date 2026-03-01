package homepage

import (
	"github.com/thumbrise/demo/golang-demo/internal/modules/homepage/endpoints/http"
	"go.uber.org/fx"
)

var Module = fx.Module("homepage",
	fx.Provide(
		http.NewHomePageRouter,
		fx.Invoke(func(router *http.HomePageRouter) {
			router.Register()
		}),
	),
)
