package main

import (
	"github.com/kindermoumoute/adventofcode/pkg/execute"
	"github.com/kindermoumoute/adventofcode/pkg/twod"
	"github.com/samber/lo"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	part1, part2 := 0, 0

	multiplier := 999999
	if input != puzzle {
		multiplier = 9
	}

	m := twod.NewMapFromInput(input)
	points := m.Filter('#')
	topRight := m.FindTopRight()

	// Find columns and rows to add
	rowToAdd := []int(nil)
	maxX, maxY := topRight.X(), topRight.Y()
nextRow:
	for j := 0; j <= maxX; j++ {
		for i := 0; i <= maxY; i++ {
			if points[twod.NewVector(i, j)] == '#' {
				continue nextRow
			}
		}
		rowToAdd = append(rowToAdd, j)
	}
	columnToAdd := []int(nil)
nextColumn:
	for i := 0; i <= maxX; i++ {
		for j := 0; j <= maxY; j++ {
			if points[twod.NewVector(i, j)] == '#' {
				continue nextColumn
			}
		}
		columnToAdd = append(columnToAdd, i)
	}

	// Expand universe
	part2Points := points.Clone()
	for i, emptyColumn := range columnToAdd {
		points = lo.MapKeys(points, shiftColumn(emptyColumn+i, 1))
		part2Points = lo.MapKeys(part2Points, shiftColumn(emptyColumn+i*multiplier, multiplier))
	}
	for i, emptyRow := range rowToAdd {
		points = lo.MapKeys(points, shiftRow(emptyRow+i, 1))
		part2Points = lo.MapKeys(part2Points, shiftRow(emptyRow+i*multiplier, multiplier))
	}

	// Count distance between each unique pair
	part1 = countDists(points)
	part2 = countDists(part2Points)

	return part1, part2
}

func shift(f func(x, y int) (int, int)) func(_ interface{}, point twod.Vector) twod.Vector {
	return func(_ interface{}, point twod.Vector) twod.Vector {
		return twod.NewVector(f(point.X(), point.Y()))
	}
}

func shiftColumn(colIdx, nbCol int) func(_ interface{}, point twod.Vector) twod.Vector {
	return shift(func(x, y int) (int, int) {
		if x >= colIdx {
			x += nbCol
		}
		return x, y
	})
}

func shiftRow(rowIdx, nbRow int) func(_ interface{}, point twod.Vector) twod.Vector {
	return shift(func(x, y int) (int, int) {
		if y >= rowIdx {
			y += nbRow
		}
		return x, y
	})
}

func countDists(points twod.Map) int {
	d := 0
	r := make(map[[2]twod.Vector]int)
	for a := range points {
		for b := range points {
			_, exist := r[[2]twod.Vector{b, a}]
			if exist || a == b {
				continue
			}
			r[[2]twod.Vector{a, b}] = 0
		}
	}

	for vectors := range r {
		d += vectors[0].DistFrom(vectors[1])
	}
	return d
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
