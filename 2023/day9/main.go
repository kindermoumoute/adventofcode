package main

import (
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
	"github.com/samber/lo"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	numbers := parse(input)

	part1, part2 := 0, 0

	for _, list := range numbers {
		part1 += getLastSum(list)
		part2 += getLastSum(lo.Reverse(list))
	}

	return part1, part2
}

func getLastSum(list []int) int {
	allLists := [][]int{list}
	currentList := 0
	for fullZeroes := false; !fullZeroes; {
		allLists = append(allLists, []int{})
		fullZeroes = true
		for i := 0; i < len(allLists[currentList])-1; i++ {
			val := allLists[currentList][i+1] - allLists[currentList][i]
			allLists[currentList+1] = append(allLists[currentList+1], val)
			fullZeroes = fullZeroes && (val == 0)
		}
		currentList++
	}
	sumLast := 0
	for i := currentList; i >= 0; i-- {
		sumLast += allLists[i][len(allLists[i])-1]
		allLists[i] = append(allLists[i], sumLast)
	}
	return sumLast
}

func parse(s string) [][]int {
	lines := strings.Split(s, "\n")
	numbers := [][]int(nil)
	for _, line := range lines {
		numbers = append(numbers, pkg.ParseIntList(line, " "))
	}
	return numbers
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
