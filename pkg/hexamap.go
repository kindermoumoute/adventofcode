package pkg

const (
	HEXASOUTH = iota
	HEXASOUTHEAST
	HEXANORTHEAST
	HEXANORTH
	HEXANORTHWEST
	HEXASOUTHWEST
)

//    \ n  /
// nw +--+ ne
//   /    \
// --+      +-
//   \    /
// sw +--+ se
//   / s  \

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

func NewHexaPoint() *HexaP {
	return &HexaP{}
}

func HexaDist(x, y int) int {
	return (Abs(x) + Abs(y) + Abs(x+y)) / 2
}

func (p *HexaP) MoveSouth(steps int) {
	p.CurrentDirection = HEXASOUTH
	p.Move(steps)
}

func (p *HexaP) MoveSouthEast(steps int) {
	p.CurrentDirection = HEXASOUTHEAST
	p.Move(steps)
}

func (p *HexaP) MoveNorthEast(steps int) {
	p.CurrentDirection = HEXANORTHEAST
	p.Move(steps)
}

func (p *HexaP) MoveNorth(steps int) {
	p.CurrentDirection = HEXANORTH
	p.Move(steps)
}
func (p *HexaP) MoveNorthWest(steps int) {
	p.CurrentDirection = HEXANORTHWEST
	p.Move(steps)
}

func (p *HexaP) MoveSouthWest(steps int) {
	p.CurrentDirection = HEXASOUTHWEST
	p.Move(steps)
}

func (p *HexaP) Move(steps int) {
	p.X += HexaDirectionMap[p.CurrentDirection].x * steps
	p.Y += HexaDirectionMap[p.CurrentDirection].y * steps
}

func (p *HexaP) TurnLeft() {
	p.CurrentDirection = (p.CurrentDirection - 1 + 6) % 6
}

func (p *HexaP) TurnRight() {
	p.CurrentDirection = (p.CurrentDirection + 1) % 6
}

func (p *HexaP) DistFromOrigin() int {
	return p.DistFrom(&P{X: 0, Y: 0})
}

func (p *HexaP) DistFrom(from *P) int {
	return HexaDist(p.X-from.X, p.Y-from.Y)
}
