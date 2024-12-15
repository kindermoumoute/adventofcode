package main

import (
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
	"github.com/kindermoumoute/adventofcode/pkg/twod"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	tmp := strings.Split(input, "\n\n")
	baseMap := twod.NewMapFromInput(tmp[0])
	path := strings.Join(strings.Split(tmp[1], "\n"), "")

	// part 1
	part1 := solve(baseMap.Clone(), path)

	// part 2
	widerMap := make(twod.Map)
	for vector, v := range baseMap {
		left := twod.NewVector(vector.X()*2, vector.Y())
		right := twod.NewVector(vector.X()*2+1, vector.Y())
		switch v {
		case 'O':
			widerMap[left] = '['
			widerMap[right] = ']'
		case '#':
			widerMap[left] = '#'
			widerMap[right] = '#'
		case '.':
			widerMap[left] = '.'
			widerMap[right] = '.'
		case '@':
			widerMap[left] = '@'
			widerMap[right] = '.'
		}
	}
	part2 := solve(widerMap, path)

	return part1, part2
}

var dirs = map[int32]twod.Vector{
	'<': twod.LEFT,
	'>': twod.RIGHT,
	'v': twod.DOWN,
	'^': twod.UP,
}

func solve(m twod.Map, path string) int {
	m = m.FilterOut('.')

	sum := 0
	robotPos := m.Find('@')[0]
	for _, dir := range path {
		movedParts := findMovables(m, robotPos, dirs[dir])
		if len(movedParts) > 0 {
			robotPos = move(m, movedParts, dirs[dir])
		}
	}
	topLeft := m.FindTopLeft()
	for vector, v := range m {
		if v == '[' || v == 'O' {
			dist := topLeft - vector
			sum += pkg.Abs(dist.Y())*100 + pkg.Abs(dist.X())
		}
	}
	return sum
}

func move(m twod.Map, movedParts []twod.Vector, dir twod.Vector) twod.Vector {
	found := map[twod.Vector]bool{}
	robot := twod.Vector(0)
	for _, pos := range movedParts {
		if found[pos] {
			continue // already moved
		}
		found[pos] = true
		val := m[pos]
		delete(m, pos)
		m[pos+dir] = val
		if val == '@' {
			robot = pos + dir
		}
	}
	return robot
}

var otherSide = map[any]twod.Vector{
	']': twod.LEFT,
	'[': twod.RIGHT,
}

func findMovables(m twod.Map, pos, dir twod.Vector) []twod.Vector {
	tile, exists := m[pos+dir]
	if !exists {
		return []twod.Vector{pos}
	}
	if tile == '#' {
		return nil
	}

	movedParts := findMovables(m, pos+dir, dir)
	if len(movedParts) == 0 {
		return nil
	}

	if otherSideDir := otherSide[tile]; otherSideDir != 0 && (dir == twod.UP || dir == twod.DOWN) {
		wideParts := findMovables(m, pos+dir+otherSideDir, dir)
		if len(wideParts) == 0 {
			return nil
		}
		movedParts = append(movedParts, wideParts...)
	}
	return append(movedParts, pos)
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
