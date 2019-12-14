package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg/execute"

	"github.com/kindermoumoute/adventofcode/pkg"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	points := parse(input)
	p := &Picture{}
	p.Reset()
	part1 := ""
	part2 := 0
	prevDiffX, prevDiffY := 0, 0
	for {
		newPoints := make(movingPointsMap)
		prevDiffX, prevDiffY = p.DiffX(), p.DiffY()
		p.Reset()
		for _, point := range points {
			point.position.X += point.velocity.X
			point.position.Y += point.velocity.Y
			if point.position.X > p.botRightX {
				p.botRightX = point.position.X
			}
			if point.position.Y > p.botRightY {
				p.botRightY = point.position.Y
			}
			if point.position.X < p.topLeftX {
				p.topLeftX = point.position.X
			}
			if point.position.Y < p.topLeftY {
				p.topLeftY = point.position.Y
			}
			newPoints[pkg.P{X: point.position.X, Y: point.position.Y}] = point
		}
		part2++
		p.points = newPoints
		if prevDiffX < p.DiffX() || prevDiffY < p.DiffY() {
			return part1, strconv.Itoa(part2 - 1)
		}
		part1 = p.String()
	}
}

type movingPoint struct {
	position pkg.P
	velocity pkg.P
}

type movingPointsMap map[pkg.P]*movingPoint

type Picture struct {
	points               movingPointsMap
	topLeftX, topLeftY   int
	botRightX, botRightY int
}

func (p *Picture) Reset() {
	p.botRightX = -9999999
	p.botRightY = -9999999
	p.topLeftX = 9999999
	p.topLeftY = 9999999
}
func (p *Picture) DiffX() int {
	return pkg.Abs(p.botRightX - p.topLeftX)
}
func (p *Picture) DiffY() int {
	return pkg.Abs(p.botRightY - p.topLeftY)
}

func (p Picture) String() string {
	const sizeLimit = 200
	if sizeLimit < p.DiffX() || sizeLimit < p.DiffY() {
		return ""
	}
	pic := "\n"
	for j := p.topLeftY; j <= p.botRightY; j++ {
		for i := p.topLeftX; i <= p.botRightX; i++ {
			_, exist := p.points[pkg.P{X: i, Y: j}]
			if exist {
				pic += "#"
			} else {
				pic += "."
			}
		}
		pic += "\n"
	}
	return pic
}

func parse(s string) []*movingPoint {
	lines := strings.Split(s, "\n")
	points := make([]*movingPoint, len(lines))
	for i, line := range lines {
		var x, y, vx, vy int
		n, err := fmt.Sscanf(line, "position=<%d, %d> velocity=<%d,  %d>", &x, &y, &vx, &vy)
		pkg.Check(err)
		if n != 4 {
			panic(fmt.Errorf("%d args expected in scanf, got %d", 0, n))
		}
		points[i] = &movingPoint{
			position: pkg.P{X: x, Y: y},
			velocity: pkg.P{X: vx, Y: vy},
		}
	}
	return points
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
