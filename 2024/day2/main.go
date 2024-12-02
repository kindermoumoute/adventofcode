package main

import (
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	reports := parse(input)

	part1, part2 := 0, 0
	for _, report := range reports {
		if isValid(report) {
			part1++
			part2++
			continue
		}
		for _, alternateReport := range generateAlternateReports(report) {
			if isValid(alternateReport) {
				part2++
				break
			}
		}
	}

	return part1, part2
}

func generateAlternateReports(report []int) [][]int {
	var result [][]int
	for i := 0; i < len(report); i++ {
		newSlice := append([]int{}, report[:i]...)
		newSlice = append(newSlice, report[i+1:]...)
		result = append(result, newSlice)
	}
	return result
}

func isValid(report []int) bool {
	positives := map[int]struct{}{1: {}, 2: {}, 3: {}}
	negatives := map[int]struct{}{-1: {}, -2: {}, -3: {}}
	for i := 1; i < len(report); i++ {
		positives[report[i]-report[i-1]] = struct{}{}
		negatives[report[i]-report[i-1]] = struct{}{}
	}
	return len(positives) == 3 || len(negatives) == 3
}

func parse(s string) [][]int {
	reports := [][]int(nil)
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		reports = append(reports, pkg.ParseIntList(line, " "))
	}
	return reports
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
