package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
)

// returns part1 and part2
func run(input string) (string, string) {
	list := parse(input)
	sum := 0
	part2 := 10000000000
	mapFreq := make(map[int]int)
	mapFreq[0]++
dance:
	for {
		for _, number := range list {
			sum += number
			_, exist := mapFreq[sum]
			if exist && part2 == 10000000000 {
				part2 = sum
				break dance
			}
			mapFreq[sum]++
		}
	}
	return strconv.Itoa(pkg.Sum(list...)), strconv.Itoa(part2)
}

func parse(s string) []int {
	lines := strings.Split(s, "\n")
	list := make([]int, len(lines))
	for i, line := range lines {
		nb, err := strconv.Atoi(line)
		pkg.Check(err)
		list[i] = nb
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
			fmt.Println("Input "+" - PART2: ", part2, " but expected ", test.expectedPart2)
			return
		}
	}
	fmt.Println(run(puzzle))
}
