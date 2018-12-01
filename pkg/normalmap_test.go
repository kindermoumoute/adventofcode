package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalMap(t *testing.T) {
	assert := assert.New(t)
	var testMap = [][]string{
		{"A", "B", "C", "D", "E"},
		{"F", "G", "H", "I", "J"},
		{"K", "L", "M", "N", "O"},
		{"P", "Q", "R", "S", "T"},
		{"U", "V", "W", "X", "Y"},
	}
	originTestMap := P{X: 2, Y: 2}

	p := NewPoint()
	p.Move(1)
	assert.Equal(0, p.X)
	assert.Equal(1, p.Y)
	p.MoveRight(2)
	assert.Equal(2, p.X)
	assert.Equal(1, p.Y)
	p.Move(-4)
	assert.Equal(-2, p.X)
	assert.Equal(1, p.Y)
	p.TurnRight()
	p.Move(2)
	cell := p.FindInMap(testMap, originTestMap)
	assert.Equal("P", cell)
}
