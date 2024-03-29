package main

import (
	"fmt"
	"image/color"

	"github.com/kindermoumoute/adventofcode/pkg/execute"
	"github.com/kindermoumoute/adventofcode/pkg/twod"
	"golang.org/x/image/colornames"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	part1, part2 := 0, 0

	m := twod.NewMapFromInput(input)
	start := twod.Vector(0)
	for vector, tile := range m {
		if tile == 'S' {
			start = vector
		}
	}

	distFromStart := make(twod.Map)
	p := twod.NewPoint(start, twod.UP) // Only work with my input

	startCount := 0
dance:
	for {
		dist, exist := distFromStart[p.Pos]
		if !exist || dist.(int) > p.Steps {
			distFromStart[p.Pos] = p.Steps
		}

		tile, exist := m[p.Pos]
		if !exist {
			panic(fmt.Errorf("%v", p))
		}

		switch tile {
		case '|':
		case '-':
		case '7':
			if p.Speed == twod.RIGHT {
				p.TurnRight()
			} else {
				p.TurnLeft()
			}
		case 'J':
			if p.Speed == twod.DOWN {
				p.TurnRight()
			} else {
				p.TurnLeft()
			}
		case 'F':
			if p.Speed == twod.UP {
				p.TurnRight()
			} else {
				p.TurnLeft()
			}
		case 'L':
			if p.Speed == twod.LEFT {
				p.TurnRight()
			} else {
				p.TurnLeft()
			}
		case '.':
			panic(fmt.Errorf("%v", p))
		case 'S':
			if startCount == 1 {
				p = twod.NewPoint(start, twod.DOWN) // Only work with my input
			} else if startCount == 2 {
				break dance
			}
			startCount++
		}

		p.Move(1)
	}

	for v := range m {
		dist, exist := distFromStart[v]
		if exist && dist.(int) > part1 {
			part1 = dist.(int)
		}
	}

	outsidepoints := make(twod.Map)
	for pos := range m {
		if _, exist := distFromStart[pos]; exist {
			continue
		}
		// reachRight := true
		reachLeft := true

		for y := pos.Y(); y > 0; y-- {
			newPos := twod.NewVector(pos.X(), y)
			_, exist := distFromStart[newPos]

			// if exist && (m[newPos] == '-' || m[newPos] == 'L' || m[newPos] == 'F') {
			// 	reachRight = !reachRight
			// }
			if exist && (m[newPos] == '-' || m[newPos] == '7' || m[newPos] == 'J' || m[newPos] == 'L' || m[newPos] == 'F') {
				reachLeft = !reachLeft
			}
		}
		if !(reachLeft) {
			outsidepoints[pos] = '.'
			part2++
		}
	}

	// Show nice visualization for fun
	twod.RenderingMap = map[interface{}]color.Color{
		'.': colornames.Blue,
	}
	for i := 0; i <= part1; i++ {
		twod.RenderingMap[i] = colornames.Green
	}

	for vector, i := range outsidepoints {
		distFromStart[vector] = i
	}
	distFromStart.Render()
	// time.Sleep(time.Minute)

	return part1, part2
}

func main() {
	execute.RunWithPixel(run, tests, puzzle, true)
}
