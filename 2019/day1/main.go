package main

import (
	"strconv"

	"github.com/kindermoumoute/adventofcode/pkg"
)

// returns part1 and part2
func run(input string) (string, string) {
	part1, part2 := 0, 0
	for _, fuelValue := range pkg.ParseIntList(input, "\n") {
		part1 += fuelValue/3 - 2
		part2 += pkg.RecursiveSum(fuelValue, func(i int) (int, bool) {
			i = i/3 - 2
			return i, i <= 0
		})
	}
	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func main() {
	pkg.Execute(run, tests, puzzle, true)
}
