package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	spreadSheet := parse(puzzle)
	fmt.Println(spreadSheet.CheckSumPart1())
	fmt.Println(spreadSheet.CheckSumPart2())

}

type SpreadSheet [][]int

func (s SpreadSheet) CheckSumPart1() int {
	cs := 0
	for i := range s {
		cs += s[i][len(s[i])-1] - s[i][0]
	}
	return cs
}

func (s SpreadSheet) CheckSumPart2() int {
	cs := 0
	for i, r := range s {
		for j := len(r) - 1; j > 0; j-- {
			for k := 0; k < j; k++ {
				if s[i][j]%s[i][k] == 0 {
					cs += s[i][j] / s[i][k]
					goto omgthisisagoto
				}
			}
		}
	omgthisisagoto:
	}
	return cs
}

func parse(p string) SpreadSheet {
	lines := strings.Split(p, "\n")
	s := make(SpreadSheet, len(lines))

	for i, l := range lines {
		line := strings.Split(l, "\t")
		s[i] = []int{}
		for _, r := range line {
			cell, _ := strconv.Atoi(r)
			s[i] = append(s[i], cell)
		}
		sort.Ints(s[i])
	}
	return s
}
