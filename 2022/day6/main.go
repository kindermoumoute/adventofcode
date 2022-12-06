package main

import (
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	return findMarker(input, 4), findMarker(input, 14)
}

func findMarker(input string, uniqueCharCount int) int {
	for position := range input {
		uniqueChar := make(map[uint8]struct{})
		for i := position; i >= 0 && i > position-uniqueCharCount; i-- {
			uniqueChar[input[i]] = struct{}{}
		}
		if len(uniqueChar) == uniqueCharCount {
			return position + 1
		}
	}
	panic("no marker")
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
