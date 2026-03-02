package main

import (
	"context"
	"fmt"

	"github.com/thumbrise/demo/golang-demo/internal/bootstrap/container"
)

func main() {
	// fx.New(
	//	modules.Build()...,
	// ).Run()
	ctx := context.Background()
	c := container.InitializeContainer(ctx)
	fmt.Printf("c = %#v\n", c)
	fmt.Println("WOW")
}
