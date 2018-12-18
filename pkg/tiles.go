package pkg

import (
	"github.com/beefsack/go-astar"
)

type Tile struct {
	P           P
	RowLenght   int
	EmptyPoints map[P]*Tile
}

func (t *Tile) SetEmptyTiles(c map[P]*Tile) {
	t.EmptyPoints = c
}

func (t *Tile) PathNeighbors() []astar.Pather {
	adjs := make([]astar.Pather, 0, 4)

	for _, newP := range t.P.FourWays() {
		ringP, exist := t.EmptyPoints[newP]
		if exist {
			//fmt.Print(" EXIST ", newP)
			ringP.EmptyPoints = t.EmptyPoints
			adjs = append(adjs, ringP)
		}
	}
	//fmt.Println()
	return adjs
}

// PathNeighborCost returns the cost of the tube leading to Truck.
func (t *Tile) PathNeighborCost(to astar.Pather) float64 {
	return 1.0
}

// PathEstimatedCost uses Manhattan distance to estimate orthogonal distance
func (t *Tile) PathEstimatedCost(to astar.Pather) float64 {
	return float64(t.P.DistFrom(&to.(*Tile).P))
}

func (t *Tile) Score() int {
	return t.P.Y*t.RowLenght + t.P.X
}

func (t *Tile) String() string {
	return t.P.String()
}
