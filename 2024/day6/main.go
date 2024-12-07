package main

import (
	"github.com/kindermoumoute/adventofcode/pkg/execute"
	"github.com/kindermoumoute/adventofcode/pkg/twod"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	part2 := 0

	m := twod.NewMapFromInput(input)

	guardInitPos := m.Find('^')[0]
	walls := m.Filter('#')
	guard := twod.NewPoint(guardInitPos, twod.UP)
	distinctPos, _ := walk(guard, walls, m)
	part1 := len(distinctPos)

	for pos := range distinctPos {
		tmpWall := walls.Clone()
		tmpWall[pos] = '#'
		_, infiniteLoop := walk(twod.NewPoint(guardInitPos, twod.UP), tmpWall, m)
		if infiniteLoop {
			part2++
		}
	}
	return part1, part2
}

func walk(guard *twod.P, walls, m twod.Map) (map[twod.Vector]bool, bool) {
	distinctPos := make(map[twod.Vector]bool)
	distinctP := make(map[[2]twod.Vector]bool)
	distinctPos[guard.Pos] = true
	distinctP[[2]twod.Vector{guard.Pos, guard.Speed}] = true
	for {
		guard.Move(1)
		if _, exist := walls[guard.Pos]; exist {
			guard.Move(-1)
			guard.TurnRight()
		} else {
			p := [2]twod.Vector{guard.Pos, guard.Speed}
			if _, exist := distinctP[p]; exist {
				return distinctPos, true
			}

			distinctPos[guard.Pos] = true
			distinctP[p] = true
		}

		if _, exist := m[guard.Pos]; !exist {
			delete(distinctPos, guard.Pos)
			return distinctPos, false
		}
	}
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
