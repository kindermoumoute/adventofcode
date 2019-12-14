package execute

import (
	"fmt"
	"os"
)

type TestCases []TestCase

type TestCase struct {
	Input                        string
	ExpectedPart1, ExpectedPart2 interface{}
}

func (t TestCases) Run(fn func(string) (interface{}, interface{}), hideInput bool) {
	for _, test := range t {
		part1I, part2I := fn(test.Input)
		part1, expectedPart1, part2, expectedPart2 := fmt.Sprint(part1I), fmt.Sprint(test.ExpectedPart1), fmt.Sprint(part2I), fmt.Sprint(test.ExpectedPart2)
		passedPart1 := part1 == expectedPart1 || expectedPart1 == ""
		passedPart2 := part2 == expectedPart2 || expectedPart2 == ""
		passed := passedPart1 && passedPart2

		if !passed && !hideInput {
			fmt.Println("Input ", test.Input)
		}
		if !passedPart1 {
			fmt.Println(" - PART1: ", part1, " but expected ", expectedPart1)
			os.Exit(1)
		}
		if !passedPart2 {
			fmt.Println(" - PART2: ", part2, " but expected ", expectedPart2)
			os.Exit(1)
		}
	}
}
