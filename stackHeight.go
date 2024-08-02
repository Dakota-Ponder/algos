package main

import (
	"math"
)

func StackHeight2d(layers int) float64 {
	
	if layers <= 0{
		return 0
	}
	
	sqrt3over2 := math.Sqrt(3) / 2

	return 1 + float64(layers - 1) * sqrt3over2
}