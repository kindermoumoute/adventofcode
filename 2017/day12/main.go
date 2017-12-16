package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	list := parse(puzzle)
	groupMap := make(map[int]int)
	for i := range list {
		nb, min := list[i].countSiblings(make(map[int]bool), 9999999)
		groupMap[min] = nb
	}
	fmt.Println(groupMap[0], len(groupMap))
}

type node struct {
	value int
	s     []*node
}

func (n *node) countSiblings(ref map[int]bool, min int) (int, int) {
	if n == nil {
		return 0, min
	}
	if _, exist := ref[n.value]; exist {
		return 0, min
	}
	if n.value < min {
		min = n.value
	}
	count := 0
	t := 0
	ref[n.value] = true
	for _, s := range n.s {
		t, min = s.countSiblings(ref, min)
		count += t
	}
	return count + 1, min
}

func parse(s string) []node {
	splittedPuzzle := strings.Split(s, "\n")
	v := make([]node, len(splittedPuzzle))
	for i, elt := range splittedPuzzle {
		splitLine := strings.Split(elt, " <-> ")
		v[i] = node{value: i, s: []*node{}}
		prgmList := strings.Split(splitLine[1], ", ")
		for _, prgm := range prgmList {
			prgmInt, _ := strconv.Atoi(prgm)
			v[i].s = append(v[i].s, &v[prgmInt])
		}
	}
	return v
}
