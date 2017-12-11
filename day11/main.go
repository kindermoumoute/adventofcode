package main

import (
	"fmt"
	"math"
	"strings"
)

type d struct {
	x, y int
}

var directionMap = map[string]d{
	"s":  {1, 0},
	"se": {1, -1},
	"ne": {0, -1},
	"n":  {-1, 0},
	"nw": {-1, 1},
	"sw": {0, 1},
}

func main() {
	x, y := 0, 0
	dMax := 0
	for _, p := range strings.Split(puzzle, ",") {
		x += directionMap[p].x
		y += directionMap[p].y
		d := dist(x, y)
		if d > dMax {
			dMax = d
		}
	}

	fmt.Println(dist(x, y), dMax)
}

func dist(x, y int) int {
	d := 0.0
	if x < 0 && y < 0 || y > 0 && x > 0 {
		d = math.Abs(float64(x)) + math.Abs(float64(y))
	} else {
		d = math.Max(math.Abs(float64(x)), math.Abs(float64(y)))
	}
	return int(d)
}
