package main

import (
	"strconv"

	"github.com/kindermoumoute/adventofcode/pkg/intcode"

	"github.com/kindermoumoute/adventofcode/pkg"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	intList := pkg.ParseIntMap(input, ",")
	part1 := len(drawMap(intList, 0))
	drawing := drawMap(intList, 1)
	screen := []int(nil)
	for j := 0; j > -8; j-- {
		for i := 1; i < 41; i++ {
			point, _ := drawing[pkg.P{X: i, Y: j}]
			screen = append(screen, point)
		}
	}
	word := pkg.NewWordFromScreen(screen, 40, 1)

	return strconv.Itoa(part1), word.String()
}

const (
	ReadColor     = 0
	ReadDirection = 1

	Left  = 0
	Right = 1
)

func drawMap(program map[int]int, input int) map[pkg.P]int {
	c := intcode.New(program, 0, input)
	c.IgnoreNonAddressedMemory = true
	c.Output.C = make(chan int)
	c.RunBackground()

	robot := pkg.NewPoint()
	robot.CurrentDirection = pkg.UP
	drawing := make(map[pkg.P]int)
	outputType := 0
	for output := range c.Output.C {
		switch outputType {
		case ReadColor:
			drawing[pkg.P{X: robot.X, Y: robot.Y}] = output
		case ReadDirection:
			switch output {
			case Left:
				robot.TurnLeft()
			case Right:
				robot.TurnRight()
			}
			robot.Move(1)

			// send current color
			color, _ := drawing[pkg.P{X: robot.X, Y: robot.Y}]
			select {
			case c.Input.C <- color:
			case <-c.Done:
			}
		}
		outputType ^= 1
	}
	return drawing
}

func main() {
	pkg.Execute(run, nil, puzzle, true)
}
