package main

import (
	"github.com/kindermoumoute/adventofcode/pkg/execute"
	"github.com/kindermoumoute/adventofcode/pkg/twod"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	part1 := cubeway(input, []twod.Vector{twod.LEFT, 0, twod.RIGHT})
	part2 := cubeway(input, append(twod.AllDirections, 0))

	return part1, part2
}

func cubeway(input string, zwDirs []twod.Vector) int {
	zwMap := make(map[twod.Vector]twod.Map)
	zwMap[0] = twod.NewMapFromInput(input).Filter('#')

	for i := 0; i < 6; i++ {
		activeCounts := make(map[twod.Vector]map[twod.Vector]int)
		for zw, xyMap := range zwMap {
			for xy := range xyMap {
				for _, zwDir := range zwDirs {
					xyDirs := twod.AllDirections
					if zwDir != 0 {
						xyDirs = append(xyDirs, 0)
					}
					_, exist := activeCounts[zw+zwDir]
					if !exist {
						activeCounts[zw+zwDir] = make(map[twod.Vector]int)
					}
					for _, xyDir := range xyDirs {
						activeCounts[zw+zwDir][xy+xyDir]++
					}
				}
			}
		}
		nextzwMap := make(map[twod.Vector]twod.Map)
		for z, m := range activeCounts {
			_, exist := nextzwMap[z]
			if !exist {
				nextzwMap[z] = make(twod.Map)
			}
			for v := range m {
				count := activeCounts[z][v]
				state, exist := zwMap[z][v]
				if !exist {
					state = '.'
				}
				if state == '#' && (count == 2 || count == 3) ||
					state == '.' && count == 3 {
					nextzwMap[z][v] = '#'
				}
			}
		}
		zwMap = nextzwMap
	}

	sum := 0
	for _, m := range zwMap {
		sum += len(m)
	}
	return sum
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
