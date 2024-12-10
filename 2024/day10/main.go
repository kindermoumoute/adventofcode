package main

import (
	"github.com/kindermoumoute/adventofcode/pkg/execute"
	"github.com/kindermoumoute/adventofcode/pkg/twod"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	part1, part2 := 0, 0

	m := twod.NewMapFromInput(input)
	for vector := range m.Filter('0') {
		trail := make(twod.Map)
		part1 += FindTrail(m, trail, vector, true)
		part2 += FindTrail(m, trail, vector, false)
	}

	return part1, part2
}

func FindTrail(m, trail twod.Map, pos twod.Vector, ignoreSeen bool) int {
	if m[pos] == '9' {
		return 1
	}

	foundTrails := 0
	for _, direction := range twod.FourDirections {
		nextSlope, exist := m[pos+direction]
		if !exist {
			continue
		}
		if _, exist = trail[direction+pos]; exist && ignoreSeen {
			continue
		}

		if nextSlope.(int32)-m[pos].(int32) == 1 {
			trail[pos+direction] = nextSlope
		} else {
			continue
		}
		foundTrails += FindTrail(m, trail, direction+pos, ignoreSeen)
	}

	return foundTrails
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
