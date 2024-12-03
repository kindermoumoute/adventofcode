package main

import (
	"regexp"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	instructions := strings.Split(input, "\n")
	part1, part2 := 0, 0

	do := true
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	re2 := regexp.MustCompile(`(do\(\)|don't\(\)|mul\((\d+),(\d+)\))`)
	for _, instruction := range instructions {
		for _, s := range re.FindAllStringSubmatch(instruction, -1) {
			part1 += pkg.MustAtoi(s[1]) * pkg.MustAtoi(s[2])
		}

		for _, s := range re2.FindAllStringSubmatch(instruction, -1) {
			if strings.HasPrefix(s[0], "mul(") && do {
				part2 += pkg.MustAtoi(s[2]) * pkg.MustAtoi(s[3])
			} else if strings.HasPrefix(s[0], "do(") {
				do = true
			} else if strings.HasPrefix(s[0], "don't(") {
				do = false
			}
		}
	}

	return part1, part2
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
