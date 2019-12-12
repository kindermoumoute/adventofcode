package main

import (
	"strconv"

	"github.com/kindermoumoute/adventofcode/pkg/twod"

	"github.com/kindermoumoute/adventofcode/pkg/font"

	"github.com/kindermoumoute/adventofcode/pkg/intcode"

	"github.com/kindermoumoute/adventofcode/pkg"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	intList := pkg.ParseIntMap(input, ",")
	part1 := len(drawMap(intList, 0))
	drawing := drawMap(intList, 1)
	word, err := font.FindWordInMap(drawing, 1)
	if err != nil {
		panic(err)
	}

	return strconv.Itoa(part1), word.String()
}

const (
	ReadColor     = 0
	ReadDirection = 1

	Left  = 0
	Right = 1
)

func drawMap(program map[int]int, input int) twod.Map {
	c := intcode.New(program, 0, input)
	c.IgnoreNonAddressedMemory = true
	c.Output.C = make(chan int)
	c.RunBackground()

	robot := pkg.NewPoint()
	robot.CurrentDirection = pkg.UP
	drawing := make(twod.Map)
	outputType := 0
	for output := range c.Output.C {
		switch outputType {
		case ReadColor:
			drawing[twod.NewVector(robot.X, robot.Y)] = output
		case ReadDirection:
			switch output {
			case Left:
				robot.TurnLeft()
			case Right:
				robot.TurnRight()
			}
			robot.Move(1)

			// send current color
			color, exist := drawing[twod.NewVector(robot.X, robot.Y)]
			if !exist {
				color = 0
			}
			select {
			case c.Input.C <- color.(int):
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
