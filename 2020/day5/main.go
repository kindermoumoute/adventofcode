package main

import (
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg/execute"
	"github.com/kindermoumoute/adventofcode/pkg/twod"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	part1, part2 := 0, 0

	m := make(twod.Map)

	for _, line := range strings.Split(input, "\n") {
		x := 0
		xd := 64
		y := 0
		yd := 4
		for _, letter := range line {
			switch letter {
			case 'F':
				xd = xd / 2
			case 'B':
				x = x + xd
				xd = xd / 2
			case 'L':
				yd = yd / 2
			case 'R':
				y = y + yd
				yd = yd / 2
			}
		}
		if x*8+y > part1 {
			part1 = x*8 + y
		}
		m[twod.NewVector(y, x)] = 'X'
	}

	for i := 0; i < 128; i++ {
		for j := 0; j < 8; j++ {
			_, exist := m[twod.NewVector(j, i)]
			if !exist && i > 6 && i < 112 {
				part2 = i*8 + j
			}
		}
	}

	return part1, part2
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
