package main

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	part1, part2 := 0, 0
	for _, line := range strings.Split(input, "\n") {
		lineTotal, _ := evalPart1(line)
		part1 += lineTotal
		part2 += pkg.MustAtoi(evalPart2(line))
	}
	return part1, part2
}

var (
	regAdd         = regexp.MustCompile(`(\d+ \+ \d+)+`)
	regMul         = regexp.MustCompile(`(\d+ \* \d+)+`)
	regNum         = regexp.MustCompile(`\d+`)
	regParenthesis = regexp.MustCompile(`\(([^\(\)]+)\)`)
)

func evalPart2(exp string) string {
	for strings.Contains(exp, "(") {
		exp = regParenthesis.ReplaceAllStringFunc(exp, func(s string) string {
			return evalNoParenthesis(s[1 : len(s)-1])
		})
	}
	return evalNoParenthesis(exp)
}

func evalNoParenthesis(exp string) string {
	for strings.Contains(exp, "+") {
		exp = regAdd.ReplaceAllStringFunc(exp, func(s string) string {
			matches := regNum.FindAllString(s, 2)
			return strconv.Itoa(pkg.MustAtoi(matches[0]) + pkg.MustAtoi(matches[1]))
		})
	}
	for strings.Contains(exp, "*") {
		exp = regMul.ReplaceAllStringFunc(exp, func(s string) string {
			matches := regNum.FindAllString(s, 2)
			return strconv.Itoa(pkg.MustAtoi(matches[0]) * pkg.MustAtoi(matches[1]))
		})
	}
	return exp
}

func evalPart1(exp string) (int, int) {
	total, previousOp := 0, uint8('+')
	for i := 0; i < len(exp); i++ {
		switch exp[i] {
		case '(':
			subTotal, offset := evalPart1(exp[i+1:])
			total = execOp(previousOp, total, subTotal)
			i += offset
		case ')':
			return total, i + 1
		case ' ':
			continue
		case '+', '*':
			previousOp = exp[i]
		default:
			total = execOp(previousOp, total, int(exp[i])-'0')
		}
	}
	return total, len(exp)
}

func execOp(op uint8, left, right int) int {
	switch op {
	case '+':
		return left + right
	case '*':
		return left * right
	}
	panic("invalid op")
}

func main() {
	execute.Run(run, tests, puzzle, false)
}
