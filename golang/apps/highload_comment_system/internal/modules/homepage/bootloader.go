package homepage

import (
	"context"

	"github.com/thumbrise/demo/golang-demo/internal/modules/homepage/endpoints/http"
)

type Bootloader struct {
	router *http.HomePageRouter
}

func NewBootloader(router *http.HomePageRouter) *Bootloader {
	return &Bootloader{router: router}
}

func (b *Bootloader) Shutdown(ctx context.Context) error {
	return nil
}

func (b *Bootloader) Boot(ctx context.Context) error {
	b.router.Register()

	return nil
}
