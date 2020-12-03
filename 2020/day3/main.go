package main

import (
	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
	"github.com/kindermoumoute/adventofcode/pkg/twod"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	m := twod.NewMapFromInput(input)
	counts := []int(nil)
	for i, slopes := range []twod.Vector{1 - 1i, 3 - 1i, 5 - 1i, 7 - 1i, 1 - 2i} {
		p := twod.NewPoint(m.FindTopLeft(), slopes)
		counts = append(counts, 0)
		width := m.Width()
		for {
			char, exist := m[p.Pos]
			if !exist {
				if p.Pos.Y() >= 0 {
					p.Pos = twod.NewVector(p.Pos.X()%width, p.Pos.Y())
					continue
				}
				break
			}
			if char == '#' {
				counts[i]++
			}
			p.Move(1)
		}
	}

	return counts[1], pkg.Multiply(counts...)
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
