package main

import (
	"github.com/thumbrise/demo/golang-demo/internal/bootstrap/modules"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		modules.Build()...,
	).Run()
}
