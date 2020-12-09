package main

import (
	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	list := pkg.ParseIntList(puzzle, "\n")

	part1, part2 := 0, 0
part1:
	for i, number := range list {
		if i < 25 {
			continue
		}
		previousNumbers := []int(nil)
		for j := i - 1; i-j <= 25; j-- {
			previousNumbers = append(previousNumbers, list[j])
		}
		for k, n1 := range previousNumbers {
			for j, n2 := range previousNumbers {
				if k == j {
					continue
				}
				if n1+n2 == number {
					continue part1
				}
			}
		}
		part1 = number
		break
	}

dance:
	for i := 0; i < len(list); i++ {
		contiguous := []int(nil)
		sum := 0
		for j := i; j < len(list); j++ {
			if sum == part1 {
				part2 = pkg.Min(contiguous...) + pkg.Max(contiguous...)
				break dance
			}
			if sum > part1 {
				break
			}

			sum += list[j]
			contiguous = append(contiguous, list[j])
		}
	}

	return part1, part2
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
