package internal

import (
	"gitlab.com/thumbrise-task-manager/task-manager-backend/cmd"
	"gitlab.com/thumbrise-task-manager/task-manager-backend/internal/contracts"
	"gitlab.com/thumbrise-task-manager/task-manager-backend/internal/modules/auth"
	"gitlab.com/thumbrise-task-manager/task-manager-backend/internal/modules/observability"
	"gitlab.com/thumbrise-task-manager/task-manager-backend/internal/modules/shared/errorsmap"
	"gitlab.com/thumbrise-task-manager/task-manager-backend/internal/modules/swagger"
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
