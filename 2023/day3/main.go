package main

import (
	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
	"github.com/kindermoumoute/adventofcode/pkg/twod"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	part1, part2 := 0, 0

	m := twod.NewMapFromInput(input)
	m = m.FilterOut('.')
	firstLetterDigits := make(twod.Map)
	symbols := make(twod.Map)
	for pos, v := range m {
		if !isDigit(v) {
			symbols[pos] = v
			continue
		}

		digitPos := findFullDigit(m, pos)
		str := ""
		for _, dp := range digitPos {
			str += string(uint8(m[dp].(int32)))
		}
		firstLetterDigits[digitPos[0]] = pkg.MustAtoi(str)
	}

	firstLettersPart1 := make(twod.Map)
	for pos, sym := range symbols {
		firstLetters := make(twod.Map)
		for _, direction := range twod.AllDirections {
			digitPos := findFullDigit(m, pos+direction)
			if len(digitPos) > 0 {
				firstLetters[digitPos[0]] = firstLetterDigits[digitPos[0]]
				firstLettersPart1[digitPos[0]] = firstLetterDigits[digitPos[0]]
			}
		}

		if sym == '*' && len(firstLetters) == 2 {
			s := 1
			for _, v := range firstLetters {
				s *= v.(int)
			}
			part2 += s
		}
	}

	for _, v := range firstLettersPart1 {
		part1 += v.(int)
	}

	return part1, part2
}

func findFullDigit(m twod.Map, pos twod.Vector) []twod.Vector {
	firstDigit := pos
	for i := 0; i < 4; i++ {
		if c, exist := m[pos+twod.LEFT.ScalarMult(i)]; exist && isDigit(c) {
			firstDigit = pos + twod.LEFT.ScalarMult(i)
		} else {
			break
		}
	}
	digits := []twod.Vector(nil)
	for i := 0; i < 4; i++ {
		if c, exist := m[firstDigit+twod.RIGHT.ScalarMult(i)]; exist && isDigit(c) {
			digits = append(digits, firstDigit+twod.RIGHT.ScalarMult(i))
		} else {
			break
		}
	}
	return digits
}

func isDigit(v any) bool {
	c := v.(int32)
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
