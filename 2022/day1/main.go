package main

import (
	"sort"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	elveWeights := parse(input)

	part1 := 0
	allSums := []int(nil)
	for _, weights := range elveWeights {
		sum := pkg.Sum(weights...)
		part1 = pkg.Max(part1, sum)
		allSums = append(allSums, sum)
	}
	sort.Ints(allSums)
	lastIndex := len(allSums) - 1
	part2 := allSums[lastIndex] + allSums[lastIndex-1] + allSums[lastIndex-2]

	return part1, part2
}

func parse(s string) [][]int {
	o := [][]int(nil)
	elves := strings.Split(s, "\n\n")
	for _, line := range elves {
		oe := []int(nil)
		weights := strings.Split(line, "\n")
		for _, w := range weights {
			oe = append(oe, pkg.MustAtoi(w))

		}
		o = append(o, oe)
	}
	return o
}

func main() {
	execute.Run(run, nil, puzzle, true)
}
