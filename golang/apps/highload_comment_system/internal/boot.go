package internal

import (
	"github.com/thumbrise/demo/golang-demo/cmd"
	"github.com/thumbrise/demo/golang-demo/internal/contracts"
	"github.com/thumbrise/demo/golang-demo/internal/modules/auth"
	"github.com/thumbrise/demo/golang-demo/internal/modules/observability"
	"github.com/thumbrise/demo/golang-demo/internal/modules/shared/errorsmap"
	"github.com/thumbrise/demo/golang-demo/internal/modules/swagger"
)

func Bootloaders(
	// main
	cmdLoader *cmd.Bootloader,
	// shared
	errorMapLoader *errorsmap.Bootloader,
	observabilityLoader *observability.Bootloader,
	swaggerLoader *swagger.Bootloader,
	// modules
	authLoader *auth.Bootloader,
) []contracts.Bootloader {
	return []contracts.Bootloader{
		// main
		cmdLoader,

		// shared
		errorMapLoader,
		swaggerLoader,

		// modules
		observabilityLoader,
		authLoader,
	}
}
