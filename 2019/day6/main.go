package main

import (
	"strconv"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	allNodes, root := parse(input)
	allNodes[root].countOrbits(0)
	part1 := 0
	for _, n := range allNodes {
		part1 += n.orbitCount
	}

	part2, _ := allNodes["YOU"].findSAN(0)

	return strconv.Itoa(part1), strconv.Itoa(part2 - 2)
}

type Node struct {
	Name   string
	Leaf   map[string]*Node
	Parent *Node

	orbitCount int
	visited    bool
}

func NewNode(name string) *Node {
	return &Node{Name: name, Leaf: map[string]*Node{}}
}

func (n *Node) connectLeaf(leaf *Node) {
	n.Leaf[leaf.Name] = leaf
	if leaf.Parent != nil {
		panic("2 parents")
	}
	leaf.Parent = n
}

func (n *Node) root() string {
	if n.Parent == nil {
		return n.Name
	}
	return n.Parent.root()
}

func (n *Node) countOrbits(currentCount int) {
	n.orbitCount = currentCount
	for _, l := range n.Leaf {
		l.countOrbits(n.orbitCount + 1)
	}
}

func (n *Node) findSAN(currentDist int) (int, bool) {
	if n == nil || n.visited {
		return 0, false
	}
	n.visited = true
	if n.Name == "SAN" {
		return currentDist, true
	}

	min := 99999999
	dist, found := n.Parent.findSAN(currentDist + 1)
	if found {
		min = dist
	}

	for _, l := range n.Leaf {
		dist, found := l.findSAN(currentDist + 1)
		if found {
			min = pkg.Min(min, dist)
		}
	}
	return min, true
}

func parse(s string) (map[string]*Node, string) {
	allNodes := map[string]*Node{}
	firstLeaf := ""
	for i, line := range strings.Split(s, "\n") {
		tmp := strings.Split(line, ")")
		left, right := tmp[0], tmp[1]
		leftLeaf := createOrGetLeaf(allNodes, left)
		rightLeaf := createOrGetLeaf(allNodes, right)
		leftLeaf.connectLeaf(rightLeaf)

		if i == 0 {
			firstLeaf = left
		}
	}

	return allNodes, allNodes[firstLeaf].root() // first leaf is not always the root
}

func createOrGetLeaf(allNodes map[string]*Node, leafName string) *Node {
	leaf, exist := allNodes[leafName]
	if !exist {
		allNodes[leafName] = NewNode(leafName)
		leaf = allNodes[leafName]
	}
	return leaf
}

func main() {
	pkg.Execute(run, tests, puzzle, true)
}
