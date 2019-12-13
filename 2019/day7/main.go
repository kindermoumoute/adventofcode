package main

import (
	"strconv"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/intcode"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	part1 := RunPart1(pkg.GetAllPermutations("01234"), input)
	part2 := RunPart2(pkg.GetAllPermutations("98765"), input)

	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func RunPart1(perms []string, program string) int {
	max := 0
	for _, permutation := range perms {
		additionalInput := 0
		for i, input := range permutation {
			c := intcode.New(program, 0, int(input-'0'), additionalInput)
			c.Name = strconv.Itoa(i) + "[" + string(input) + "]"
			additionalInput = c.Run()
			if additionalInput > max {
				max = additionalInput
			}
		}
	}
	return max
}

func RunPart2(settings []string, program string) int {
	max := 0
	for _, setting := range settings {
		// create a program for each setting character
		programs := []*intcode.IntCode(nil)
		for i, input := range setting {
			programs = append(programs, intcode.New(program, 0, int(input-'0')))
			programs[i].Name = strconv.Itoa(i) + "[" + string(input) + "]"
		}
		programs[0].Input.Buff = append(programs[0].Input.Buff, 0) // first program must have a second input set to 0

		// match outputs with inputs
		for i, c := range programs {
			programs[(i-1+len(programs))%len(programs)].IntCodeToIntCodeProtocol(c.Input.C)
		}

		// run all programs but the last
		for _, c := range programs[:len(programs)-1] {
			c.RunBackground()
		}

		max = pkg.Max(max, programs[len(programs)-1].Run())
	}
	return max
}

func main() {
	pkg.Execute(run, tests, puzzle, true)
}
