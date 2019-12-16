package main

import (
	"math"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {

	initList := make([]int, len(input))
	for i := range initList {
		initList[i] = int(input[i] - '0')
	}

	////
	// part 1
	////

	phase := initList[:]

	// calculate phase
	for i := 0; i < 100; i++ {
		phase = getPhase(phase)
	}

	// output result
	part1 := getIntInList(phase, 0, 7)

	////
	// part 2
	////

	offset := getIntInList(initList, 0, 6)
	phase = make([]int, len(initList)*10000)
	for i := range phase {
		phase[i] = initList[i%len(initList)]
	}

	// calculate phase
	for i := 0; i < 100; i++ {
		phase = getPhase2(phase)
	}

	// output result
	offset %= len(phase)
	part2 := getIntInList(phase, offset, offset+7)

	return part1, part2
}

func getIntInList(list []int, start, end int) int {
	sum := 0
	for i, v := range list[start : end+1] {
		sum += v * (int(math.Pow(10, float64((end-start)-i))))
	}
	return sum
}

// phase 2 only compute phase for phase[n] where n>len(phase)/2
func getPhase2(phase []int) []int {
	newPhase := make([]int, len(phase))
	sum := pkg.Sum(phase...)
	for j := range newPhase {
		newPhase[j] = sum % 10
		sum -= phase[j]
	}
	return newPhase
}

func getPhase(phase []int) []int {
	newPhase := make([]int, len(phase))
	for j := range newPhase {
		for k := range phase {
			newPhase[j] += phase[k] * pattern(j, k)
		}
		newPhase[j] = pkg.Abs(newPhase[j]) % 10
	}
	return newPhase
}

func pattern(j, k int) int {
	j++
	k = (k + 1) % (4 * j)
	switch {
	case k < j:
		return 0
	case k < 2*j:
		return 1
	case k < 3*j:
		return 0
	case k < 4*j:
		return -1
	}
	return 0
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
