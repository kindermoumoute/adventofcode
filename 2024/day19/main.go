package main

import (
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	patterns, towels := parse(input)

	part1, part2 := 0, 0
	for _, towel := range towels {
		if l := countTowelPatterns(towel, patterns); l > 0 {
			part1++
			part2 += l
		}
	}

	return part1, part2
}

var cache = map[string]int{}

func countTowelPatterns(s string, patterns *pattern) int {
	if len(s) == 0 {
		return 1
	}

	if c, exist := cache[s]; exist {
		return c
	}

	part2 := 0
	for _, nextPattern := range patterns.findNextPatterns(s) {
		part2 += countTowelPatterns(s[len(nextPattern):], patterns)
	}

	cache[s] = part2
	return part2
}

type pattern struct {
	nextPatterns map[int32]*pattern
	val          string
}

func (p *pattern) findNextPatterns(s string) []string {
	nextPatterns := []string(nil)
	if p.val != "" {
		nextPatterns = append(nextPatterns, p.val)
	}
	if len(s) == 0 {
		return nextPatterns
	}

	if nextP, exist := p.nextPatterns[int32(s[0])]; exist {
		nextPatterns = append(nextPatterns, nextP.findNextPatterns(s[1:])...)
	}

	return nextPatterns
}

func parse(s string) (*pattern, []string) {
	tmp := strings.Split(s, "\n\n")
	rawPatterns := strings.Split(tmp[0], ", ")
	patterns := &pattern{
		nextPatterns: make(map[int32]*pattern),
	}
	for _, rawPattern := range rawPatterns {
		p := patterns
		for _, char := range rawPattern {
			if _, exist := p.nextPatterns[char]; !exist {
				p.nextPatterns[char] = &pattern{
					nextPatterns: make(map[int32]*pattern),
				}
			}
			p = p.nextPatterns[char]
		}
		p.val = rawPattern
	}
	return patterns, strings.Split(tmp[1], "\n")
}

func main() {
	execute.Run(run, nil, puzzle, true)
}
