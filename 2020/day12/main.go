package main

import (
	"regexp"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
	"github.com/kindermoumoute/adventofcode/pkg/twod"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	//list := parse(input)

	p := twod.NewPoint(0, twod.LEFT)
	req := regexp.MustCompile(`([NSEWLRF]{1})(\d+)`)
	lines := req.FindAllStringSubmatch(input, -1)
	for _, matches := range lines {
		char := matches[1][0]
		value := pkg.MustAtoi(matches[2])
		switch char {
		case 'N':
			p.MoveUp(value)
		case 'S':
			p.MoveDown(value)
		case 'E':
			p.MoveLeft(value)
		case 'W':
			p.MoveRight(value)
		case 'L':
			p.Speed = p.Speed.RotateDegree(-value)
		case 'R':
			p.Speed = p.Speed.RotateDegree(value)
		case 'F':
			p.Move(value)
		}
	}
	part1 := p.Pos.DistFromOrigin()

	p = twod.NewPoint(0, -10+1i)
	for _, matches := range lines {
		char := matches[1][0]
		value := pkg.MustAtoi(matches[2])
		switch char {
		case 'N':
			p.Speed += twod.NewVector(0, value)
		case 'S':
			p.Speed += twod.NewVector(0, -value)
		case 'E':
			p.Speed += twod.NewVector(-value, 0)
		case 'W':
			p.Speed += twod.NewVector(value, 0)
		case 'L':
			p.Speed = p.Speed.RotateDegree(-value)
		case 'R':
			p.Speed = p.Speed.RotateDegree(value)
		case 'F':
			p.Move(value)
		}
	}
	part2 := p.Pos.DistFromOrigin()
	return part1, part2
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
