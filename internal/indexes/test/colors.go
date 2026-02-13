package test

import (
	"math/rand"
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
	colorIndex := rand.Intn(len(PossibleColors))
	return PossibleColors[colorIndex]
}

func GetRandomColors() []string {
	colorCount := rand.Intn(len(PossibleColors))
	colors := make([]string, 0, colorCount)
	for range colorCount {
		colors = append(colors, GetRandomColor())
	}
	return colors
}
