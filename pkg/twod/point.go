package twod

import (
	"math"
)

const (
	UP    = Vector(1i)
	LEFT  = Vector(-1)
	DOWN  = Vector(-1i)
	RIGHT = Vector(1)
)

var DirectionNames = map[Vector]string{
	UP:    "up",
	LEFT:  "left",
	DOWN:  "down",
	RIGHT: "right",
}

type P struct {
	Pos   Vector
	Speed Vector
	Steps int
}

func NewPoint(position Vector, direction Vector) *P {
	p := &P{
		Pos:   position,
		Speed: direction,
	}
	return p
}

func (p *P) MoveLeft(steps int) {
	p.Speed = LEFT
	p.Move(steps)
}

func (p *P) MoveRight(steps int) {
	p.Speed = RIGHT
	p.Move(steps)
}

func (p *P) MoveUp(steps int) {
	p.Speed = UP
	p.Move(steps)
}

func (p *P) MoveDown(steps int) {
	p.Speed = DOWN
	p.Move(steps)
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
