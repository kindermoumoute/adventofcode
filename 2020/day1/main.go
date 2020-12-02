package main

import (
	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	list := pkg.ParseIntList(input, "\n")
	m := make(map[int]int)
	m2 := make(map[int]int)
	for _, v1 := range list {
		for _, v2 := range list {
			m[v1+v2] = v1 * v2
			for _, v3 := range list {
				m2[v1+v2+v3] = v1 * v2 * v3
			}
		}
	}

	part1 := m[2020]
	part2 := m2[2020]

	return part1, part2
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
