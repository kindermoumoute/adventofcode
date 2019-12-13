package main

import (
	"strconv"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/intcode"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {

	// part 1
	incCode := intcode.New(input, 0)
	incCode.Memory[1] = 12
	incCode.Memory[2] = 2
	incCode.Run()
	part1 := incCode.Memory[0]

	// part 2
	part2 := 0
dance:
	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			newIntCode := intcode.New(input, 0)
			newIntCode.Memory[1] = i
			newIntCode.Memory[2] = j
			newIntCode.Run()
			if newIntCode.Memory[0] == 19690720 {
				part2 = i*100 + j
				break dance
			}
		}
	}

	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func main() {
	pkg.Execute(run, tests, puzzle, true)
}
