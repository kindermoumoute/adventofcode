package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
)

type day2 struct {
	squares []dimsAndmult

	maxX, maxY int
}

type dimsAndmult struct {
	marginLeft, marginTop int

	wide, tall int
}

// returns part1 and part2
func run(input string) (string, string) {
	day2 := parse(input)
	mapPoints := make(map[pkg.P]int)
	for _, square := range day2.squares {
		topLeft := square.topLeftCorner()
		botRight := square.botRightCorner()

		for i := topLeft.X; i < botRight.X; i++ {
			for j := topLeft.Y; j < botRight.Y; j++ {
				mapPoints[pkg.P{X: i, Y: j}]++
			}
		}
	}
	part1 := 0
	for _, point := range mapPoints {

		if point > 1 {
			part1++
		}
	}

	part2 := 0
theGame:
	for i, square := range day2.squares {
		topLeft := square.topLeftCorner()
		botRight := square.botRightCorner()

		for i := topLeft.X; i < botRight.X; i++ {
			for j := topLeft.Y; j < botRight.Y; j++ {

				if mapPoints[pkg.P{X: i, Y: j}] != 1 {
					continue theGame
				}
			}
		}
		part2 = i + 1
	}

	// fmt.Println("PASS")
	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func parse(s string) day2 {
	var err error
	lines := strings.Split(s, "\n")
	day2Str := day2{}
	day2Str.squares = make([]dimsAndmult, len(lines))

	for i, line := range lines {
		tmp := strings.Split(line, " @ ")
		tmp2 := strings.Split(tmp[1], ": ")
		dims := strings.Split(tmp2[0], ",")
		mult := strings.Split(tmp2[1], "x")
		day2Str.squares[i].marginLeft, err = strconv.Atoi(dims[0])
		pkg.Check(err)
		day2Str.squares[i].marginTop, err = strconv.Atoi(dims[1])
		pkg.Check(err)

		day2Str.squares[i].wide, err = strconv.Atoi(mult[0])
		pkg.Check(err)
		day2Str.squares[i].tall, err = strconv.Atoi(mult[1])
		pkg.Check(err)

		if day2Str.squares[i].marginLeft+day2Str.squares[i].wide > day2Str.maxX {
			day2Str.maxX = day2Str.squares[i].marginLeft + day2Str.squares[i].wide
		}

		if day2Str.squares[i].marginTop+day2Str.squares[i].tall > day2Str.maxY {
			day2Str.maxY = day2Str.squares[i].marginTop + day2Str.squares[i].tall
		}

	}
	return day2Str
}

func main() {
	tests.Run(run, false)
	fmt.Println(run(puzzle))
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
