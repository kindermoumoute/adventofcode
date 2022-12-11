package main

import (
	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
	"github.com/kindermoumoute/adventofcode/pkg/twod"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {

	part1, part2 := 0, 0

	m := twod.NewMapFromInput(input)
	topLeft := m.FindTopLeft()
	for vector, v := range m {
		m[vector] = pkg.MustAtoi(string(v.(int32)))
	}

	// Part 1
	visibles := make(twod.Map)
	p := twod.NewPoint(m.FindTopLeft(), twod.DOWN)
	for {
		p.TurnLeft()
		movesInRow := 0
		treeHeight := -1
		for {
			currentPoint, exist := m[p.Pos]
			if !exist {
				break
			}
			nextTreeHeight := currentPoint.(int)
			if nextTreeHeight > treeHeight {
				visibles[p.Pos] = m[p.Pos]
				treeHeight = nextTreeHeight
			}
			p.Move(1)
			movesInRow++
		}

		p.Move(-movesInRow)
		p.TurnRight()

		p.Move(1)
		_, exist := m[p.Pos]
		if !exist {
			p.Move(-1)
			p.TurnLeft()
			if p.Pos == topLeft && p.Speed == twod.DOWN {
				break
			}
		}
	}
	part1 = len(visibles)

	// Part 2
	scenicScores := make(twod.Map)
	for pos, treeHeight := range m {
		dirScores := make(map[twod.Vector]int)
		for _, dir := range twod.FourDirections {
			p := twod.NewPoint(pos, dir)
			for {
				p.Move(1)
				height, exist := m[p.Pos]
				if !exist {
					break
				}
				dirScores[dir]++
				if treeHeight.(int) <= height.(int) {
					break
				}
			}
		}
		sum := 1
		if len(dirScores) != 4 {
			sum = 0
		}
		for _, s := range dirScores {
			sum *= s
		}
		scenicScores[pos] = sum
	}

	for _, score := range scenicScores {
		if score.(int) > part2 {
			part2 = score.(int)
		}
	}

	return part1, part2
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
