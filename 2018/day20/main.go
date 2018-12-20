package main

import (
	"fmt"
	"strconv"

	"github.com/beefsack/go-astar"

	"github.com/kindermoumoute/adventofcode/pkg"
)

// returns part1 and part2
func run(input string) (string, string) {
	path := []byte(input)
	stack := make([]pkg.P, 1, 10)
	i := 0
	maxX, minX := 0, 0
	maxY, minY := 0, 0
	pathMap := make(map[pkg.P]*pkg.Tile)
	pathMap[stack[0]] = &pkg.Tile{P: stack[0], EmptyPoints: pathMap}
	for _, dir := range path[1 : len(path)-1] {
		switch dir {
		case 'N':
			stack[i].MoveUp(1)
		case 'W':
			stack[i].MoveLeft(1)
		case 'E':
			stack[i].MoveRight(1)
		case 'S':
			stack[i].MoveDown(1)
		case '(':
			stack = append(stack, stack[i])
			i++
			continue
		case ')':
			stack = stack[:len(stack)-1]
			i--
			continue
		case '|':
			stack[i] = stack[i-1]
			continue
		}
		dir := stack[i].CurrentDirection
		stack[i].CurrentDirection = 0
		pathMap[stack[i]] = &pkg.Tile{P: stack[i], EmptyPoints: pathMap}
		stack[i].CurrentDirection = dir
		stack[i].Move(1)
		stack[i].CurrentDirection = 0
		pathMap[stack[i]] = &pkg.Tile{P: stack[i], EmptyPoints: pathMap}

		if minX > stack[i].X {
			minX = stack[i].X
		}
		if maxX < stack[i].X {
			maxX = stack[i].X
		}
		if minY > stack[i].Y {
			minY = stack[i].Y
		}
		if maxY < stack[i].Y {
			maxY = stack[i].Y
		}
	}

	str := ""
	orig := pathMap[pkg.P{X: 0, Y: 0}]
	maxDist := 0.0
	part2 := 0
	for j := minY - 1; j <= maxY+1; j++ {
		fmt.Println(j, "/", maxY+1)
		for i := minX - 1; i <= maxX+1; i++ {
			p := pkg.P{X: i, Y: j}
			tile, exist := pathMap[p]
			if !exist {
				str += "#"
			} else {

				_, dist, found := astar.Path(orig, tile)
				if dist >= 1999 {
					part2++
				}
				if found && dist > maxDist {
					maxDist = dist
				}
				str += "."
			}
		}
		str += "\n"
	}
	fmt.Println(str)
	return strconv.Itoa(int(maxDist) / 2), strconv.Itoa(part2 / 2)
}

func main() {
	pkg.Execute(run, tests, puzzle, true)
}
