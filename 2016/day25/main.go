package main

import (
	"fmt"
	"strconv"
	"strings"
)

func run(input string) (string, string) {
	lowest := 0
dance:
	for ; ; lowest++ {
		c := parse(input)
		c.register["a"] = lowest
		signal := 0
		for i := 0; i < len(c.instruction); i++ {
			ins := c.instruction[i]
			switch ins[0] {
			case "out":
				value, err := strconv.Atoi(ins[1])
				if err != nil {
					value = c.register[ins[1]]
				}
				if value > 1 {
					break
				} else {
					signal = signal<<1 + value
				}
			case "cpy":
				_, err := strconv.Atoi(ins[2])
				if err != nil {
					value, err := strconv.Atoi(ins[1])
					if err != nil {
						c.register[ins[2]] = c.register[ins[1]]
					} else {
						c.register[ins[2]] = value
					}
				}
			case "inc":
				_, err := strconv.Atoi(ins[1])
				if err != nil {
					c.register[ins[1]]++
					//if i+2 < len(c.instruction) && c.instruction[i+1][0] == "dec" && c.instruction[i+2][0] == "jnz" && c.instruction[i+2][2] == "-2" {
					//	_, err = strconv.Atoi(c.instruction[i+1][1])
					//	if err != nil {
					//		c.register[ins[1]] *= c.register[c.instruction[i+1][1]]
					//		c.register[c.instruction[i+1][1]] = 0
					//		i += 2
					//	}
					//} else {
					//	c.register[ins[1]]++
					//}
				}

			case "dec":
				_, err := strconv.Atoi(ins[1])
				if err != nil {
					c.register[ins[1]]--
					//if i+2 < len(c.instruction) && c.instruction[i+1][0] == "inc" && c.instruction[i+2][0] == "jnz" && c.instruction[i+2][2] == "-2" {
					//	_, err = strconv.Atoi(c.instruction[i+1][1])
					//	if err != nil {
					//		c.register[c.instruction[i+1][1]] *= c.register[ins[1]]
					//		c.register[ins[1]] = 0
					//		i += 2
					//	}
					//} else {
					//	c.register[ins[1]]--
					//}
				}
			case "jnz":
				test, err := strconv.Atoi(ins[1])
				if err != nil {
					test = c.register[ins[1]]
				}
				if test != 0 {
					value, err := strconv.Atoi(ins[2])
					if err != nil {
						value = c.register[ins[2]]
					}
					i += value - 1
				}
				//case "tgl":
				//	away, err := strconv.Atoi(ins[1])
				//	if err != nil {
				//		away = c.register[ins[1]]
				//	}
				//	away += i
				//	if away < len(c.instruction) {
				//		if len(c.instruction[away]) == 2 {
				//			if c.instruction[away][0] == "inc" {
				//				c.instruction[away][0] = "dec"
				//			} else {
				//				c.instruction[away][0] = "inc"
				//			}
				//		} else {
				//			if c.instruction[away][0] == "jnz" {
				//				c.instruction[away][0] = "cpy"
				//			} else {
				//				c.instruction[away][0] = "jnz"
				//			}
				//		}
				//	}
			}
			if signal == 0x5 {
				break dance
			}
			//fmt.Println(c.instruction[i], c.register)
		}
		fmt.Println(lowest)
	}
	return strconv.Itoa(lowest), ""
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
