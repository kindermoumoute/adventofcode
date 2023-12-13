package main

import (
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg/execute"
	"github.com/kindermoumoute/adventofcode/pkg/twod"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	part1, part2 := 0, 0

	for _, pattern := range strings.Split(input, "\n\n") {
		m := twod.NewMapFromInput(pattern)
		topRight := m.FindTopRight()

		if nbLeftColumns := findReflection(m, topRight, false, 0); nbLeftColumns != 0 {
			part1 += nbLeftColumns
		}
		if nbLeftColumns := findReflection(m, topRight, false, 1); nbLeftColumns != 0 {
			part2 += nbLeftColumns
		}

		if nbBottomRows := findReflection(m, topRight, true, 0); nbBottomRows != 0 {
			part1 += (topRight.Y() - nbBottomRows + 1) * 100 // we want the upper rows, not bottom
		}
		if nbBottomRows := findReflection(m, topRight, true, 1); nbBottomRows != 0 {
			part2 += (topRight.Y() - nbBottomRows + 1) * 100 // we want the upper rows, not bottom
		}
	}

	return part1, part2
}

func findReflection(m twod.Map, topRight twod.Vector, isHorizontal bool, allowedMistake int) int {
	coord := func(x, y int) twod.Vector {
		if isHorizontal {
			return twod.NewVector(y, x)
		}
		return twod.NewVector(x, y)
	}
	topRight = coord(topRight.X(), topRight.Y())
nextMirror:
	for i := 1; i <= topRight.X(); i++ {
		foundMistake := 0
		for j := 0; j+i <= topRight.X() && i-j-1 >= 0; j++ {
			for k := 0; k <= topRight.Y(); k++ {
				right, exist1 := m[coord(j+i, k)]
				left, exist2 := m[coord(i-j-1, k)]
				if !exist1 || !exist2 {
					panic("not happen")
				}
				if left != right {
					foundMistake++
					if foundMistake > allowedMistake {
						continue nextMirror
					}
				}
			}
		}
		if foundMistake != allowedMistake {
			continue
		}
		return i
	}
	return 0
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
