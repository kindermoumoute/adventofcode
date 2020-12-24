package main

import (
	"regexp"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg/execute"
	"github.com/kindermoumoute/adventofcode/pkg/twod"
)

var hexaDirs = map[string]twod.Vector{
	"e":  twod.NewVector(1, 0),
	"se": twod.NewVector(0, -1),
	"sw": twod.NewVector(-1, -1),
	"w":  twod.NewVector(-1, 0),
	"nw": twod.NewVector(0, 1),
	"ne": twod.NewVector(1, 1),
}

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	hexaPositions := parse(input)

	// part 1
	m := twod.Map{}
	for _, p := range hexaPositions {
		flip(m, p)
	}
	part1 := len(m)

	// part 2
	for i := 0; i < 100; i++ {
		blackCounts := make(twod.Map)
		for pos := range m {
			for _, dir := range hexaDirs {
				_, exist := blackCounts[pos+dir]
				if !exist {
					blackCounts[pos+dir] = 0
				}
				blackCounts[pos+dir] = blackCounts[pos+dir].(int) + 1
			}
		}
		newM := make(twod.Map)
		for pos := range m {
			count, exist := blackCounts[pos]
			if !exist || count.(int) > 2 {
				continue
			}
			newM[pos] = '#'
		}
		for pos, count := range blackCounts {
			_, exist := m[pos]
			if !exist && count == 2 {
				newM[pos] = '#'
			}
		}
		m = newM
	}
	part2 := len(m)

	return part1, part2
}

func flip(m twod.Map, pos twod.Vector) {
	_, exist := m[pos]
	if !exist {
		m[pos] = '#'
		return
	}
	delete(m, pos)
}

func parse(s string) []twod.Vector {
	points := []twod.Vector(nil)
	reg := regexp.MustCompile(`e|se|sw|w|nw|ne`)
	for _, line := range strings.Split(s, "\n") {
		p := twod.NewPoint(0, 0)
		for _, dir := range reg.FindAllString(line, -1) {
			p.Speed = hexaDirs[dir]
			p.Move(1)
		}
		points = append(points, p.Pos)
	}
	return points
}

func main() {
	execute.Run(run, tests, puzzle, false)
}
