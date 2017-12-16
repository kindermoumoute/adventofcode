package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	part1, picosec := 0, 0

	list := parse(puzzle)

	// part 1
	for _, l := range list {
		if sevScore(picosec, l.depth, l.size) == 0 {
			part1 += l.depth * l.size
		}
	}
	// part 2
	for severity := true; severity; picosec++ {
		severity = false
		for _, l := range list {
			if sevScore(picosec, l.depth, l.size) == 0 {
				severity = true
				break
			}
		}
	}
	fmt.Println(part1, picosec-1)
}
func sevScore(p, d, s int) int {
	return (p + d) % ((s - 1) * 2)
}

type layer struct {
	size, depth int
}

func parse(s string) []layer {
	lines := strings.Split(s, "\n")
	list := make([]layer, len(lines))
	for i, line := range lines {
		linesSplitted := strings.Split(line, ": ")
		list[i].depth, _ = strconv.Atoi(linesSplitted[0])
		list[i].size, _ = strconv.Atoi(linesSplitted[1])
	}
	return list
}
