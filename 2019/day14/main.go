package main

import (
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	allNodes := parse(input)
	root := allNodes["FUEL"]

	part1 := 0 // part1 is the consumed Ore to produce exactly one fuel
	part2 := 0 // part2 is the produced Fuel when consuming 1e12 ores
	target := int(1e12)
	for {
		root.Reset()
		root.produce(part2 + 1)
		requiredOre := allNodes["ORE"].Used
		if part2+1 == 1 {
			part1 = requiredOre
		}
		if requiredOre > target {
			break
		} else {
			part2 = pkg.Max(part2+1, (part2+1)*target/requiredOre)
		}
	}

	return part1, part2
}

func (n *Node) produce(need int) {
	n.Leftovers += need / n.Unit * n.Unit
	if need%n.Unit != 0 && n.Leftovers < need {
		n.Leftovers += n.Unit
	}

	newNeed := n.Leftovers / n.Unit
	n.Leftovers -= need
	n.Used += need

	for _, arc := range n.ChildrenArcs {
		arc.Node.produce(newNeed * arc.Multiplier)
	}
}

func (n *Node) Reset() {
	n.Leftovers = 0
	n.Used = 0
	for _, arc := range n.ChildrenArcs {
		arc.Node.Reset()
	}
}

type Arc struct {
	Multiplier int
	Node       *Node
	Visited    bool
}

type Node struct {
	ParentsArcs  map[string]*Arc
	ChildrenArcs map[string]*Arc
	Unit         int
	Name         string
	Used         int
	Leftovers    int
}

func parse(s string) map[string]*Node {
	lines := strings.Split(s, "\n")
	allNode := make(map[string]*Node)
	for _, line := range lines {
		tmp := strings.Split(line, " => ")
		producedAtOnce, name := scanArc(tmp[1])
		parentNode, exist := allNode[name]
		if !exist {
			parentNode = newNode()
			allNode[name] = parentNode

		}
		parentNode.Name = name
		parentNode.Unit = producedAtOnce

		rights := strings.Split(tmp[0], ", ")
		for _, right := range rights {
			multiplier, name := scanArc(right)
			childNode, exist := allNode[name]
			if !exist {
				childNode = newNode()
				allNode[name] = childNode
				childNode.Name = name
			}
			childNode.ParentsArcs[parentNode.Name] = &Arc{
				Multiplier: multiplier,
				Node:       parentNode,
			}
			parentNode.ChildrenArcs[childNode.Name] = &Arc{
				Multiplier: multiplier,
				Node:       childNode,
			}
		}
	}
	return allNode
}

func scanArc(s string) (int, string) {
	n, name := 0, ""
	pkg.MustScanf(s, "%d %s", &n, &name)
	return n, name
}

func newNode() *Node {
	return &Node{
		ParentsArcs:  make(map[string]*Arc),
		ChildrenArcs: make(map[string]*Arc),
		Unit:         1,
	}

}

func main() {
	pkg.Execute(run, tests, puzzle, true)
}
