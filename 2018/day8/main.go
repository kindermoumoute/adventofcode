package main

import (
	"strconv"

	"github.com/kindermoumoute/adventofcode/pkg"
)

type Node struct {
	metadata []int
	children []*Node
}

func (n *Node) Part1() int {
	if n == nil {
		return 0
	}
	count := pkg.Sum(n.metadata...)
	for _, c := range n.children {
		count += c.Part1()
	}
	return count
}

func (n *Node) Part2() int {
	if n == nil {
		return 0
	}
	if len(n.children) == 0 {
		return pkg.Sum(n.metadata...)
	}
	count := 0
	for _, m := range n.metadata {
		if m != 0 && m <= len(n.children) {
			count += n.children[m-1].Part2()
		}
	}
	return count
}

func buildNode(b []int) (*Node, []int) {
	if len(b) == 0 {
		return nil, b
	}
	childrenCount := b[0]
	metadataCount := b[1]
	b = b[2:]
	n := &Node{
		metadata: make([]int, metadataCount),
		children: make([]*Node, childrenCount),
	}
	for i := range n.children {
		n.children[i], b = buildNode(b)
	}
	for i := range n.metadata {
		n.metadata[i] = b[i]
	}
	return n, b[metadataCount:]
}

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	intList := pkg.ParseIntList(input, " ")
	root, _ := buildNode(intList)
	return strconv.Itoa(root.Part1()), strconv.Itoa(root.Part2())
}
func main() {
	pkg.Execute(run, tests, puzzle, true)
}
