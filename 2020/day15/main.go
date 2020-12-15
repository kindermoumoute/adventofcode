package main

import (
	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	list := pkg.ParseIntMap(input, ",")
	usageIndices := make(map[int][]int)
	part1 := 0
	recently := 0
	for i := 0; i < 30000000; i++ {
		v, exist := list[i]
		if !exist {
			lastIndex, exist := usageIndices[recently]
			if exist && len(lastIndex) > 1 {
				v = lastIndex[len(lastIndex)-1] - lastIndex[len(lastIndex)-2]
			} else {
				v = 0
			}
		}

		recently = v
		usageIndices[v] = append(usageIndices[v], i)
		if i == 2019 {
			part1 = v
		}
	}
	part2 := recently

	return part1, part2
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
