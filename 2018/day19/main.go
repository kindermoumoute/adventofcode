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
func run(input string) (string, string) {
	instructions, insPointer := parse(input)
	part1 := 0
	for i := 0; i < 2; i++ {

		registers := &Registers{
			ip: insPointer,
			r:  make([]int, 6),
		}
		registers.r[0] = i
		for registers.r[registers.ip] >= 0 && registers.r[registers.ip] < len(instructions) {
			if registers.r[registers.ip] == 2 && i == 1 {
				fmt.Println("start p2", registers.r)
				part2 := 0
				for j := 1; j < registers.r[4]+1; j++ {
					if registers.r[4]%j == 0 {
						part2 += j
					}
				}
				fmt.Println(registers)
				return strconv.Itoa(part1), strconv.Itoa(part2)
			}

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
			case "gtrr":
				registers.gtrr(ins.a, ins.b, ins.c)
			case "eqrr":
				registers.eqrr(ins.a, ins.b, ins.c)
			default:
				panic("forget")
			}
			registers.r[registers.ip]++
		}
		part1 = registers.r[0]
	}
	return strconv.Itoa(0), "" // never happen
}

type Instruction struct {
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

func (r *Registers) eqrr(a, b, c int) {
	if r.r[b] == r.r[a] {
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
