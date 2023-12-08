package main

import (
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	road, steps := parse(input)
	part1, part2 := 0, 0

	part1 = solve(road, steps, "AAA", func(s string) bool {
		return s == "ZZZ"
	})

	part2 = 1
	for node := range road {
		if node[2] == 'A' {
			part2 = pkg.LCM(part2, solve(road, steps, node, func(s string) bool {
				return s[2] == 'Z'
			}))
		}
	}

	return part1, part2
}

func solve(road map[string]Step, steps string, cursor string, stopCondition func(string) bool) int {
	count := 0
	for {
		for _, way := range steps {
			count++
			switch way {
			case 'L':
				cursor = road[cursor].L
			case 'R':
				cursor = road[cursor].R
			}
			if stopCondition(cursor) {
				return count
			}
		}
	}
}

type Step struct {
	L, R string
}

func parse(s string) (map[string]Step, string) {
	lines := strings.Split(s, "\n")
	steps := make(map[string]Step)
	for _, line := range lines[2:] {
		tmp := strings.Split(line, " = ")
		tmp2 := strings.Split(tmp[1][1:len(tmp[1])-1], ", ")
		steps[tmp[0]] = Step{
			L: tmp2[0],
			R: tmp2[1],
		}
	}
	return steps, lines[0]
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
