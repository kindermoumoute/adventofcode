package main

import (
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
	"github.com/kindermoumoute/adventofcode/pkg/twod"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {

	part1 := len(moveRope(parse(input), 2))
	part2 := len(moveRope(parse(input), 10))

	return part1, part2
}

func moveRope(moves []Move, knots int) twod.Map {
	H := twod.NewPoint(0, 0)
	tailMap := make(twod.Map)
	rope := make(map[int]twod.Vector)
	for i := 0; i < knots; i++ {
		rope[i] = 0
	}
	ropeIndex := 0
	for _, move := range moves {
		H.Speed = move.Dir
		for i := 0; i < move.Steps; i++ {
			H.Move(1)
			ropeIndex++
			for currentKnotIndex := 0; currentKnotIndex < knots; currentKnotIndex++ {
				if currentKnotIndex != 0 {
					headingKnotIndex := currentKnotIndex - 1
					diff := rope[headingKnotIndex] - rope[currentKnotIndex]
					diffX := diff.X()
					diffY := diff.Y()
					toUpdate := false
					if pkg.Abs(diff.X()) == 2 {
						diffX /= 2
						toUpdate = true
					}
					if pkg.Abs(diff.Y()) == 2 {
						diffY /= 2
						toUpdate = true
					}
					if toUpdate {
						rope[currentKnotIndex] = twod.NewVector(rope[currentKnotIndex].X()+diffX, rope[currentKnotIndex].Y()+diffY)
					}
				} else {
					rope[currentKnotIndex] = H.Pos
				}
			}

			tailMap[rope[knots-1]] = "#"
		}
	}
	return tailMap
}

type Move struct {
	Dir   twod.Vector
	Steps int
}

var directions = map[string]twod.Vector{
	"R": twod.RIGHT,
	"U": twod.UP,
	"L": twod.LEFT,
	"D": twod.DOWN,
}

func parse(s string) []Move {
	moves := []Move(nil)
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		m := Move{}
		str := ""
		pkg.MustScanf(line, "%s %d", &str, &m.Steps)
		m.Dir = directions[str]
		moves = append(moves, m)
	}
	return moves
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
