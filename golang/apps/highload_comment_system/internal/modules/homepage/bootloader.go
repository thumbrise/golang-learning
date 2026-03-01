package homepage

import (
	"context"
	"fmt"

	"github.com/thumbrise/demo/golang-demo/internal/modules/homepage/endpoints/http"
	"go.uber.org/fx"
)

type Bootloader struct {
	router *http.HomePageRouter
}

func NewBootloader(router *http.HomePageRouter) *Bootloader {
	return &Bootloader{router: router}
}

func (b *Bootloader) Name() string {
	return "homepage"
}

func (b *Bootloader) Bind() []fx.Option {
	fmt.Println("HIIII, im bind homepage")

	return []fx.Option{
		fx.Provide(NewBootloader),
		fx.Provide(http.NewHomePageRouter),
	}
}

func (b *Bootloader) BeforeStart() error {
	b.router.Register()
	fmt.Println("HIIII, im register homepage")

	return nil
}

func (b *Bootloader) OnStart(ctx context.Context) error {
	fmt.Println("HIIII, im boot homepage")

	return nil
}

func (b *Bootloader) Shutdown(ctx context.Context) error {
	fmt.Println("HIIII, im shutdown homepage")

	return nil
}
