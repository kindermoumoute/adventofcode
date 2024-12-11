package main

import (
	"strconv"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	part1, part2 := 0, 0
	stonesMemo := make(map[[2]int]int)
	for _, v := range pkg.ParseIntList(input, " ") {
		part1 += blink(v, 25, stonesMemo)
	}
	for _, v := range pkg.ParseIntList(input, " ") {
		part2 += blink(v, 75, stonesMemo)
	}
	return part1, part2
}

func blink(stone, blinks int, countMemo map[[2]int]int) (out int) {
	if blinks == 0 {
		return 1
	}

	if c, found := countMemo[[2]int{stone, blinks}]; found {
		return c
	}
	defer func() {
		countMemo[[2]int{stone, blinks}] = out
	}()

	if stone == 0 {
		return blink(1, blinks-1, countMemo)
	}

	if str := strconv.Itoa(stone); len(str)%2 == 0 {
		left := blink(pkg.MustAtoi(str[:len(str)/2]), blinks-1, countMemo)
		right := blink(pkg.MustAtoi(str[len(str)/2:]), blinks-1, countMemo)
		return left + right
	}

	return blink(stone*2024, blinks-1, countMemo)
}

func main() {
	execute.Run(run, nil, puzzle, true)
}
