package main

import (
	"github.com/kindermoumoute/adventofcode/pkg/execute"
	"github.com/kindermoumoute/adventofcode/pkg/intcode"
	"github.com/kindermoumoute/adventofcode/pkg/twod"
	"golang.org/x/image/colornames"
)

const (
	empty = iota
	wall
	brick
	bar
	ball
)

var tiles = map[int]string{
	0: " ",
	1: "|",
	2: "#",
	3: "-",
	4: "o",
}

// returns part1 and part2
func run(input string) (interface{}, interface{}) {

	c := intcode.New(input, 0)
	c.Output.C = make(chan int)
	c.Done = make(chan bool)
	c.Memory[0] = 2
	c.RunBackground()

	gameMap := make(twod.Map)

	part1, part2 := 0, 0
	for {
		select {
		case x := <-c.Output.C:
			c.Output.C <- 0 // ack

			y := <-c.Output.C
			c.Output.C <- 0 // ack

			v := <-c.Output.C
			c.Output.C <- 0 // ack

			if x == -1 {
				if part1 == 0 {
					part1 = len(gameMap.Filter(brick))
				}
				part2 = v
			} else {
				gameMap[twod.NewVector(x, y)] = v
			}

		case <-c.Input.C: // input request
			c.Input.C <- strat(gameMap.Find(bar)[0], gameMap.Find(ball)[0])
			gameMap.Render(twod.RenderingMap{
				wall:  colornames.Brown,
				brick: colornames.Purple,
				bar:   colornames.Black,
				ball:  colornames.Red,
			})
		case <-c.Done:
			return part1, part2
		}
	}
}

func strat(you, ball twod.Vector) int {
	if ball.X() == you.X() {
		return 0
	} else if ball.X() < you.X() {
		return -1
	}
	return 1
}

func main() {
	execute.RunWithPixel(run, tests, puzzle, true)
}
