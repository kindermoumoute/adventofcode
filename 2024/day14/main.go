package main

import (
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
	"github.com/kindermoumoute/adventofcode/pkg/twod"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	robots := parse(input)

	wide, tall := 11, 7
	if input == puzzle {
		wide, tall = 101, 103
	}

	part1, part2 := 0, 0
	for i := 0; ; i++ {
		m := make(map[twod.Vector]int)
		for _, robot := range robots {
			m[robot.Pos]++
			robot.Move(1)
			robot.Pos = twod.NewVector((robot.Pos.X()+wide)%wide, (robot.Pos.Y()+tall)%tall)
		}

		if len(m) == len(robots) {
			part2 = i
			break
		}

		if i == 99 {
			quad1, quad2, quad3, quad4 := 0, 0, 0, 0
			for _, robot := range robots {
				switch {
				case robot.Pos.X() < wide/2 && robot.Pos.Y() < tall/2:
					quad1++
				case robot.Pos.X() > wide/2 && robot.Pos.Y() < tall/2:
					quad2++
				case robot.Pos.X() > wide/2 && robot.Pos.Y() > tall/2:
					quad3++
				case robot.Pos.X() < wide/2 && robot.Pos.Y() > tall/2:
					quad4++
				}
			}
			part1 = quad1 * quad2 * quad3 * quad4
		}
	}

	return part1, part2
}

func parse(s string) []*twod.P {
	points := []*twod.P(nil)
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		pX, pY, vX, vY := 0, 0, 0, 0
		pkg.MustScanf(line, "p=%d,%d v=%d,%d", &pX, &pY, &vX, &vY)
		points = append(points, twod.NewPoint(twod.NewVector(pX, pY), twod.NewVector(vX, vY)))
	}
	return points
}

func main() {
	execute.Run(run, nil, puzzle, true)
}
