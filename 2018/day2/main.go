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
	twoTimes := 0
	threeTimes := 0
	for _, line := range lines {
		mapLineChar[line] = make(map[int32]int)
		for _, char := range line {
			mapLineChar[line][char]++
		}
		countTwoExist, countThreeExist := false, false
		for _, count := range mapLineChar[line] {
			if !countTwoExist && count == 2 {
				countTwoExist = true
				twoTimes++
			}
			if !countTwoExist && count == 2 {
				countTwoExist = true
				twoTimes++
			}
			if !countThreeExist && count == 3 {
				countThreeExist = true
				threeTimes++
			}
			if countTwoExist && countThreeExist {
				break
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
	fmt.Println(run(puzzle))
}
