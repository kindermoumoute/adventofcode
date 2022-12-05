package main

import (
	"fmt"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	line := parse(input)

	part1, part2 := 0, 0

	for j, line := range line {
		if contains(line[0], line[1]) || contains(line[1], line[0]) {
			part1++
		}
		if overlaps(line[0], line[1]) {
			part2++
			fmt.Println(j)
		}
	}

	return part1, part2
}

func contains(a, b Range) bool {
	return a.Left >= b.Left && a.Right <= b.Right
}
func overlaps(a, b Range) bool {
	return a.Left <= b.Right && b.Left <= a.Right
}

type Range struct {
	Left, Right int
}

func parse(s string) [][2]Range {
	lines := strings.Split(s, "\n")
	r := [][2]Range(nil)
	for _, line := range lines {
		s := [2]Range{}
		pkg.MustScanf(line, "%d-%d,%d-%d", &s[0].Left, &s[0].Right, &s[1].Left, &s[1].Right)
		r = append(r, s)
	}
	return r
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
