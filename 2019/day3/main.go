package main

import (
	"strconv"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
)

// returns part1 and part2
func run(input string) (string, string) {
	list1, list2 := parse(input)

	crossMap := map[pkg.P]int{}
	stepsMap := []map[pkg.P]int(nil)
	for n, currentList := range [][]string{list1, list2} {
		p := pkg.NewPoint()
		stepsMap = append(stepsMap, map[pkg.P]int{})
		for _, move := range currentList {
			steps, err := strconv.Atoi(move[1:])
			pkg.Check(err)
			dir := map[uint8]int{
				'R': pkg.RIGHT,
				'D': pkg.DOWN,
				'U': pkg.UP,
				'L': pkg.LEFT,
			}[move[0]]
			markPoints(p, dir, steps, n, crossMap, stepsMap[n])
		}
	}

	// closest cross point from origin
	part1 := 9999999999999999
	crossedPoints := []pkg.P(nil)
	for p, v := range crossMap {
		if v == 3 {
			crossedPoints = append(crossedPoints, p)
			newMin := p.DistFromOrigin()
			if newMin < part1 {
				part1 = newMin
			}
		}
	}

	// fewest steps to cross point
	part2 := 9999999999999999
	for _, p := range crossedPoints {
		score := stepsMap[0][p] + stepsMap[1][p]

		if score < part2 {
			part2 = score
		}
	}

	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func markPoints(p *pkg.P, direction, steps, n int, crossMap, stepsMap map[pkg.P]int) {
	p.CurrentDirection = direction

	for i := 0; i < steps; i++ {
		p.Move(1)
		newP := pkg.P{X: p.X, Y: p.Y}
		crossMap[newP] |= 1 << n
		_, exist := stepsMap[newP]
		if !exist {
			stepsMap[newP] = p.Steps
		}
	}
}

func parse(s string) ([]string, []string) {
	lines := strings.Split(s, "\n")
	return strings.Split(lines[0], ","), strings.Split(lines[1], ",")
}

func main() {
	pkg.Execute(run, tests, puzzle, true)
}
