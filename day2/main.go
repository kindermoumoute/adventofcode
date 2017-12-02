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

type SpreadSheet struct {
	Row     []Row
	cs, cs2 int
}

type Row struct {
	Cell []int
	div  int
}

func (s *SpreadSheet) CheckSumPart1() int {
	for i := range s.Row {
		s.cs += s.Row[i].Cell[len(s.Row[i].Cell)-1] - s.Row[i].Cell[0]
	}
	return s.cs
}

func (s *SpreadSheet) CheckSumPart2() int {
	for i, r := range s.Row {
		for j := len(r.Cell) - 1; j > 0; j-- {
			for k := 0; k < j; k++ {
				if s.Row[i].Cell[j]%s.Row[i].Cell[k] == 0 {
					s.cs2 += s.Row[i].Cell[j] / s.Row[i].Cell[k]
					goto omgthisisagoto
				}
			}
		}
	omgthisisagoto:
	}
	return s.cs2
}

func parse(p string) *SpreadSheet {
	ss := &SpreadSheet{}
	lines := strings.Split(p, "\n")
	ss.Row = make([]Row, len(lines))

	for i, l := range lines {
		line := strings.Split(l, "\t")
		ss.Row[i].Cell = []int{}
		for _, r := range line {
			cell, _ := strconv.Atoi(r)
			ss.Row[i].Cell = append(ss.Row[i].Cell, cell)
		}
		sort.Ints(ss.Row[i].Cell)
	}
	return ss
}
