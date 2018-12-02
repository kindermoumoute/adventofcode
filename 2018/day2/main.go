package main

import (
	"fmt"
	"strconv"
	"strings"
)

// returns part1 and part2
func run(input string) (string, string) {
	part1, part2 := parse(input)

	return strconv.Itoa(part1), part2
}

func parse(s string) (int, string) {
	lines := strings.Split(s, "\n")

	// part1
	mapLineChar := make(map[string]map[int32]int)
	twoTimes, threeTimes := 0, 0
	for _, line := range lines {
		mapLineChar[line] = make(map[int32]int)
		for _, char := range line {
			mapLineChar[line][char]++
		}
		countTwoExist, countThreeExist := false, false
		for _, count := range mapLineChar[line] {
			switch {
			case countTwoExist && countThreeExist:
				break
			case !countTwoExist && count == 2:

				countTwoExist = true
				twoTimes++
			case !countThreeExist && count == 3:
				countThreeExist = true
				threeTimes++
			}
		}
	}

	//part2
	commonLetters := ""
dance:
	for _, line := range lines {
		for _, line2 := range lines {
			miss := -1
			for i := range line {
				if line[i] != line2[i] {
					if miss != -1 {
						miss = -1
						break
					}
					miss = i
				}
			}
			if miss != -1 {
				commonLetters = line[:miss] + line[miss+1:]
				break dance
			}
		}
	}
	return twoTimes * threeTimes, commonLetters
}

func main() {
	tests.Run(run, false)
	fmt.Println(run(puzzle))
}
