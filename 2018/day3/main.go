package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
)

// returns part1 and part2
func run(input string) (string, string) {
	squares := parse(input)
	mapPoints := make(map[pkg.P]int)
	part1 := 0
	for _, square := range squares {
		topLeft := square.topLeftCorner()
		botRight := square.botRightCorner()

		for i := topLeft.X; i < botRight.X; i++ {
			for j := topLeft.Y; j < botRight.Y; j++ {
				mapPoints[pkg.P{X: i, Y: j}]++
				if mapPoints[pkg.P{X: i, Y: j}] == 2 {
					part1++
				}
			}
		}
	}

theGame:
	for n, square := range squares {
		topLeft := square.topLeftCorner()
		botRight := square.botRightCorner()

		for i := topLeft.X; i < botRight.X; i++ {
			for j := topLeft.Y; j < botRight.Y; j++ {
				if mapPoints[pkg.P{X: i, Y: j}] != 1 {
					continue theGame
				}
			}
		}
		return strconv.Itoa(part1), strconv.Itoa(n + 1)
	}
	return strconv.Itoa(part1), ""
}

func (d dimsAndmult) topLeftCorner() pkg.P {
	return pkg.P{
		X: d.marginLeft,
		Y: d.marginTop,
	}
}

func (d dimsAndmult) botRightCorner() pkg.P {
	return pkg.P{
		X: d.marginLeft + d.wide,
		Y: d.marginTop + d.tall,
	}
}

type dimsAndmult struct {
	marginLeft, marginTop int

	wide, tall int
}

func parse(s string) []dimsAndmult {
	lines := strings.Split(s, "\n")
	squares := make([]dimsAndmult, len(lines))
	for i, line := range lines {
		var tmpI int
		n, err := fmt.Sscanf(line, "#%d @ %d,%d: %dx%d",
			&tmpI,
			&squares[i].marginLeft,
			&squares[i].marginTop,
			&squares[i].wide,
			&squares[i].tall,
		)
		pkg.Check(err)
		if n != 5 {
			panic("less than 5 values")
		}
	}
	return squares
}

func main() {
	tests.Run(run, false)
	fmt.Println(run(puzzle))
}
