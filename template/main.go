package main

import (
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	//list := parse(input)

	part1, part2 := 0, 0

	//// 2D lib
	//
	//m := twod.NewMapFromInput(input)
	//m:= make( twod.Map)
	//twod.NewPoint()
	////

	return part1, part2
}

func parse(s string) {
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		pkg.MustScanf(line, "")
	}
	return
}

func main() {
	execute.Run(run, nil, puzzle, true)
}
