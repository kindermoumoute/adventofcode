package main

import (
	"sort"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {

	list := pkg.ParseIntList(input, "\n")
	list = append(list, 0, pkg.Max(list...)+3)
	sort.Slice(list, func(i, j int) bool {
		return list[i] < list[j]
	})

	previous := 0
	diff := make(map[int]int)
	for _, n := range list {
		diff[n-previous]++
		previous = n
	}
	part1 := diff[1] * diff[3]

	arrangements := make([]int, len(list))
	arrangements[0]++
	for i := range list {
		for j := i + 1; j < len(list); j++ {
			if list[j]-list[i] > 3 {
				break
			}
			arrangements[j] += arrangements[i]
		}
	}
	part2 := arrangements[len(arrangements)-1]

	return part1, part2
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
