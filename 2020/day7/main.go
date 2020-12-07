package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	bags := parse(input)

	part1, part2 := 0, 0
	for _, b := range bags {
		if b.name == "shiny gold" {
			part2 = b.countBags() - 1
		}
		if b.contain("shiny gold") {
			part1++
		}
	}

	return part1, part2
}

type bag struct {
	name          string
	children      map[string]*bag
	childrenCount map[string]int
}

func (b *bag) contain(name string) bool {
	if b == nil {
		return false
	}
	for childName := range b.children {
		if childName == name {
			return true
		}
	}
	for _, child := range b.children {
		if child.contain(name) {
			return true
		}
	}
	return false
}

func (b *bag) countBags() int {
	currentCount := 1
	if b == nil {
		return currentCount
	}
	for k, count := range b.childrenCount {
		currentCount += count * b.children[k].countBags()
	}
	return currentCount
}

func parse(s string) []*bag {
	bags := []*bag(nil)
	index := make(map[string]*bag)
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		tmp := strings.Split(line, " bags contain ")
		b := &bag{
			name:          tmp[0],
			children:      make(map[string]*bag),
			childrenCount: make(map[string]int),
		}
		bags = append(bags, b)
		if _, exist := index[b.name]; exist {
			panic(fmt.Errorf("doublon %s", b.name))
		}
		index[b.name] = b
		for _, child := range strings.Split(strings.TrimRight(tmp[1], "."), ", ") {
			tmp3 := strings.Split(child, " bag")
			child = tmp3[0]
			tmp2 := strings.Split(child, " ")
			childName := strings.Join(tmp2[1:], " ")
			childCount, err := strconv.Atoi(tmp2[0])
			if err != nil {
				childCount = 0
			}
			b.children[childName] = nil
			b.childrenCount[childName] = childCount

		}
	}
	for _, b := range bags {
		for k := range b.children {
			b.children[k] = index[k]
		}
	}
	return bags
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
