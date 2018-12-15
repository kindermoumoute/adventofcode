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

//
//func TestAstar(t *testing.T) {
//	assert := assert.New(t)
//	emptyPoints := make(map[P]*PAstar)
//	points := []*PAstar{
//		{P: P{X: 2, Y: 0}},
//		{P: P{X: 2, Y: 1}},
//		{P: P{X: 1, Y: 1}},
//		{P: P{X: 2, Y: 2}},
//	}
//	orig := points[0]
//	end := points[3]
//	for i := 0; i < 5; i++ {
//		for j := 0; j < 5; j++ {
//			p := P{X: i, Y: j}
//			emptyPoints[p] = &PAstar{P: p, EmptyPoints: emptyPoints}
//		}
//	}
//
//	for i := range points {
//		p := points[i].P
//		points[i] = emptyPoints[p]
//		points[i].EmptyPoints = emptyPoints
//		//delete(emptyPoints, p)
//
//		fmt.Println(emptyPoints[p] == points[i])
//	}
//
//	pathers, dist, found := astar.Path(orig, end)
//	assert.Equal(true, found, "found")
//	assert.Equal(4.0, dist, "distance")
//	assert.Equal(1, pathers[len(pathers)-2].(*PAstar).P.X, "next X")
//	assert.Equal(0, pathers[len(pathers)-2].(*PAstar).P.Y, "next Y")
//	fmt.Println(pathers, dist, found)
//}
