package homepage

import (
	"github.com/thumbrise/demo/golang-demo/internal/modules/homepage/endpoints/http"
	"github.com/thumbrise/demo/golang-demo/internal/modules/homepage/infrastucture/generator"
	"go.uber.org/fx"
)

var Module = fx.Module("homepage",
	fx.Provide(
		http.NewHomePageRouter,
		generator.NewGenerator,
	),
	fx.Invoke(func(router *http.HomePageRouter) {
		router.Register()
	}),
)
