package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	puzzle := "5	1	10	0	1	7	13	14	3	12	8	10	7	12	0	6"
	//puzzle := "2	4	1	2"
	memoryBanks := parse(puzzle)
	states := make(map[string]int)
	for {
		states[sliceToString(memoryBanks)]++
		if states[sliceToString(memoryBanks)] > 2 {
			break
		}
		topBank := 0

		for i := range memoryBanks {
			if memoryBanks[i] > memoryBanks[topBank] {
				topBank = i
			}
		}
		blocksToDistribute := memoryBanks[topBank]
		memoryBanks[topBank] = 0
		for blocksToDistribute > 0 {
			topBank = (topBank + 1) % len(memoryBanks)
			memoryBanks[topBank]++
			blocksToDistribute--
		}
	}
	part2 := 0
	for _, n := range states {
		if n > 1 {
			part2++
		}
	}
	fmt.Println(len(states))
	fmt.Println(part2)

}
func sliceToString(s []int) string {
	str := ""
	for _, n := range s {
		str += strconv.Itoa(n) + "-"
	}
	return str
}
func parse(p string) []int {
	tmp := strings.Split(p, "\t")
	numbers := make([]int, len(tmp))
	for i, n := range tmp {
		numbers[i], _ = strconv.Atoi(n)
	}
	return numbers
}
