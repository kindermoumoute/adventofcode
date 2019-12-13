package main

import (
	"strconv"

	"github.com/kindermoumoute/adventofcode/pkg/intcode"

	"github.com/kindermoumoute/adventofcode/pkg"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	return strconv.Itoa(intcode.New(input, 0, 1).Run()), strconv.Itoa(intcode.New(input, 0, 2).Run())
}

func main() {
	pkg.Execute(run, nil, puzzle, true)
}
