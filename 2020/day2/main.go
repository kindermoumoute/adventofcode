package main

import (
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	passwords := parse(input)

	part1, part2 := 0, 0
	for _, p := range passwords {
		c := strings.Count(p.password, string(p.char))
		if c >= p.min && c <= p.max {
			part1++
		}
	}
	for _, p := range passwords {
		match := 0
		if len(p.password) >= p.max && p.password[p.max-1] == p.char {
			match++
		}
		if len(p.password) >= p.min && p.password[p.min-1] == p.char {
			match++
		}
		if match == 1 {
			part2++
		}
	}

	return part1, part2
}

type password struct {
	min, max int
	char     byte
	password string
}

func parse(s string) []password {
	passwords := []password(nil)
	for _, line := range strings.Split(s, "\n") {
		p := password{}
		pkg.MustScanf(line, "%d-%d %c: %s", &p.min, &p.max, &p.char, &p.password)
		passwords = append(passwords, p)
	}
	return passwords
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
