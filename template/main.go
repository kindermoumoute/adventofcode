package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Run(s string) (string, string) {
	//list := parse(puzzle)
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
		part1, part2 := Run(test.input)
		if part1 != test.expectedPart1 {
			fmt.Println("PART1: ", part1, " but expected ", test.expectedPart1)
			return
		}
		if part2 != test.expectedPart2 {
			fmt.Println("PART2: ", part2, " but expected ", test.expectedPart2)
			return
		}
	}
	fmt.Println(Run(puzzle))
}
