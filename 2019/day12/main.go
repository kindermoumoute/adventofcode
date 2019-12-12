package main

import (
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	list := parse(input)
	memoized := [3]map[[4]Dim]MemoizedDim{ // memoization is done per axis
		make(map[[4]Dim]MemoizedDim), // TODO remove MemoizedDim type
		make(map[[4]Dim]MemoizedDim),
		make(map[[4]Dim]MemoizedDim),
	}
	byTime := [3]map[int][4]Dim{ // memoize time as well as position for part 1
		make(map[int][4]Dim), // TODO use slice
		make(map[int][4]Dim),
		make(map[int][4]Dim),
	}
	i := 0
	var memoizedTime [3]int
	for {
		memoizedCount := 0
		for dim := 0; dim < 3; dim++ {
			_, exist := memoized[dim][list[dim]]
			if exist {
				memoizedCount++
				if memoizedTime[dim] == 0 {
					memoizedTime[dim] = i
				}
				continue
			}
			previousDim := copyDim(list[dim])
			for j := range list[dim] {
				for k := range list[dim] {
					if j == k {
						continue
					}
					if list[dim][j].Pos < list[dim][k].Pos {
						list[dim][j].Vel++
					} else if list[dim][j].Pos > list[dim][k].Pos {
						list[dim][j].Vel--
					}
				}
			}

			// apply speed velocity
			for j := range list[dim] {
				list[dim][j].Pos += list[dim][j].Vel
			}
			memoized[dim][previousDim] = MemoizedDim{
				NextDim: copyDim(list[dim]),
				Time:    i,
			}
			byTime[dim][i] = copyDim(list[dim])
		}

		if memoizedCount == 3 {
			break
		}

		i++
	}

	sum := 0
	for i := 0; i < 4; i++ {
		sumDimPos := 0
		sumDimVel := 0
		for dim := 0; dim < 3; dim++ {
			sumDimPos += pkg.Abs(byTime[dim][999%len(byTime[dim])][i].Pos)
			sumDimVel += pkg.Abs(byTime[dim][999%len(byTime[dim])][i].Vel)
		}
		sum += sumDimPos * sumDimVel
	}
	return sum, getLCM(memoizedTime)
}

func getLCM(times [3]int) int {
	lcm := times[0]
	for _, time := range times[1:] {
		lcm = lcm * time / pkg.GCD(lcm, time)
	}
	return lcm
}

func copyDim(in [4]Dim) [4]Dim {
	var copy [4]Dim
	for i := range in {
		copy[i] = in[i]
	}
	return copy
}

type MemoizedDim struct {
	NextDim [4]Dim
	Time    int
}

type List [3][4]Dim

type Dim struct {
	Pos int
	Vel int
}

func parse(s string) List {
	lines := strings.Split(s, "\n")
	var dims List
	for i, line := range lines {
		pkg.MustScanf(line, "<x=%d, y=%d, z=%d>", &dims[0][i].Pos, &dims[1][i].Pos, &dims[2][i].Pos)

	}
	return dims
}

func main() {
	pkg.Execute(run, tests, puzzle, false)
}
