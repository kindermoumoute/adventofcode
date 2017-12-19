package main

import (
	"fmt"
	"strconv"
	"strings"
)

type p struct {
	x, y  int
	route byte
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
	tab := parse(input)
	word := []byte{}
	part2 := 0
	x, y := 0, 0
	dirI := 1
	for i := range tab[0] {
		if tab[0][i] == '|' {
			x = i
		}
	}
	for {
		if tab[y][x] >= 'A' && tab[y][x] <= 'Z' {
			word = append(word, tab[y][x])
		}
		x += dir[dirI].x
		y += dir[dirI].y
		part2++
		if tab[y][x] == '+' {
			dirI = (dirI + 2) % len(dir)
			for {
				dirI = (dirI + 1) % len(dir)
				if tab[y+dir[dirI].y][x+dir[dirI].x] == dir[dirI].route {
					break
				}

				if tab[y+dir[dirI].y][x+dir[dirI].x] >= 'A' && tab[y+dir[dirI].y][x+dir[dirI].x] <= 'Z' {
					break
				}
			}
		} else if tab[y][x] == ' ' {
			break
		}

	}
	return string(word), strconv.Itoa(part2)
}

func parse(s string) [][]byte {
	add := []byte{}
	for i := 0; i < 200; i++ {
		add = append(add, ' ')
	}

	lines := strings.Split(s, "\n")
	list := make([][]byte, len(lines))
	for i, line := range lines {
		list[i] = append([]byte(line), add...)
	}

	return list
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
