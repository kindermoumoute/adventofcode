package main

import (
	"strconv"
	"strings"
)

func main() {
	//list := parse(puzzle)

}

func parse(s string) []int {
	lines := strings.Split(s, "\n")
	list := make([]int, len(lines))
	for i, line := range lines {
		nb, err := strconv.Atoi(line)
		check(err)
		list[i] = nb
	}
	return list
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
