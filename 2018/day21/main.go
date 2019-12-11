package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
)

type Registers struct {
	ip int
	r  []int
}

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	instructions, insPointer := parse(input)
	part1 := 99999999999999999
	part2 := -1
	fewestIns := 100000000
	valuesMap := map[int]struct{}{}
	values := []int{0}
dance:
	for i := 0; i < len(values); i++ {
		time := 0
		registers := &Registers{
			ip: insPointer,
			r:  make([]int, 6),
		}
		registers.r[0] = values[i]

		for registers.r[registers.ip] >= 0 && registers.r[registers.ip] < len(instructions) {
			if registers.r[registers.ip] == 28 && i == 0 {
				_, exist := valuesMap[registers.r[1]]
				if exist {
					continue dance
				}
				valuesMap[registers.r[1]] = struct{}{}
				part2 = registers.r[1]
				values = append(values, registers.r[1])
			}
			instructions[registers.r[registers.ip]].hit++
			ins := instructions[registers.r[registers.ip]]
			switch ins.ins {
			case "addr":
				registers.addr(ins.a, ins.b, ins.c)
			case "addi":
				registers.addi(ins.a, ins.b, ins.c)
			case "setr":
				registers.setr(ins.a, ins.b, ins.c)
			case "seti":
				registers.seti(ins.a, ins.b, ins.c)
			case "muli":
				registers.muli(ins.a, ins.b, ins.c)
			case "mulr":
				registers.mulr(ins.a, ins.b, ins.c)
			case "bani":
				registers.bani(ins.a, ins.b, ins.c)
			case "banr":
				registers.banr(ins.a, ins.b, ins.c)
			case "bori":
				registers.bori(ins.a, ins.b, ins.c)
			case "borr":
				registers.borr(ins.a, ins.b, ins.c)
			case "gtrr":
				registers.gtrr(ins.a, ins.b, ins.c)
			case "gtir":
				registers.gtir(ins.a, ins.b, ins.c)
			case "eqrr":
				registers.eqrr(ins.a, ins.b, ins.c)
			case "eqri":
				registers.eqri(ins.a, ins.b, ins.c)
			default:
				panic("forget " + ins.ins)
			}
			registers.r[registers.ip]++
			time++
			if i != 0 && time == fewestIns {
				continue dance
			}
		}
		if fewestIns > time {
			fewestIns = time
			part1 = values[i]
		} else if fewestIns == time && part1 > values[i] {
			part1 = values[i]
		}
	}
	return strconv.Itoa(part1), strconv.Itoa(part2) // never happen
}

type Instruction struct {
	hit     int
	ins     string
	a, b, c int
}

func parse(s string) ([]Instruction, int) {
	lines := strings.Split(s, "\n")
	IPN := 0
	n, err := fmt.Sscanf(lines[0], "#ip %d", &IPN)
	pkg.Check(err)
	if n != 1 {
		panic(fmt.Errorf("%d args expected in scanf, got %d", 0, n))
	}
	IPList := make([]Instruction, len(lines)-1)
	for i, line := range lines[1:] {
		n, err := fmt.Sscanf(line, "%s %d %d %d", &IPList[i].ins, &IPList[i].a, &IPList[i].b, &IPList[i].c)
		pkg.Check(err)
		if n != 4 {
			panic(fmt.Errorf("%d args expected in scanf, got %d", 0, n))
		}
	}
	return IPList, IPN
}

func main() {
	pkg.Execute(run, tests, puzzle, true)
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
