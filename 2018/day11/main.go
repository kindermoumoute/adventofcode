package main

import (
	"strconv"

	"github.com/kindermoumoute/adventofcode/pkg/execute"

	"github.com/kindermoumoute/adventofcode/pkg"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	serial, err := strconv.Atoi(input)
	pkg.Check(err)

	scores := map[pkg.P]int{
		pkg.P{X: 0, Y: 0}: 0,
		pkg.P{X: 0, Y: 1}: 0,
		pkg.P{X: 1, Y: 0}: 0,
	}
	for i := 1; i <= 300; i++ {
		for j := 1; j <= 300; j++ {
			rackID := i + 10
			score := rackID * j
			score += serial
			score *= rackID
			score = (score / 100) % 10
			score -= 5
			scores[pkg.P{X: i, Y: j}] = score + scores[pkg.P{X: i, Y: j - 1}] + scores[pkg.P{X: i - 1, Y: j}] - scores[pkg.P{X: i - 1, Y: j - 1}]
		}
	}

	maxXPart1, maxYPart1, maxPart1 := 0, 0, -99999999
	maxXPart2, maxYPart2, maxPart2, maxSizePart2 := 0, 0, -99999999, 1
	for size := 1; size <= 300; size++ {
		for i := size; i <= 300; i++ {
			for j := size; j <= 300; j++ {
				sum := scores[pkg.P{X: i, Y: j}] - scores[pkg.P{X: i, Y: j - size}] - scores[pkg.P{X: i - size, Y: j}] + scores[pkg.P{X: i - size, Y: j - size}]
				if size == 3 && sum > maxPart1 {
					maxPart1 = sum
					maxXPart1 = i - 2
					maxYPart1 = j - 2
				}
				if sum > maxPart2 {
					maxPart2 = sum
					maxXPart2 = i - size + 1
					maxYPart2 = j - size + 1
					maxSizePart2 = size
				}
			}
		}
	}
	return strconv.Itoa(maxXPart1) + "," + strconv.Itoa(maxYPart1), strconv.Itoa(maxXPart2) + "," + strconv.Itoa(maxYPart2) + "," + strconv.Itoa(maxSizePart2)
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
