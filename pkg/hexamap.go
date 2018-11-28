package pkg

const (
	HEXASOUTH = iota
	HEXASOUTHEAST
	HEXANORTHEAST
	HEXANORTH
	HEXANORTHWEST
	HEXASOUTHWEST
)

var HexaDirectionMap = map[int]struct {
	x, y int
	name string
}{
	HEXASOUTH:     {1, 0, "SOUTH"},
	HEXASOUTHEAST: {1, -1, "SOUTHEAST"},
	HEXANORTHEAST: {0, -1, "NORTHEAST"},
	HEXANORTH:     {-1, 0, "NORTH"},
	HEXANORTHWEST: {-1, 1, "NORTHWEST"},
	HEXASOUTHWEST: {0, 1, "SOUTHWEST"},
}

type HexaP struct {
	X, Y             int
	CurrentDirection int
}

func HexaDist(x, y int) int {
	return (Abs(x) + Abs(y) + Abs(x+y)) / 2
}

func (p *HexaP) Move(steps int) {
	p.X += HexaDirectionMap[p.CurrentDirection].x * steps
	p.Y += HexaDirectionMap[p.CurrentDirection].y * steps
}

func (p *HexaP) Left() {
	p.CurrentDirection = (p.CurrentDirection - 1 + 6) % 6
}

func (p *HexaP) Right() {
	p.CurrentDirection = (p.CurrentDirection + 1) % 6
}

func (p *HexaP) DistFromOrigin() int {
	return p.DistFrom(&P{X: 0, Y: 0})
}

func (p *HexaP) DistFrom(from *P) int {
	return HexaDist(Abs(p.X-from.X), Abs(p.Y-from.Y))
}
