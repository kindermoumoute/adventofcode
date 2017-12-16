package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	p := newParser()
	part2 := p.parse(puzzle)
	part1 := -999999999
	for _, value := range p.variable {
		if value > part1 {
			part1 = value
		}
	}
	fmt.Println(part1, part2)
}

type parser struct {
	fn       map[string]func(int, int) int
	compare  map[string]func(int, int) bool
	variable map[string]int
}

func (p parser) parse(s string) int {
	highest := -999999999
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		instruction := strings.Split(line, " if ")
		fnArg := strings.Split(instruction[0], " ")
		cond := strings.Split(instruction[1], " ")
		a, exist := p.variable[cond[0]]
		if !exist {
			p.variable[cond[0]] = 0
			a = 0
		}
		b, _ := strconv.Atoi(cond[2])
		if p.compare[cond[1]](a, b) {
			a := p.variable[fnArg[0]]
			b, _ := strconv.Atoi(fnArg[2])
			p.variable[fnArg[0]] = p.fn[fnArg[1]](a, b)
			if p.variable[fnArg[0]] > highest {
				highest = p.variable[fnArg[0]]
			}
		}
	}
	return highest
}

func newParser() parser {
	fn := map[string]func(int, int) int{
		"inc": inc,
		"dec": dec,
	}
	compare := map[string]func(int, int) bool{
		">":  gt,
		">=": ge,
		"==": e,
		"<":  lt,
		"<=": le,
		"!=": ne,
	}
	return parser{fn, compare, make(map[string]int)}
}

func inc(a, b int) int {
	return a + b
}

func dec(a, b int) int {
	return a - b
}

func gt(a, b int) bool {
	return a > b
}

func ge(a, b int) bool {
	return a >= b
}

func e(a, b int) bool {
	return a == b
}

func lt(a, b int) bool {
	return a < b
}

func le(a, b int) bool {
	return a <= b
}

func ne(a, b int) bool {
	return a != b
}
