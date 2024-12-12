package main

import (
	"github.com/kindermoumoute/adventofcode/pkg/execute"
	"github.com/kindermoumoute/adventofcode/pkg/twod"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	part1, part2 := 0, 0

	m := twod.NewMapFromInput(input)
	zones := []twod.Map(nil)
	for vector, v := range m {
		zone := make(twod.Map)
		findZone(m, zone, vector, v)
		zones = append(zones, zone)
		for vec := range zone {
			delete(m, vec)
		}
	}

	for _, zone := range zones {
		perim := 0
		sides := map[twod.Vector]map[twod.Vector]struct{}{}
		for vector := range zone {
			for _, direction := range twod.FourDirections {
				_, exist := zone[vector+direction]
				if !exist {
					if _, exist = sides[direction]; !exist {
						sides[direction] = map[twod.Vector]struct{}{}
					}
					sides[direction][vector] = struct{}{}
					perim++
				}
			}
		}

		sideCount := 0
		for dir, vectors := range sides {
			seen := map[twod.Vector]struct{}{}
			searchDir := dir.RotateDegree(90)
			for vector := range vectors {
				if _, exist := seen[vector]; exist {
					continue
				}
				sideCount++
				for _, d := range []twod.Vector{searchDir, -searchDir} {
					p := vector
					for _, exist := vectors[p]; exist; _, exist = vectors[p] {
						seen[p] = struct{}{}
						p += d
					}
				}
			}

		}
		part1 += perim * len(zone)
		part2 += sideCount * len(zone)
	}

	return part1, part2
}

func findZone(m, zone twod.Map, vector twod.Vector, zoneValue any) {
	v, exist := m[vector]
	if !exist || v != zoneValue {
		return
	}
	if _, exist = zone[vector]; exist {
		return
	}

	zone[vector] = zoneValue
	for _, direction := range twod.FourDirections {
		findZone(m, zone, vector+direction, zoneValue)
	}
	return
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
