package main

import (
	"math/cmplx"
	"sort"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	m := parse(input)

	// part1
	var bestAsteroidCoord complex128
	for p := range m {
		for toP := range m {
			if p == toP {
				continue
			}
			radius, angle := cmplx.Polar(toP - p)
			m[p][angle] = append(m[p][angle], radius)
		}
		if len(m[p]) > len(m[bestAsteroidCoord]) {
			bestAsteroidCoord = p
		}
	}

	// part 2 starts here
	bestAsteroid := m[bestAsteroidCoord]
	angles := []float64(nil)
	for a := range bestAsteroid {
		angles = append(angles, a)
	}
	sort.Sort(sort.Reverse(sort.Float64Slice(angles)))

	finalAngleIndex := 199 % len(angles)
	finalAngle := angles[finalAngleIndex]

	finalTargetIndex := finalAngleIndex % len(bestAsteroid[finalAngle])
	sort.Float64s(bestAsteroid[finalAngle])
	finalRadius := bestAsteroid[finalAngle][finalTargetIndex]

	finalTarget := cmplx.Rect(finalRadius, finalAngle) + bestAsteroidCoord

	part1 := len(bestAsteroid)
	part2 := int(imag(finalTarget)+0.1)*100 + int(real(finalTarget)+0.1) // don't forget to adjust due to float precision
	return part1, part2
}

func parse(s string) map[complex128]map[float64][]float64 {
	lines := strings.Split(s, "\n")
	m := make(map[complex128]map[float64][]float64)
	for i, line := range lines {
		for j, char := range line {
			if char == '#' {
				m[complex(float64(i), float64(j))] = make(map[float64][]float64)
			}
		}
	}
	return m
}

func main() {
	execute.Run(run, tests, puzzle, false)
}
