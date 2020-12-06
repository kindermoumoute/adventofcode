package main

import (
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {

	part1, part2 := 0, 0

	for _, group := range strings.Split(input, "\n\n") {
		uniq := make(map[int32]int)
		people := strings.Split(group, "\n")
		for _, person := range people {
			for _, c := range person {
				uniq[c]++
				if uniq[c] == len(people) {
					part2++
				}
			}
		}
		part1 += len(uniq)
	}

	return part1, part2
}

func main() {
	execute.Run(run, nil, puzzle, true)
}
