package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
)

// returns part1 and part2
func run(input string) (string, string) {
	part1, part2 := parse(input)
	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func parse(s string) (int, int) {
	var (
		currentPosition = pkg.NewPoint()
		visited         = map[struct{ x, y int }]struct{}{}
	)

	lines := strings.Split(s, ", ")
	firstRevisited := 0
	for _, line := range lines {
		if len(line) == 0 {
			fmt.Println(line)
			continue
		}
		if line[0] == 'R' {
			currentPosition.Right()
		} else {
			currentPosition.Left()
		}
		nb, err := strconv.Atoi(line[1:])
		pkg.Check(err)

		for i := 0; i < nb; i++ {
			_, alreadyVisited := visited[struct{ x, y int }{currentPosition.X, currentPosition.Y}]
			if alreadyVisited && firstRevisited == 0 {
				firstRevisited = currentPosition.DistFromOrigin()
			}
			visited[struct{ x, y int }{currentPosition.X, currentPosition.Y}] = struct{}{}
			currentPosition.Move(1)
		}
	}

	return currentPosition.DistFromOrigin(), firstRevisited
}

func main() {
	for _, test := range tests {
		part1, part2 := run(test.input)
		if part1 != test.expectedPart1 {
			fmt.Println("Input ", test.input, " - PART1: ", part1, " but expected ", test.expectedPart1)
			return
		}
		if part2 != test.expectedPart2 {
			fmt.Println("Input ", test.input, " - PART2: ", part2, " but expected ", test.expectedPart2)
			return
		}
	}
	fmt.Println(run(puzzle))
}
