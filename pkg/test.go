package pkg

import (
	"fmt"
	"os"
)

type TestCases []TestCase

type TestCase struct {
	Input, ExpectedPart1, ExpectedPart2 string
}

func (t TestCases) Run(fn func(string) (string, string), hideInput bool) {
	for _, test := range t {
		part1, part2 := fn(test.Input)
		passedPart1 := part1 == test.ExpectedPart1 || test.ExpectedPart1 == ""
		passedPart2 := part2 == test.ExpectedPart2 || test.ExpectedPart2 == ""
		passed := passedPart1 && passedPart2

		if !passed && !hideInput {
			fmt.Println("Input ", test.Input)
		}
		if !passedPart1 {
			fmt.Println(" - PART1: ", part1, " but expected ", test.ExpectedPart1)
			os.Exit(1)
		}
		if !passedPart2 {
			fmt.Println(" - PART2: ", part2, " but expected ", test.ExpectedPart2)
			os.Exit(1)
		}
	}
}
