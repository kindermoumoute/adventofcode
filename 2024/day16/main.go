package main

import (
	"github.com/kindermoumoute/adventofcode/pkg/execute"
	"github.com/kindermoumoute/adventofcode/pkg/twod"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/path"
	"gonum.org/v1/gonum/graph/simple"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	part1, part2 := 0, 0

	m := twod.NewMapFromInput(input)

	g := simple.NewWeightedDirectedGraph(0, 0)
	var startNode graph.Node
	endNode := g.NewNode()
	g.AddNode(endNode)
	nodes := make(map[[2]twod.Vector]graph.Node)
	posFromNode := make(map[int64]twod.Vector)
	for vector, val := range m.FilterOut('#') {
		for _, dir := range twod.FourDirections {
			n := g.NewNode()
			nodes[[2]twod.Vector{vector, dir}] = n
			posFromNode[n.ID()] = vector
			g.AddNode(n)

			if val == 'S' && dir == twod.RIGHT {
				startNode = n
			}
			if val == 'E' {
				g.SetWeightedEdge(g.NewWeightedEdge(n, endNode, 0))
			}
		}
	}
	for vector, fromNode := range nodes {
		if toNode, exist := nodes[[2]twod.Vector{vector[0] + vector[1], vector[1]}]; exist {
			g.SetWeightedEdge(g.NewWeightedEdge(fromNode, toNode, 1))
		}

		g.SetWeightedEdge(g.NewWeightedEdge(fromNode, nodes[[2]twod.Vector{vector[0], vector[1].RotateDegree(90)}], 1000))
		g.SetWeightedEdge(g.NewWeightedEdge(fromNode, nodes[[2]twod.Vector{vector[0], vector[1].RotateDegree(-90)}], 1000))
	}

	paths := path.YenKShortestPaths(g, -1, 0, startNode, endNode)
	uniq := make(twod.Map)
	for i, p := range paths {
		if i == 0 {
			part1 = int(computePathCost(g, p))
		}
		for _, n := range p[:len(p)-1] {
			uniq[posFromNode[n.ID()]] = 'O'
		}
	}
	part2 = len(uniq)

	return part1, part2
}

func computePathCost(g *simple.WeightedDirectedGraph, path []graph.Node) float64 {
	totalCost := 0.0

	// Iterate through the path nodes and sum edge weights
	for i := 0; i < len(path)-1; i++ {
		edge := g.WeightedEdge(path[i].ID(), path[i+1].ID())
		if edge == nil {
			panic("not possible")
		}
		totalCost += edge.Weight()
	}

	return totalCost
}
func main() {
	execute.Run(run, tests, puzzle, true)
}
