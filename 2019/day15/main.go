package main

import (
	"image/color"

	"github.com/kindermoumoute/adventofcode/pkg/execute"
	"github.com/kindermoumoute/adventofcode/pkg/intcode"
	"github.com/kindermoumoute/adventofcode/pkg/twod"
	"golang.org/x/image/colornames"
)

const (
	wall = iota
	empty
	oxygen
	deadend
	robot
	lastPoint
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {

	// initialize intcode
	c := intcode.New(input, 0)
	c.Output.C = make(chan int)
	c.Done = make(chan bool)
	c.RunBackground()

	// action convert a speed vector into its corresponding intcode input
	action := map[twod.Vector]int{0: 0, twod.UP: 1, twod.DOWN: 2, twod.LEFT: 3, twod.RIGHT: 4}

	// graphic rendering config
	twod.SetFPS(0)
	twod.RenderingMap = map[interface{}]color.Color{
		wall:      colornames.Black,
		empty:     colornames.Yellow,
		oxygen:    colornames.Blue,
		deadend:   colornames.Brown,
		robot:     colornames.Green,
		lastPoint: colornames.Aliceblue,
	}

	// create initial map
	maze := make(twod.Map)
	droid := twod.NewPoint(0, twod.UP)

	// draw droid's initial position
	maze[droid.Pos] = empty
	maze.Render()

	part1 := 0
	oxygenPos := twod.Vector(0)
dance:
	for {
		select {
		case output := <-c.Output.C:
			c.Output.C <- 0 // ack

			// draw updated maze
			maze[droid.Pos] = robot
			maze.Render()
			maze[droid.Pos] = output

			switch output {
			case 0:
				droid.Move(-1) // droid hits a wall, and move back
			case 2:
				oxygenPos = droid.Pos
				_, dist, _ := droid.ShortestPathToPos(0, maze.Filter(1, 2))
				part1 = int(dist)
			}

		case <-c.Input.C: // input request
			c.Input.C <- action[strat(maze, droid)]
		case <-c.Done:
			break dance
		}
	}

	// at this point the full maze has been explored, every empty tile is now a deadend

	// set the droid at the oxygen position
	droid.Pos = oxygenPos
	maze[droid.Pos] = robot

	// find the longest distance in the maze
	_, part2 := droid.LongestPos(maze.Filter(deadend))

	return part1, part2
}

// this strat is way too complicated because it is trying relative direction, but constant direction should work as well
func strat(m twod.Map, p *twod.P) twod.Vector {
	goodDirs := []twod.Vector(nil)
	for i := 0; i < 8; i++ {
		p.Move(1)
		v, exist := m[p.Pos]
		if !exist || i >= 4 && v != wall && v != deadend {
			goodDirs = append(goodDirs, p.Speed)
		}

		p.Move(-1)

		// straight, left, right, back
		switch i % 4 {
		case 0:
			p.TurnLeft()
		case 1, 3:
			p.Speed *= -1
		case 2:
			p.TurnRight()
		}
	}
	if len(goodDirs) == 0 {
		return 0
	}
	if len(goodDirs) == 1 {
		m[p.Pos] = deadend
	}
	p.Speed = goodDirs[0]
	p.Move(1)
	return p.Speed
}

func main() {
	execute.RunWithPixel(run, tests, puzzle, true)
}
