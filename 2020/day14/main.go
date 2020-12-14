package main

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	part1, part2 := 0, 0

	mem := make(map[int]int)
	andMask, orMask := 0, 0

	regMask := regexp.MustCompile(`mask = (.*)`)
	regMem := regexp.MustCompile(`mem\[(\d+)] = (\d+)`)
	for _, line := range strings.Split(input, "\n") {
		switch {
		case regMask.MatchString(line):
			andMask, orMask = toMasks(regMask.FindAllStringSubmatch(line, -1))
		case regMem.MatchString(line):
			key, value := toMem(regMem.FindAllStringSubmatch(line, -1))
			mem[key] = value&andMask | orMask
		}
	}
	for _, v := range mem {
		part1 += v
	}

	mem = make(map[int]int)
	orMasks := []int(nil)
	andMask = 0
	for _, line := range strings.Split(input, "\n") {
		switch {
		case regMask.MatchString(line):
			orMasks, andMask = toOrMasks(regMask.FindAllStringSubmatch(line, -1))
		case regMem.MatchString(line):
			key, value := toMem(regMem.FindAllStringSubmatch(line, -1))
			for _, mask := range orMasks {
				mem[key&andMask|mask] = value
			}
		}
	}
	for _, v := range mem {
		part2 += v
	}

	return part1, part2
}

func toMem(matches [][]string) (int, int) {
	key := matches[0][1]
	value := matches[0][2]
	return pkg.MustAtoi(key), pkg.MustAtoi(value)
}

func toMasks(matches [][]string) (int, int) {
	str := matches[0][1]
	and := strings.ReplaceAll(str, "X", "1")
	or := strings.ReplaceAll(str, "X", "0")
	andMask, err := strconv.ParseInt(and, 2, 64)
	if err != nil {
		panic(err)
	}
	orMask, err := strconv.ParseInt(or, 2, 64)
	if err != nil {
		panic(err)
	}
	return int(andMask), int(orMask)
}

func toOrMasks(matches [][]string) ([]int, int) {
	str := matches[0][1]
	or := strings.ReplaceAll(str, "X", "0")
	tmp, err := strconv.ParseInt(or, 2, 64)
	if err != nil {
		panic(err)
	}
	orMask := int(tmp)
	masks := []int{orMask}
	for i, v := range str {
		if v == 'X' {
			xOrMask := 1 << (len(str) - 1 - i)
			xAndMask := 0xFFFFFFFFFFFFFFF ^ xOrMask
			lnMasks := len(masks)
			for j := 0; j < lnMasks; j++ {
				masks = append(masks, masks[j]|xOrMask)
				masks[j] = masks[j] & xAndMask
			}
		}
	}

	andStr := strings.ReplaceAll(str, "0", "1")
	andStr = strings.ReplaceAll(andStr, "X", "0")
	andMask, err := strconv.ParseInt(andStr, 2, 64)
	if err != nil {
		panic(err)
	}

	return masks, int(andMask)
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
