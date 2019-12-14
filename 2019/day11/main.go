package main

import (
	"strconv"

	"github.com/kindermoumoute/adventofcode/pkg/execute"

	"github.com/kindermoumoute/adventofcode/pkg/twod"

	"github.com/kindermoumoute/adventofcode/pkg/font"

	"github.com/kindermoumoute/adventofcode/pkg/intcode"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	part1 := len(drawMap(input, 0))
	drawing := drawMap(input, 1).Filter(1)
	word, err := font.FindWordInMap(drawing)
	if err != nil {
		panic(err)
	}

	return strconv.Itoa(part1), word.String()
}

const (
	Left  = 0
	Right = 1
)

func drawMap(program string, firstInput int) twod.Map {
	c := intcode.New(program, 0, firstInput)
	c.Output.C = make(chan int)
	c.Done = make(chan bool)
	c.RunBackground()
	robot := twod.NewPoint(0, twod.UP)
	drawing := make(twod.Map)
	for {
		select {
		case drawing[robot.Pos] = <-c.Output.C:
			c.Output.C <- 0 // ack

			direction := <-c.Output.C
			c.Output.C <- 0 // ack

			switch direction {
			case Left:
				robot.TurnLeft()
			case Right:
				robot.TurnRight()
			}
			robot.Move(1)

		case <-c.Input.C:
			color, exist := drawing[robot.Pos]
			if !exist {
				color = 0
			}
			c.Input.C <- color.(int)
		case <-c.Done:
			return drawing
		}
	}
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
