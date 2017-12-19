package main

import (
	"fmt"
	"strconv"
	"strings"
)

type p struct {
	x, y int
	path byte
}

var (
	dir = []p{
		{1, 0, '-'},
		{0, 1, '|'},
		{-1, 0, '-'},
		{0, -1, '|'},
	}
)

// returns part1 and part2
func run(input string) (string, string) {
	p := parse(input)
	part1 := []byte{}
	part2 := 0
	for p.path != ' ' {

		// part 1 & 2
		if p.isLetter(p.path) {
			part1 = append(part1, p.path)
		}
		part2++

		// walk and rotate
		p.walk()
		if p.path == '+' {
			p.findPath()
		}
	}
	return string(part1), strconv.Itoa(part2)
}

type walker struct {
	p
	dirI int
	tab  [][]byte
}

func (p *walker) findPath() {
	p.rotateLeft()
	for !p.followsRoute() && !p.isLetter(p.next()) {
		p.rotateRight()
	}
}

func (p *walker) walk() {
	p.x += dir[p.dirI].x
	p.y += dir[p.dirI].y
	p.path = p.tab[p.y][p.x]
}

func (p *walker) rotateRight() {
	p.dirI = (p.dirI + 1) % len(dir)
}

func (p *walker) rotateLeft() {
	p.dirI = (p.dirI + 3) % len(dir)
}

func (p *walker) isLetter(b byte) bool {
	return b >= 'A' && b <= 'Z'
}

func (p *walker) next() byte {
	return p.tab[p.y+dir[p.dirI].y][p.x+dir[p.dirI].x]
}

func (p *walker) followsRoute() bool {
	return p.next() == dir[p.dirI].path
}

func parse(s string) *walker {
	add := []byte{}
	for i := 0; i < 200; i++ {
		add = append(add, ' ')
	}

	lines := strings.Split(s, "\n")
	list := make([][]byte, len(lines))
	for i, line := range lines {
		list[i] = append([]byte(line), add...)
	}
	p := &walker{tab: list}
	p.path = '|'
	p.dirI = 1
	for i := range p.tab[0] {
		if p.tab[0][i] == '|' {
			p.x = i
		}
	}
	return p
}

func main() {
	for _, test := range tests {
		part1, part2 := run(test.input)
		if part1 != test.expectedPart1 {
			fmt.Println("Input ", test.input, " - PART1: ", part1, " but expected ", test.expectedPart1)
			return
		}
		if part2 != test.expectedPart2 {
			fmt.Println("Input ", test.input, " - PART2: ", part2, " but expected ", test.expectedPart2)
			return
		}
	}
	fmt.Println(run(puzzle))
}
