package main

import (
	"slices"
	"sort"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	orderingRules, updates := parse(input)

	part1, part2 := 0, 0
	for _, update := range updates {
		tmp := make([]int, len(update))
		copy(tmp, update)

		sort.Slice(tmp, func(i, j int) bool {
			return orderingRules[tmp[i]][tmp[j]]
		})

		middlePage := tmp[len(tmp)/2]
		if slices.Equal(tmp, update) {
			part1 += middlePage
		} else {
			part2 += middlePage
		}
	}

	return part1, part2
}

func parse(s string) (map[int]map[int]bool, [][]int) {
	tmp := strings.Split(s, "\n\n")

	orderingRules := make(map[int]map[int]bool)
	for _, line := range strings.Split(tmp[0], "\n") {
		a, b := 0, 0
		pkg.MustScanf(line, "%d|%d", &a, &b)
		_, exist := orderingRules[a]
		if !exist {
			orderingRules[a] = make(map[int]bool)
		}
		orderingRules[a][b] = true
	}

	updates := [][]int(nil)
	for _, line := range strings.Split(tmp[1], "\n") {
		updates = append(updates, pkg.ParseIntList(line, ","))
	}

	return orderingRules, updates
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
