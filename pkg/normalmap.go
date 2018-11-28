package pkg

import "math"

const (
	NORTH = iota
	EAST
	SOUTH
	WEST
)

var DirectionMap = map[int]struct {
	x, y int
	name string
}{
	NORTH: {0, 1, "north"},
	EAST:  {1, 0, "east"},
	SOUTH: {0, -1, "south"},
	WEST:  {-1, 0, "west"},
}

type P struct {
	X, Y             int
	CurrentDirection int
}

func NewPoint() *P {
	return &P{}
}

func (p *P) Move(steps int) {
	p.X += DirectionMap[p.CurrentDirection].x * steps
	p.Y += DirectionMap[p.CurrentDirection].y * steps
}

func (p *P) Left() {
	p.CurrentDirection = (p.CurrentDirection - 1 + 4) % 4
}

func (p *P) Right() {
	p.CurrentDirection = (p.CurrentDirection + 1) % 4
}

func (p *P) DistFromOrigin() int {
	return p.DistFrom(&P{X: 0, Y: 0})
}

func (p *P) DistFrom(from *P) int {
	return Abs(p.X-from.X) + Abs(p.Y-from.Y)
}

func (p *P) EuclideanDistFrom(from *P) float64 {
	return math.Sqrt(float64(p.DistFrom(from)))
}

func (p *P) EuclideanDistFromOrigin() float64 {
	return math.Sqrt(float64(p.DistFromOrigin()))
}
