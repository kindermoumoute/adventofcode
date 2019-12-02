package pkg

import (
	"strconv"
	"strings"
)

func ParseIntList(s, sep string) []int {
	lines := strings.Split(s, sep)
	list := make([]int, len(lines))
	for i, line := range lines {
		nb, err := strconv.Atoi(line)
		Check(err)
		list[i] = nb
	}
	return list
}

func ParseIntMap(s, sep string) map[int]int {
	m := make(map[int]int)
	for i, line := range strings.Split(s, sep) {
		nb, err := strconv.Atoi(line)
		Check(err)
		m[i] = nb
	}
	return m
}
