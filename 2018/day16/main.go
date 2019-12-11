package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	var r = &Registers{r: make([]int, 4)}
	var Ops = []func(int, int, int){r.eqri, r.banr, r.bori, r.mulr, r.seti, r.bani, r.muli, r.gtrr, r.setr, r.addi, r.gtir, r.borr, r.addr, r.eqrr, r.gtri, r.eqir}
	part1, part2 := parse(input)
	part1Score := 0
dance:
	for _, p1 := range part1 {
		count := 0
		for _, op := range Ops {
			r.SetRegisters(p1.in)
			op(p1.a, p1.b, p1.c)
			if r.Equals(p1.out) {
				count++
				if count >= 3 {
					part1Score++
					continue dance
				}
			}
		}

	}
	r.SetRegisters([]int{0, 0, 0, 0})
	for _, ins := range part2 {
		Ops[ins.op](ins.a, ins.b, ins.c)
	}
	return strconv.Itoa(part1Score), strconv.Itoa(r.r[0])
}

type Part1 struct {
	in          []int
	out         []int
	op, a, b, c int
}

type Instruction struct {
	op      int
	a, b, c int
}

func parse(s string) ([]Part1, []Instruction) {
	part := strings.Split(s, "\n\n\n\n")
	part1 := strings.Split(part[0], "\n")
	part2 := strings.Split(part[1], "\n")
	p1 := make([]Part1, len(part1)/4+1)
	for i := 0; i < len(part1); i += 4 {
		p1[i/4].in = make([]int, 4)
		n, err := fmt.Sscanf(part1[i], "Before: [%d,%d,%d,%d]", &p1[i/4].in[0], &p1[i/4].in[1], &p1[i/4].in[2], &p1[i/4].in[3])
		pkg.Check(err)
		if n != 4 {
			panic(fmt.Errorf("%d args expected in scanf, got %d", 0, n))
		}

		n, err = fmt.Sscanf(part1[i+1], "%d %d %d %d", &p1[i/4].op, &p1[i/4].a, &p1[i/4].b, &p1[i/4].c)
		pkg.Check(err)
		if n != 4 {
			panic(fmt.Errorf("%d args expected in scanf, got %d", 0, n))
		}

		p1[i/4].out = make([]int, 4)
		n, err = fmt.Sscanf(part1[i+2], "After:  [%d, %d, %d, %d]", &p1[i/4].out[0], &p1[i/4].out[1], &p1[i/4].out[2], &p1[i/4].out[3])
		pkg.Check(err)
		if n != 4 {
			panic(fmt.Errorf("%d args expected in scanf, got %d", 0, n))
		}
	}
	p2 := make([]Instruction, len(part2))
	for i := range part2 {
		n, err := fmt.Sscanf(part2[i], "%d %d %d %d", &p2[i].op, &p2[i].a, &p2[i].b, &p2[i].c)
		pkg.Check(err)
		if n != 4 {
			panic(fmt.Errorf("%d args expected in scanf, got %d", 0, n))
		}
	}
	return p1, p2
}

func main() {
	pkg.Execute(run, nil, puzzle, true)
}

type Registers struct {
	ip int
	r  []int
}

func (r *Registers) SetRegisters(init []int) {
	for i := 0; i < 4; i++ {
		r.r[i] = init[i]
	}
}

func (r *Registers) Equals(out []int) bool {
	sameRegister := 0
	for i := 0; i < 4; i++ {
		if r.r[i] == out[i] {
			sameRegister++
		}
	}
	return sameRegister == 4
}

func (r *Registers) setr(a, b, c int) {
	r.r[c] = r.r[a]
}

func (r *Registers) seti(a, b, c int) {
	r.r[c] = a
}
func (r *Registers) addr(a, b, c int) {
	r.r[c] = r.r[a] + r.r[b]
}

func (r *Registers) addi(a, b, c int) {
	r.r[c] = r.r[a] + b
}

func (r *Registers) mulr(a, b, c int) {
	r.r[c] = r.r[a] * r.r[b]
}

func (r *Registers) muli(a, b, c int) {
	r.r[c] = r.r[a] * b
}

func (r *Registers) banr(a, b, c int) {
	r.r[c] = r.r[a] & r.r[b]
}

func (r *Registers) bani(a, b, c int) {
	r.r[c] = r.r[a] & b
}

func (r *Registers) borr(a, b, c int) {
	r.r[c] = r.r[a] | r.r[b]
}

func (r *Registers) bori(a, b, c int) {
	r.r[c] = r.r[a] | b
}

func (r *Registers) eqrr(a, b, c int) {
	if r.r[b] == r.r[a] {
		r.r[c] = 1
		return
	}
	r.r[c] = 0
}

func (r *Registers) eqri(a, b, c int) {
	if b == r.r[a] {
		r.r[c] = 1
		return
	}
	r.r[c] = 0
}
func (r *Registers) eqir(a, b, c int) {
	if a == r.r[b] {
		r.r[c] = 1
		return
	}
	r.r[c] = 0
}
func (r *Registers) gtrr(a, b, c int) {
	if r.r[b] < r.r[a] {
		r.r[c] = 1
		return
	}
	r.r[c] = 0
}

func (r *Registers) gtir(a, b, c int) {
	if r.r[b] < a {
		r.r[c] = 1
		return
	}
	r.r[c] = 0
}

func (r *Registers) gtri(a, b, c int) {
	if b < r.r[a] {
		r.r[c] = 1
		return
	}
	r.r[c] = 0
}
