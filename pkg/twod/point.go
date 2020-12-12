package twod

import (
	"math"

	"github.com/beefsack/go-astar"
)

const (
	UP    = Vector(1i)
	LEFT  = Vector(-1)
	DOWN  = Vector(-1i)
	RIGHT = Vector(1)

	UPRIGHT   = UP + RIGHT
	UPLEFT    = UP + LEFT
	DOWNRIGHT = DOWN + RIGHT
	DOWNLEFT  = DOWN + LEFT
)

var (
	DirectionNames = map[Vector]string{
		UP:    "up",
		LEFT:  "left",
		DOWN:  "down",
		RIGHT: "right",
	}
	FourDirections = []Vector{DOWN, RIGHT, LEFT, UP}
	AllDirections  = []Vector{DOWN, RIGHT, LEFT, UP, UPRIGHT, UPLEFT, DOWNRIGHT, DOWNLEFT}
)

type P struct {
	Pos   Vector
	Speed Vector
	Steps int

	// emptyPoints and target are used for astar algorithm
	emptyPoints   Map
	exitCondition func(currentPoint *P, neighbors []astar.Pather) []astar.Pather
}

func NewPoint(position Vector, direction Vector) *P {
	p := &P{
		Pos:   position,
		Speed: direction,
	}
	return p
}

func (p *P) MoveLeft(steps int) {
	speedOrig := p.Speed
	p.Speed = LEFT
	p.Move(steps)
	p.Speed = speedOrig
}

func (p *P) MoveRight(steps int) {
	speedOrig := p.Speed
	p.Speed = RIGHT
	p.Move(steps)
	p.Speed = speedOrig
}

func (p *P) MoveUp(steps int) {
	speedOrig := p.Speed
	p.Speed = UP
	p.Move(steps)
	p.Speed = speedOrig
}

func (p *P) MoveDown(steps int) {
	speedOrig := p.Speed
	p.Speed = DOWN
	p.Move(steps)
	p.Speed = speedOrig
}

func (p *P) Move(steps int) {
	p.Pos += p.Speed.ScalarMult(steps)
	p.Steps += steps
}

func (p *P) TurnLeft() {
	p.Speed = p.Speed.Rotate(math.Pi / 2)
}

func (p *P) TurnRight() {
	p.Speed = p.Speed.Rotate(-math.Pi / 2)
}
