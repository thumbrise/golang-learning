package container

import "errors"

var (
	ErrBootloaderBoot     = errors.New("bootloader failed")
	ErrBootloaderShutdown = errors.New("bootloader shutdown")
)
