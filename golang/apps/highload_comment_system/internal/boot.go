package internal

import (
	"github.com/thumbrise/demo/golang-demo/cmd"
	"github.com/thumbrise/demo/golang-demo/internal/contracts"
	"github.com/thumbrise/demo/golang-demo/internal/modules/auth"
	"github.com/thumbrise/demo/golang-demo/internal/modules/homepage"
	"github.com/thumbrise/demo/golang-demo/internal/modules/observability"
	"github.com/thumbrise/demo/golang-demo/internal/modules/shared/database"
	"github.com/thumbrise/demo/golang-demo/internal/modules/shared/errorsmap"
	"github.com/thumbrise/demo/golang-demo/internal/modules/swagger"
)

func Bootloaders() []contracts.Bootloader {
	return []contracts.Bootloader{
		&cmd.Bootloader{},
		// shared
		&errorsmap.Bootloader{},
		&swagger.Bootloader{},
		// modules
		&observability.Bootloader{},
		&auth.Bootloader{},
		&homepage.Bootloader{},
		&database.Bootloader{},
	}
}
