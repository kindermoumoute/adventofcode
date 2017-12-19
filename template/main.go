package main

import (
	"fmt"
	"strconv"
	"strings"
)

// returns part1 and part2
func run(input string) (string, string) {
	//list := parse(input)

	// fmt.Println("PASS")
	return "", ""
}

func parse(s string) []int {
	lines := strings.Split(s, "\n")
	list := make([]int, len(lines))
	for i, line := range lines {
		nb, err := strconv.Atoi(line)
		check(err)
		list[i] = nb
	}
	return list
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
