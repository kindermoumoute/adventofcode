package main

import (
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	part1, part2 := 0, 0
	for _, line := range strings.Split(input, "\n") {
		part1Numbers := []int(nil)
		part2Numbers := []int(nil)
		for i := 0; i < len(line); i++ {
			if isDigit(line[i]) {
				part1Numbers = append(part1Numbers, int(line[i]-'0'))
				part2Numbers = append(part2Numbers, int(line[i]-'0'))
			} else if ok, value := isCharDigit(line, i); ok {
				part2Numbers = append(part2Numbers, value)
			}
		}
		if len(part1Numbers) > 0 {
			part1 += part1Numbers[0]*10 + part1Numbers[len(part1Numbers)-1]
		}
		if len(part2Numbers) > 0 {
			part2 += part2Numbers[0]*10 + part2Numbers[len(part2Numbers)-1]
		}
	}

	return part1, part2
}

func isDigit(c uint8) bool {
	return c >= '0' && c <= '9'
}

func isCharDigit(line string, i int) (bool, int) {
	for word, value := range charDigits {
		if i+len(word) > len(line) {
			continue
		}
		if line[i:i+len(word)] == word {
			return true, value
		}
	}

	return false, 0

}

var charDigits = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
