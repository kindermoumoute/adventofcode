package main

import (
	"github.com/kindermoumoute/adventofcode/pkg/execute"
	"github.com/kindermoumoute/adventofcode/pkg/twod"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	dim3 := make(map[int]twod.Map)
	dim3[0] = twod.NewMapFromInput(input)

	part1, part2 := 0, 0
	for i := 0; i < 6; i++ {
		diff := make(map[int]twod.Map)
		dim3[i+1] = make(twod.Map)
		dim3[-i-1] = make(twod.Map)
		for v := range dim3[0] {
			dim3[i+1][v] = '.'
			dim3[-i-1][v] = '.'
		}
		for z, m := range dim3 {
			topLeft := m.FindTopLeft()
			bottomRight := m.FindBottomRight()
			for x := topLeft.X() - 1; x <= bottomRight.X()+1; x++ {
				m[twod.NewVector(x, topLeft.Y()+1)] = '.'
				m[twod.NewVector(x, bottomRight.Y()-1)] = '.'
			}
			for y := bottomRight.Y(); y <= topLeft.Y(); y++ {
				m[twod.NewVector(topLeft.X()-1, y)] = '.'
				m[twod.NewVector(bottomRight.X()+1, y)] = '.'
			}
			diff[z] = make(twod.Map)
			for v := range m {
				activeCount := 0
				for j := -1; j <= 1; j++ {
					dirs := twod.AllDirections
					if j != 0 {
						dirs = append(dirs, 0)
					}
					mz, exist := dim3[z+j]
					if !exist {
						continue
					}
					for _, dir := range dirs {
						if mz[v+dir] == '#' {
							activeCount++
						}
					}
				}
				switch {
				case m[v] == '#' && (activeCount == 2 || activeCount == 3) ||
					m[v] == '.' && activeCount == 3:
					diff[z][v] = '#'
				default:
					diff[z][v] = '.'
				}
			}
		}

		for z, m := range diff {
			for v, state := range m {
				dim3[z][v] = state
			}
		}
	}

	for _, m := range dim3 {
		for _, state := range m {
			if state == '#' {
				part1++
			}
		}
	}

	// part 2
	dim4 := make(twod.Map)
	dim4[0] = twod.NewMapFromInput(input)

	for i := 0; i < 6; i++ {
		diff := make(twod.Map)

		topLeft := dim4.FindTopLeft()
		bottomRight := dim4.FindBottomRight()
		initMap := make(twod.Map)
		for v := range dim4[0].(twod.Map) {
			initMap[v] = '.'
		}

		for x := topLeft.X() - 1; x <= bottomRight.X()+1; x++ {
			dim4[twod.NewVector(x, topLeft.Y()+1)] = initMap.Clone()
			dim4[twod.NewVector(x, bottomRight.Y()-1)] = initMap.Clone()
		}
		for y := bottomRight.Y(); y <= topLeft.Y(); y++ {
			dim4[twod.NewVector(topLeft.X()-1, y)] = initMap.Clone()
			dim4[twod.NewVector(bottomRight.X()+1, y)] = initMap.Clone()
		}
		for z, mI := range dim4 {
			m := mI.(twod.Map)
			topLeft := m.FindTopLeft()
			bottomRight := m.FindBottomRight()
			for x := topLeft.X() - 1; x <= bottomRight.X()+1; x++ {
				m[twod.NewVector(x, topLeft.Y()+1)] = '.'
				m[twod.NewVector(x, bottomRight.Y()-1)] = '.'
				m[twod.NewVector(x, bottomRight.Y()-2)] = '.'
			}
			for y := bottomRight.Y(); y <= topLeft.Y(); y++ {
				m[twod.NewVector(topLeft.X()-1, y)] = '.'
				m[twod.NewVector(bottomRight.X()+1, y)] = '.'
			}
			diff[z] = make(twod.Map)

			for v := range m {
				activeCount := 0
				tmpCount := 0
				for _, v2dir := range append(twod.AllDirections, 0) {
					dirs := twod.AllDirections
					if v2dir != 0 {
						dirs = append(dirs, 0)
					}
					mz, exist := dim4[z+v2dir]
					if !exist {
						continue
					}
					for _, dir := range dirs {
						tmpCount++
						if mz.(twod.Map)[v+dir] == '#' {
							activeCount++
						}
					}
				}
				switch {
				case m[v] == '#' && (activeCount == 2 || activeCount == 3) ||
					m[v] == '.' && activeCount == 3:
					diff[z].(twod.Map)[v] = '#'
				default:
					diff[z].(twod.Map)[v] = '.'
				}
			}
		}

		for z, m := range diff {
			for v, state := range m.(twod.Map) {
				dim4[z].(twod.Map)[v] = state
			}
		}
	}

	for _, m := range dim4 {
		for _, state := range m.(twod.Map) {
			if state == '#' {
				part2++
			}
		}
	}

	return part1, part2
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
