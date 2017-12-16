package main

import (
	"fmt"
	"strconv"
	"strings"
)

func run(input string) (string, string) {
	c := parse(input)
	c.register["c"] = 1 // remove to get part 1
	for i := 0; i < len(c.instruction); i++ {
		ins := c.instruction[i]
		switch ins[0] {
		case "cpy":
			value, err := strconv.Atoi(ins[1])
			if err != nil {
				value = c.register[ins[1]]
			}
			c.register[ins[2]] = value
		case "inc":
			c.register[ins[1]]++
		case "dec":
			c.register[ins[1]]--
		case "jnz":
			test, err := strconv.Atoi(ins[1])
			if err != nil {
				test = c.register[ins[1]]
			}
			if test != 0 {
				value, err := strconv.Atoi(ins[2])
				check(err)
				i += value - 1
			}
		}
	}
	return strconv.Itoa(c.register["a"]), ""
}

type computer struct {
	register    map[string]int
	instruction [][]string
}

func parse(s string) computer {
	c := computer{register: make(map[string]int), instruction: [][]string{}}
	for _, line := range strings.Split(s, "\n") {
		c.instruction = append(c.instruction, strings.Split(line, " "))
	}
	return c
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
