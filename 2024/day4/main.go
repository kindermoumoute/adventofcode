package main

import (
	"github.com/kindermoumoute/adventofcode/pkg/execute"
	"github.com/kindermoumoute/adventofcode/pkg/twod"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	part1, part2 := 0, 0

	m := twod.NewMapFromInput(input)
	for vector := range m.Filter('X') {
		for _, direction := range twod.AllDirections {
			v, exist := m[vector+direction]
			if exist && v == 'M' {
				v, exist = m[vector+2*direction]
				if exist && v == 'A' {
					v, exist = m[vector+3*direction]
					if exist && v == 'S' {
						part1++
					}
				}
			}
		}
	}

	for vector := range m.Filter('A') {
		v, exist := m[vector+twod.UPRIGHT]
		v2, exist2 := m[vector+twod.DOWNLEFT]
		if exist2 && exist {
			letters := string(v.(int32)) + string(v2.(int32))
			if letters == "MS" || letters == "SM" {
				v, exist = m[vector+twod.UPLEFT]
				v2, exist2 = m[vector+twod.DOWNRIGHT]
				if exist2 && exist {
					letters = string(v.(int32)) + string(v2.(int32))
					if letters == "MS" || letters == "SM" {
						part2++
					}
				}
			}
		}
	}

	return part1, part2
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
