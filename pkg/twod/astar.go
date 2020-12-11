package twod

import (
	"github.com/beefsack/go-astar"
)

func (p *P) Clone() *P {
	return &P{
		Pos:           p.Pos,
		Speed:         p.Speed,
		Steps:         p.Steps,
		emptyPoints:   p.emptyPoints.Clone(),
		exitCondition: p.exitCondition,
	}
}

func (p *P) LongestPos(emptyPoints Map) (maxPos Vector, maxDist int) {
	p.emptyPoints = emptyPoints
	return p.longestPath(0)
}

func (p *P) longestPath(currentDist int) (maxPos Vector, maxDist int) {
	maxPos = p.Pos
	maxDist = currentDist
	for _, neighbor := range p.PathNeighbors() {
		newPos, dist := neighbor.(*P).longestPath(currentDist + 1)
		if dist > maxDist {
			maxDist = dist
			maxPos = newPos
		}
	}
	return
}

func (p *P) ShortestPathToPos(pos Vector, emptyPoints Map) (path []*P, distance float64, found bool) {
	p.emptyPoints = emptyPoints
	target := NewPoint(pos, 0)

	// exit condition is when current point pos is target position
	p.exitCondition = func(currentPoint *P, neighbors []astar.Pather) []astar.Pather {
		if currentPoint.Pos == pos {
			return []astar.Pather{target}
		}
		return neighbors
	}

	// find shortest path
	patherPath, dist, found := astar.Path(p, target)

	// convert path to list of point
	pointPath := []*P(nil)
	for _, step := range patherPath {
		pointPath = append(pointPath, step.(*P))
	}

	return pointPath, dist, found
}

func (p *P) PathNeighbors() []astar.Pather {
	p.emptyPoints.Render()
	adjs := make([]astar.Pather, 0, 4)
	for _, dir := range FourDirections {

		// move the point to the adjacent position
		p.Speed = dir
		p.Move(1)

		// if the position is valid, explore it
		_, exist := p.emptyPoints[p.Pos]
		if exist {
			newP := p.Clone()
			delete(newP.emptyPoints, p.Pos) // remove position already explored
			adjs = append(adjs, newP)
		}
		p.Move(-1)
	}

	// if this position is the end point, we found it
	if p.exitCondition != nil {
		return p.exitCondition(p, adjs)
	}

	return adjs
}

// PathNeighborCost returns the cost of the tube leading to Truck.
func (p *P) PathNeighborCost(to astar.Pather) float64 {
	return 1.0
}

// PathEstimatedCost uses Manhattan distance to estimate orthogonal distance
func (p *P) PathEstimatedCost(to astar.Pather) float64 {
	return float64(p.Pos.DistFrom(to.(*P).Pos))
}
