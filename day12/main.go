package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	list := parse(puzzle)
	groupMap := make(map[int]bool)
	part1 := 0
	for i := range list {
		min := list[i].minInGroup(make(map[int]bool), 9999999)
		groupMap[min] = true
		if part1 == 0 && min == 0 {
			part1 = list[i].countSiblings(make(map[int]bool))
		}
	}
	fmt.Println(part1, len(groupMap))
}

type node struct {
	value int
	s     []*node

	tmp string
}

func (n *node) minInGroup(ref map[int]bool, min int) int {
	if n == nil {
		return min
	}
	if _, exist := ref[n.value]; exist {
		return min
	}
	if n.value < min {
		min = n.value
	}

	ref[n.value] = true
	for _, s := range n.s {
		min = s.minInGroup(ref, min)
	}
	return min
}

func (n *node) countSiblings(ref map[int]bool) int {
	if n == nil {
		return 0
	}
	if _, exist := ref[n.value]; exist {
		return 0
	}
	count := 0

	ref[n.value] = true
	for _, s := range n.s {
		count += s.countSiblings(ref)
	}
	return count + 1
}

func (n *node) String() string {
	s := "value: " + strconv.Itoa(n.value)
	return s
}
func parse(s string) []*node {
	splittedPuzzle := strings.Split(s, "\n")
	v := make([]*node, len(splittedPuzzle))
	for i, elt := range splittedPuzzle {
		splitLine := strings.Split(elt, " <-> ")
		v[i] = &node{value: i, s: []*node{}, tmp: splitLine[1]}
	}

	for i, e := range v {
		prgmList := strings.Split(e.tmp, ", ")
		for _, prgm := range prgmList {
			prgmInt, _ := strconv.Atoi(prgm)
			v[i].s = append(v[i].s, v[prgmInt])
		}
	}
	return v
}
