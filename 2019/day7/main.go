package main

import (
	"strconv"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/intcode"
)

// returns part1 and part2
func run(input string) (string, string) {
	intList := pkg.ParseIntMap(input, ",")
	part1 := RunPart1(pkg.GetAllPermutations("01234"), intList)
	part2 := RunPart2(pkg.GetAllPermutations("98765"), intList)

	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func RunPart1(perms []string, intList map[int]int) int {
	max := 0
	for _, permutation := range perms {
		additionalInput := 0
		for i, input := range permutation {
			c := intcode.New(intList, 0, int(input-'0'), additionalInput)
			c.Name = strconv.Itoa(i) + "[" + string(input) + "]"
			additionalInput = c.Run()
			if additionalInput > max {
				max = additionalInput
			}
		}
	}
	return max
}

func RunPart2(settings []string, intList map[int]int) int {
	max := 0
	for _, setting := range settings {
		// create a program for each setting character
		programs := []*intcode.IntCode(nil)
		for i, input := range setting {
			inputs := []int{int(input - '0')}
			if i == 0 {
				inputs = append(inputs, 0) // first program must have a second input set to 0
			}
			programs = append(programs, intcode.New(intList, 0, inputs...))
			programs[i].Name = strconv.Itoa(i) + "[" + string(input) + "]"
		}

		// match outputs with inputs
		for i := range programs {
			programs[(i-1+len(programs))%len(programs)].Output.C = programs[i].Input.C
		}

		// run
		for i, c := range programs {
			if i == len(programs)-1 {
				max = pkg.Max(c.Run(), max)
			} else {
				c.RunBackground()
			}
		}
	}
	return max
}

func main() {
	pkg.Execute(run, tests, puzzle, true)
}
