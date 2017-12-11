package main

import (
	"fmt"
	"math"
	"strings"
)

var directionMap = map[string]struct{ x, y int }{
	"s":  {1, 0},
	"se": {1, -1},
	"ne": {0, -1},
	"n":  {-1, 0},
	"nw": {-1, 1},
	"sw": {0, 1},
}

func main() {
	x, y, d, dMax := 0, 0, 0, 0
	for _, p := range strings.Split(puzzle, ",") {
		x += directionMap[p].x
		y += directionMap[p].y
		d = dist(float64(x), float64(y))
		if d > dMax {
			dMax = d
		}
	}
	fmt.Println(d, dMax)
}

func dist(x, y float64) int {
	return int((math.Abs(x) + math.Abs(y) + math.Abs(x+y)) / 2)
}
