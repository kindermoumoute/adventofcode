package twod

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPoint(t *testing.T) {
	p := NewPoint(0, UP)
	p.Move(1)
	assert.Equal(t, 0, p.Pos.X())
	assert.Equal(t, 1, p.Pos.Y())
	p.MoveRight(2)
	assert.Equal(t, 2, p.Pos.X())
	assert.Equal(t, 1, p.Pos.Y())
	p.Speed = p.Speed.RotateDegree(270)
	p.Move(-4)
	assert.Equal(t, -2, p.Pos.X())
	assert.Equal(t, 1, p.Pos.Y())
	p.TurnRight()
	p.Move(2)
}
