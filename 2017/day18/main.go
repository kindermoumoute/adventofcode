package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

const timeOut = 10 * time.Millisecond

type val struct {
	sync.Mutex
	ue int
}

func run(input string) (string, string) {
	c := computer{0, []program{0: parse(input), 1: parse(input)}}
	for t := range c.p {
		p := c.p[t]
		p.ID = t
		p.register["p"] = t
		go func() {
			timer := time.NewTimer(timeOut)
		dance:
			for i := 0; i < len(p.instruction); i++ {
				ins := p.instruction[i]
				switch ins[0] {
				case "set":
					p.register[ins[1]] = p.getValue(ins[2])
				case "add":
					p.register[ins[1]] += p.getValue(ins[2])
				case "mul":
					p.register[ins[1]] *= p.getValue(ins[2])
				case "mod":
					p.register[ins[1]] %= p.getValue(ins[2])
				case "snd":
					if p.ID == 1 {
						c.part2++
					}
					c.p[1-p.ID].receiver <- p.getValue(ins[1])
				case "rcv":
					select {
					case value := <-p.receiver:
						timer = time.NewTimer(timeOut)
						p.register[ins[1]] = value
					case <-timer.C:
						break dance
					}
				case "jgz":
					test := p.getValue(ins[1])
					if test > 0 {
						i += p.getValue(ins[2]) - 1
					}
				}
			}
			p.done <- true
		}()
	}
	for _, p := range c.p {
		<-p.done
	}
	return "", strconv.Itoa(c.part2)
}

type computer struct {
	part2 int
	p     []program
}

type program struct {
	ID          int
	register    map[string]int
	instruction [][]string
	receiver    chan int

	done chan bool
}

func (c program) getValue(s string) int {
	value, err := strconv.Atoi(s)
	if err != nil {
		value = c.register[s]
	}
	return value
}

func parse(s string) program {
	c := program{
		register:    make(map[string]int),
		instruction: [][]string{},
		receiver:    make(chan int, 99999),
		done:        make(chan bool),
	}
	for _, line := range strings.Split(s, "\n") {
		c.instruction = append(c.instruction, strings.Split(line, " "))
	}
	return c
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

// part 1 - unoptimized
//func run(input string) (string, string) {
//	sound, part1 := "", ""
//	c := parse(input)
//	//c.register["c"] = 1 // remove to get part 1
//dance:
//	for i := 0; i < len(c.instruction); i++ {
//		if i < 15 {
//			fmt.Println(i, " : ", c.instruction[i], " a=", c.register["a"])
//		}
//		ins := c.instruction[i]
//		switch ins[0] {
//		case "set":
//			value, err := strconv.Atoi(ins[2])
//			if err != nil {
//				value = c.register[ins[2]]
//			}
//			c.register[ins[1]] = value
//			fmt.Println("VALUE ", value, " : ", ins[2])
//		case "add":
//			value, err := strconv.Atoi(ins[2])
//			if err != nil {
//				value = c.register[ins[2]]
//			}
//			c.register[ins[1]] += value
//		case "mul":
//			value, err := strconv.Atoi(ins[2])
//			if err != nil {
//				value = c.register[ins[2]]
//			}
//			c.register[ins[1]] *= value
//		case "mod":
//			value, err := strconv.Atoi(ins[2])
//			if err != nil {
//				value = c.register[ins[2]]
//			}
//			c.register[ins[1]] %= value
//		case "snd":
//			value, err := strconv.Atoi(ins[1])
//			if err != nil {
//				value = c.register[ins[1]]
//			}
//			sound = strconv.Itoa(value)
//		case "rcv":
//			value, err := strconv.Atoi(ins[1])
//			if err != nil {
//				value = c.register[ins[1]]
//			}
//			if value != 0 {
//				part1 = sound
//				break dance
//			}
//			//break dance
//		case "jgz":
//			test, err := strconv.Atoi(ins[1])
//			if err != nil {
//				test = c.register[ins[1]]
//			}
//			fmt.Println("register ", ins[1], " : ", test)
//			if test > 0 {
//				value, err := strconv.Atoi(ins[2])
//				check(err)
//				i += value - 1
//			}
//		}
//	}
//	fmt.Println(part1)
//	return part1, ""
//}
