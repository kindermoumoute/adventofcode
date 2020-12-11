package main

import (
	"github.com/kindermoumoute/adventofcode/pkg/execute"
	"github.com/kindermoumoute/adventofcode/pkg/twod"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	m := twod.NewMapFromInput(input)
	part1 := solve(m.Clone(), 4, occupied)
	part2 := solve(m, 5, occupied2)

	return part1, part2
}

func solve(m twod.Map, seatTolerance int, occupiedFunc func(m twod.Map, v twod.Vector) int) int {
	for {
		changed := twod.Map{}
		occupiedCount := 0
		for v, char := range m {
			switch char {
			case 'L':
				if occupiedFunc(m, v) > 0 {
					continue
				}
				changed[v] = '#'
			case '#':
				if occupiedFunc(m, v) >= seatTolerance {
					changed[v] = 'L'
				} else {
					occupiedCount++
				}
			}

		}

		if len(changed) == 0 {
			return occupiedCount
		}
		for k, v := range changed {
			m[k] = v
		}
	}
}

var dirs = []twod.Vector{
	twod.UP,
	twod.DOWN,
	twod.LEFT,
	twod.RIGHT,
	twod.RIGHT + twod.UP,
	twod.RIGHT + twod.DOWN,
	twod.LEFT + twod.DOWN,
	twod.LEFT + twod.UP,
}

func occupied2(m twod.Map, v twod.Vector) int {
	o := 0
	for _, dir := range dirs {
		for i := 1; ; i++ {
			delta := twod.NewVector(i*dir.X(), i*dir.Y())
			char, exist := m[v+delta]
			if !exist || char == 'L' {
				break
			}
			if char == '#' {
				o++
				break
			}
		}
	}
	return o
}

func occupied(m twod.Map, v twod.Vector) int {
	o := 0
	for _, dir := range dirs {
		if m[v+dir] == '#' {
			o++
		}
	}
	return o
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
