package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
)

// returns part1 and part2
func run(input string) (string, string) {
	part2Limit := 10000
	if input != puzzle {
		part2Limit = 32
	}
	list := parse(input)

	sizes := [][]int{
		//TODO: switch to a static retry number
		{-800, 900, -900, 900},
		{-901, 901, -901, 901},
	}
	iMin, iMax, jMin, jMax := 0, 0, 0, 0
	surfaces := make([]map[pkg.P]int, len(sizes))
	part2 := 0
	part2Found := false
	for n, size := range sizes {
		iMin, iMax, jMin, jMax = size[0], size[1], size[2], size[3]
		surfaces[n] = make(map[pkg.P]int)
		surfacesMap := surfaces[n]
		for i := iMin; i < iMax; i++ {
			for j := jMin; j < jMax; j++ {
				point := pkg.P{X: i, Y: j}
				min := 999999999999999
				pID := 0
				ignore := false
				for p := range list {
					currentDist := point.DistFrom(&list[p])
					if currentDist < min {
						min = currentDist
						pID = p
						ignore = false
					} else if currentDist == min {
						ignore = true
					}
				}
				if !ignore {
					surfacesMap[list[pID]]++
				}

				// part2
				if !part2Found {
					sum := 0
					for p := range list {
						sum += point.DistFrom(&list[p])
					}
					if sum < part2Limit {
						part2++
					}
				}
			}
		}
		part2Found = true
	}

	surfacesLists := make([][]int, len(surfaces))
	for n, surfacesMap := range surfaces {
		surfacesLists[n] = make([]int, 0, len(list))
		for _, surf := range surfacesMap {
			surfacesLists[n] = append(surfacesLists[n], surf)
		}
		sort.Ints(surfacesLists[n])
	}

	part1 := 0
	for i := range surfacesLists[0] {
		diff := false
		surface := surfacesLists[0][i]
		for n := range surfacesLists {
			if surface != surfacesLists[n][i] {
				diff = true
				break
			}
		}
		if diff {
			break
		}
		part1 = surface
	}

	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func parse(s string) []pkg.P {
	lines := strings.Split(s, "\n")
	points := make([]pkg.P, len(lines))
	for i, line := range lines {
		n, err := fmt.Sscanf(line, "%d, %d", &points[i].X, &points[i].Y)
		pkg.Check(err)
		if n != 2 {
			panic(fmt.Errorf("%d args expected in scanf, got %d", 2, n))
		}
	}
	return points
}

func main() {
	pkg.Execute(run, tests, puzzle, true)
}
