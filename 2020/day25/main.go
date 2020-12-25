package main

import (
	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	a, b := parse(input)

	nicePrimaryNumber := 20201227
	loop := 0
	for transformed := 1; transformed != a; loop++ {
		transformed = transformed * 7 % nicePrimaryNumber
	}

	part1 := 1
	for i := 0; i < loop; i++ {
		part1 = part1 * b % nicePrimaryNumber
	}

	return part1, 0
}

func parse(s string) (int, int) {
	tmp := pkg.ParseIntList(s, "\n")
	return tmp[0], tmp[1]
}

func main() {
	execute.Run(run, nil, puzzle, true)
}
