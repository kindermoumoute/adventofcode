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
