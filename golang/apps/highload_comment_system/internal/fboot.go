package internal

import (
	"github.com/thumbrise/demo/golang-demo/internal/contracts"
	"github.com/thumbrise/demo/golang-demo/internal/modules/homepage"
)

func FBootloaders() []contracts.FBootloader {
	return []contracts.FBootloader{
		&homepage.FBootloader{},
	}
}
