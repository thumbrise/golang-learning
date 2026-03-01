package internal

import (
	"github.com/thumbrise/demo/golang-demo/cmd"
	"github.com/thumbrise/demo/golang-demo/internal/contracts"
	"github.com/thumbrise/demo/golang-demo/internal/modules/auth"
	"github.com/thumbrise/demo/golang-demo/internal/modules/homepage"
	"github.com/thumbrise/demo/golang-demo/internal/modules/shared/database"
	"github.com/thumbrise/demo/golang-demo/internal/modules/shared/errorsmap"
	"github.com/thumbrise/demo/golang-demo/internal/modules/swagger"
)

func Bootloaders(
// main
	cmdLoader *cmd.Bootloader,
// shared
	errorMapLoader *errorsmap.Bootloader,
	swaggerLoader *swagger.Bootloader,
// modules
//observabilityLoader *observability.Bootloader,
	authLoader *auth.Bootloader,
	homepageLoader *homepage.Bootloader,
	databaseLoader *database.Bootloader,
) []contracts.Bootloader {
	return []contracts.Bootloader{
		// main
		cmdLoader,

		// shared
		errorMapLoader,
		swaggerLoader,
		databaseLoader,

		// modules
		//observabilityLoader,
		authLoader,
		homepageLoader,
	}
}
