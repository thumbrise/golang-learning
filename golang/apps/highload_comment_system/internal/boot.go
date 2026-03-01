package internal

import (
	"github.com/thumbrise/demo/golang-demo/internal/contracts"
	"github.com/thumbrise/demo/golang-demo/internal/modules/homepage"
)

func Bootloaders() []contracts.Bootloader {
	return []contracts.Bootloader{
		// main
		//cmdLoader *cmd.Bootloader,
		// shared
		//errorMapLoader *errorsmap.Bootloader,
		//swaggerLoader *swagger.Bootloader,
		// modules
		//observabilityLoader *observability.Bootloader,
		//authLoader *auth.Bootloader,
		&homepage.Bootloader{},
		//databaseLoader *database.Bootloader,

	}
}
