package main

import (
	"fmt"
	"strconv"
	"strings"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	grid := parse(input)
	ection := 0
	part1 := 0
	pos := p{0, 0}
	for i := 0; i < 10000; i++ {
		if _, exist := grid[pos]; !exist {
			grid[pos] = '.'
		}
		switch grid[pos] {
		case '#':
			ection = (ection + 1) % 4
			grid[pos] = '.'
		case '.':
			part1++
			ection = (4 + ection - 1) % 4
			grid[pos] = '#'
		}
		pos.x += dir[ection].x
		pos.y += dir[ection].y
	}

	grid = parse(input)
	ection = 0
	part2 := 0
	pos = p{0, 0}
	for i := 0; i < 10000000; i++ {
		_, exist := grid[pos]
		if !exist {
			grid[pos] = '.'
		}

		switch grid[pos] {
		case '.':
			ection = (4 + ection - 1) % 4
			grid[pos] = 'W'
		case 'W':
			part2++
			grid[pos] = '#'
		case '#':
			ection = (ection + 1) % 4
			grid[pos] = 'F'
		case 'F':
			ection = (ection + 2) % 4
			grid[pos] = '.'
		}
		pos.x += dir[ection].x
		pos.y += dir[ection].y
	}

	return strconv.Itoa(part1), strconv.Itoa(part2)
}

var dir = []p{
	p{0, 1},
	p{1, 0},
	p{0, -1},
	p{-1, 0},
}

type p struct {
	x, y int
}

func parse(s string) map[p]byte {
	myMap := make(map[p]byte)
	lines := strings.Split(s, "\n")
	offset := len(lines) / 2
	for y, line := range lines {
		b := []byte(line)
		for x, v := range b {
			myMap[p{x - offset, len(lines) - y - offset - 1}] = v
		}
	}
	return myMap
}

func check(err error) {
	if err != nil {
		panic(err)
	}
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
