package main

import (
	"strconv"

	"github.com/kindermoumoute/adventofcode/pkg/execute"

	"github.com/kindermoumoute/adventofcode/pkg/intcode"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	return strconv.Itoa(intcode.New(input, 0, 1).Run()), strconv.Itoa(intcode.New(input, 0, 2).Run())
}

func main() {
	execute.Run(run, nil, puzzle, true)
}
