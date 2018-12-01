package main

import (
	"fmt"
	"strconv"
	"strings"
)

var Salt = []int{17, 31, 73, 47, 23}

func Run(input string) (string, string) {
	part1Input := []int{}
	for _, number := range strings.Split(input, ",") {
		n, _ := strconv.Atoi(number)
		part1Input = append(part1Input, n)
	}
	part2Input := []int{}
	for _, number := range input {
		part2Input = append(part2Input, int(number))
	}
	part2Input = append(part2Input, Salt...)
	list := newList(256)
	knotHash(list, part1Input, 0, 0)

	part1 := fmt.Sprint(list[0] * list[1])

	list = newList(256)
	pos, skipSize := 0, 0
	for i := 0; i < 64; i++ {
		pos, skipSize = knotHash(list, part2Input, pos, skipSize)
	}
	hash := make([]int, 16)
	for i, v := range list {
		hash[i/len(hash)] ^= v
	}
	part2 := ""
	for _, h := range hash {
		if h < 16 {
			part2 += fmt.Sprintf("0")
		}
		part2 += fmt.Sprintf("%x", h)
	}
	return part1, part2
}

func newList(ln int) []int {
	list := make([]int, ln)
	for i := range list {
		list[i] = i
	}
	return list
}

func knotHash(list, puzzle []int, pos, skipSize int) (int, int) {
	for _, length := range puzzle {
		for i := 0; i < length/2; i++ {
			swapStart := (pos + i) % len(list)
			swapEnd := (pos + length - i - 1) % len(list)
			list[swapStart], list[swapEnd] = list[swapEnd], list[swapStart]
		}
		pos = (pos + length + skipSize) % len(list)
		skipSize = (skipSize + 1) % len(list)
	}
	return pos, skipSize
}

func main() {
	for _, test := range tests {
		part1, part2 := Run(test.input)
		if part1 != test.expectedPart1 {
			fmt.Println("Input ", test.input, " - PART1: ", part1, " but expected ", test.expectedPart1)
			return
		}
		if part2 != test.expectedPart2 {
			fmt.Println("Input ", test.input, " - PART2: ", part2, " but expected ", test.expectedPart2)
			return
		}
	}
	fmt.Println(Run(puzzle))
}
