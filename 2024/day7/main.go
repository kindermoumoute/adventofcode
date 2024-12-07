package main

import (
	"strconv"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	tv := parse(input)

	part1, part2 := 0, 0
	for _, value := range tv {
		if canCompute(value.value, value.list, false) {
			part1 += value.value
		}
		if canCompute(value.value, value.list, true) {
			part2 += value.value
		}
	}

	return part1, part2
}

func canCompute(value int, list []int, part2 bool) bool {
	switch {
	case len(list) == 1:
		return value == list[0]
	case canCompute(value, append([]int{list[0] + list[1]}, list[2:]...), part2):
		return true
	case canCompute(value, append([]int{list[0] * list[1]}, list[2:]...), part2):
		return true
	case part2 && canCompute(value, append([]int{pkg.MustAtoi(strconv.Itoa(list[0]) + strconv.Itoa(list[1]))}, list[2:]...), part2):
		return true
	}
	return false
}

type testValues struct {
	value int
	list  []int
}

func parse(s string) []testValues {
	tv := []testValues(nil)
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		tmp := strings.Split(line, ": ")
		tv = append(tv, testValues{
			value: pkg.MustAtoi(tmp[0]),
			list:  pkg.ParseIntList(tmp[1], " "),
		})
	}
	return tv
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
