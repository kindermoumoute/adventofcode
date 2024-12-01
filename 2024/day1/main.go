package main

import (
	"sort"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	lists := parse(input)
	// pkg.ParseIntMap()
	sort.Ints(lists[0])
	sort.Ints(lists[1])
	part1, part2 := 0, 0

	intMapLeft := map[int]int{}
	intMapRight := map[int]int{}
	for i := 0; i < len(lists[0]); i++ {
		intMapLeft[lists[0][i]]++
		intMapRight[lists[1][i]]++
		part1 += pkg.Abs(lists[0][i] - lists[1][i])
	}
	for val, countLeft := range intMapLeft {
		countRight, exist := intMapRight[val]
		if exist {
			part2 += countRight * val * countLeft
		}
	}
	return part1, part2
}

func parse(s string) [][]int {
	lists := [][]int{{}, {}}
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		l := pkg.ParseIntList(line, "   ")
		lists[0] = append(lists[0], l[0])
		lists[1] = append(lists[1], l[1])
	}
	return lists
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
