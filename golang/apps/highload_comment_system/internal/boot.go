package internal

import (
	"github.com/thumbrise/demo/golang-demo/cmd"
	"github.com/thumbrise/demo/golang-demo/internal/app"
	"github.com/thumbrise/demo/golang-demo/internal/contracts"
	"github.com/thumbrise/demo/golang-demo/internal/modules/auth"
	"github.com/thumbrise/demo/golang-demo/internal/modules/homepage"
	"github.com/thumbrise/demo/golang-demo/internal/modules/observability"
	"github.com/thumbrise/demo/golang-demo/internal/modules/shared/database"
	"github.com/thumbrise/demo/golang-demo/internal/modules/shared/errorsmap"
	"github.com/thumbrise/demo/golang-demo/internal/modules/shared/http"
	"github.com/thumbrise/demo/golang-demo/internal/modules/shared/mail"
	"github.com/thumbrise/demo/golang-demo/internal/modules/swagger"
)

func Bootloaders() []contracts.Bootloader {
	return []contracts.Bootloader{
		&app.Bootloader{},
		&cmd.Bootloader{},
		&http.Bootloader{},
		&database.Bootloader{},
		&mail.Bootloader{},
		&errorsmap.Bootloader{},
		&swagger.Bootloader{},
		&observability.Bootloader{},
		&auth.Bootloader{},
		&homepage.Bootloader{},
	}
}
