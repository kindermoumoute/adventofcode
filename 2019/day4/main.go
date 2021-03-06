package main

import (
	"strconv"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg/execute"

	"github.com/kindermoumoute/adventofcode/pkg"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	a, b := parse(input)
	part1, part2 := 0, 0
	for i := a; i <= b; i++ {
		password := []byte(strconv.Itoa(i))
		if checkPassword(password, true) {
			part1++
		}
		if checkPassword(password, false) {
			part2++
		}
	}

	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func parse(s string) (int, int) {
	numbers := strings.Split(s, "-")
	return pkg.MustAtoi(numbers[0]), pkg.MustAtoi(numbers[1])
}

func checkPassword(password []byte, part1 bool) bool {
	if len(password) < 6 { // 6 characters
		return false
	}

	countChar := map[byte]int{}
	for i := 0; i < len(password); i++ {
		if i != 0 && password[i-1] > password[i] { // no increasing order
			return false
		}
		countChar[password[i]]++
	}

	for _, count := range countChar {
		if count == 2 || part1 && count >= 2 { // exactly 2 matching characters
			return true
		}
	}

	return false
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
