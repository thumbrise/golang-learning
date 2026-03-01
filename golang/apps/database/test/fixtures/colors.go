package fixtures

import (
	"github.com/thumbrise/demo/golang/pkg/random"
)

var PossibleColors = []string{
	"red",
	"blue",
	"green",
	"yellow",
	"black",
	"white",
	"orange",
	"purple",
	"pink",
	"brown",
	"gray",
	"grey",
	"silver",
	"gold",
	"bronze",
	"platinum",
	"transparent",
}

func GetRandomColor() string {
	return PossibleColors[random.Int64(int64(len(PossibleColors)))]
}

func GetRandomColors() []string {
	colorCount := random.Int64(int64(len(PossibleColors) / 2))

	colors := make([]string, 0, colorCount)
	for range colorCount {
		colors = append(colors, GetRandomColor())
	}

	return colors
}
