package main

import (
	"fmt"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {

	// fmt.Println("PASS")
	return parse(input)
}

var mapPart1 = [][]string{
	{"1", "2", "3"},
	{"4", "5", "6"},
	{"7", "8", "9"},
}
var mapPart2 = [][]string{
	{"", "", "1", "", ""},
	{"", "2", "3", "4", ""},
	{"5", "6", "7", "8", "9"},
	{"", "A", "B", "C", ""},
	{"", "", "D", "", ""},
}

func parse(s string) (string, string) {
	lines := strings.Split(s, "\n")
	part1 := ""
	for _, line := range lines {
		p := pkg.NewPoint()
		for _, dir := range line {
			switch dir {
			case 'R':
				p.MoveRight(1)
			case 'U':
				p.MoveUp(1)
			case 'D':
				p.MoveDown(1)
			case 'L':
				p.MoveLeft(1)
			default:
				panic(fmt.Errorf("wrong direction %c", dir))
			}
			if pkg.Abs(p.X) > 1 || pkg.Abs(p.Y) > 1 {
				p.Move(-1)
			}
		}
		part1 += p.FindInMap(mapPart1, pkg.P{X: 1, Y: 1})
	}

	// part 2
	part2 := ""
	for _, line := range lines {
		p := pkg.NewPoint()
		p.X = -2
		for _, dir := range line {
			switch dir {
			case 'R':
				p.MoveRight(1)
			case 'U':
				p.MoveUp(1)
			case 'D':
				p.MoveDown(1)
			case 'L':
				p.MoveLeft(1)
			default:
				panic(fmt.Errorf("wrong direction %c", dir))
			}
			if p.DistFromOrigin() > 2 {
				p.Move(-1)
			}
		}
		part2 += p.FindInMap(mapPart2, pkg.P{X: 2, Y: 2})
	}
	return part1, part2
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
