package main

import (
	"fmt"
	"strconv"
)

type cBuffer struct {
	buffer []int
	pos    int
}

func run(input string) (string, string) {
	forwardSteps := parse(input)
	buf := cBuffer{buffer: []int{0}, pos: 0}
	part1, part2 := -1, -1
	lenBuf := len(buf.buffer)
	for i := 1; ; i++ {
		buf.pos = (buf.pos + forwardSteps) % lenBuf
		//fmt.Println("i=", i, " pos ", buf.pos, " ", forwardSteps, len(buf.buffer), (i+forwardSteps)%len(buf.buffer))
		if i == 2017 {
			part1 = buf.buffer[buf.pos]
			break
		}
		if part1 == -1 {
			buf.buffer = append(buf.buffer, 0)
			copy(buf.buffer[buf.pos+1:], buf.buffer[buf.pos:])
			buf.buffer[buf.pos] = i
		}

		buf.pos++
		lenBuf++
	}

	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func parse(s string) int {
	nb, err := strconv.Atoi(s)
	check(err)
	return nb
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
