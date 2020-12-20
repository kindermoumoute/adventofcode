package main

import (
	"fmt"
	"image/color"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
	"github.com/kindermoumoute/adventofcode/pkg/twod"
	"golang.org/x/image/colornames"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	twod.RenderingMap = map[interface{}]color.Color{
		'#': colornames.Black,
		'O': colornames.Red,
	}
	twod.SetFPS(10)
	tiles := parse(input)

	part1, part2 := 0, 0

	borders := make(map[int][]*Rotation)
	bordersPerTile := make(map[int][]*Rotation)
	for id, tile := range tiles {
		topLeft := tile.FindTopLeft()
		bottomRight := tile.FindBottomRight()
		topBorder, botBorder := 0, 0
		shiftedTopBorder, shiftedBotBorder := 0, 0
		shift := 0
		for x := topLeft.X(); x <= bottomRight.X(); x++ {
			if tile[twod.NewVector(x, topLeft.Y())] == '#' {
				topBorder += 1 << shift
				shiftedTopBorder += 1 << (9 - shift)
			}
			if tile[twod.NewVector(x, bottomRight.Y())] == '#' {
				botBorder += 1 << shift
				shiftedBotBorder += 1 << (9 - shift)
			}
			shift++
		}
		addBorder(borders, bordersPerTile, &Rotation{
			ID:        id,
			Value:     topBorder,
			Direction: twod.UP,
		})
		addBorder(borders, bordersPerTile, &Rotation{
			ID:        id,
			Value:     shiftedTopBorder,
			Direction: twod.UP,
			Inverted:  true,
		})
		addBorder(borders, bordersPerTile, &Rotation{
			ID:        id,
			Value:     botBorder,
			Direction: twod.DOWN,
		})
		addBorder(borders, bordersPerTile, &Rotation{
			ID:        id,
			Value:     shiftedBotBorder,
			Direction: twod.DOWN,
			Inverted:  true,
		})

		leftBorder, rightBorder := 0, 0
		shiftedLeftBorder, shiftedRightBorder := 0, 0
		shift = 0
		for y := topLeft.Y(); y >= bottomRight.Y(); y-- {
			if tile[twod.NewVector(topLeft.X(), y)] == '#' {
				leftBorder += 1 << shift
				shiftedLeftBorder += 1 << (9 - shift)
			}
			if tile[twod.NewVector(bottomRight.X(), y)] == '#' {
				rightBorder += 1 << shift
				shiftedRightBorder += 1 << (9 - shift)
			}
			shift++
		}
		addBorder(borders, bordersPerTile, &Rotation{
			ID:        id,
			Value:     leftBorder,
			Direction: twod.LEFT,
		})
		addBorder(borders, bordersPerTile, &Rotation{
			ID:        id,
			Value:     shiftedLeftBorder,
			Direction: twod.LEFT,
			Inverted:  true,
		})
		addBorder(borders, bordersPerTile, &Rotation{
			ID:        id,
			Value:     rightBorder,
			Direction: twod.RIGHT,
		})
		addBorder(borders, bordersPerTile, &Rotation{
			ID:        id,
			Value:     shiftedRightBorder,
			Direction: twod.RIGHT,
			Inverted:  true,
		})
	}

	adjacent := make(map[int]int)
	for _, rotations := range borders {
		for _, rotation := range rotations {
			if len(rotations) > 1 {
				adjacent[rotation.ID]++
			}
		}
	}

	part1 = 1
	for id, value := range adjacent {
		if value == 4 {
			part1 *= id
		}
	}

	// part 2
	bordersPerTileID := make(map[int]map[twod.Vector][]*Border)
	arbitraryTile := -1
	for id, tile := range tiles {
		if arbitraryTile == -1 {
			arbitraryTile = id
			continue
		}
		bordersPerTileID[id] = make(map[twod.Vector][]*Border)
		for i := 0; i < 8; i++ {
			tile = tile.SetPositive()
			tileBorders := findBorders(tile)
			for dir, borders := range tileBorders {
				for _, b := range borders {
					b.Tile = tile
					b.TileID = id
				}
				bordersPerTileID[id][dir] = append(bordersPerTileID[id][dir], borders...)
			}
			tile = tile.RotateRight() // try every rotation
			if i%4 == 3 {
				tile = tile.InvertY() // try every inversion
			}
		}
	}

	m := tiles[arbitraryTile]
dance:
	for len(bordersPerTileID) > 0 {
		removeCouronne(m).Render()
		actualBordersPerType := findBorders(m)
		for _, direction := range twod.FourDirections {
			for id, potentialBordersPerDirection := range bordersPerTileID {
				for _, potentialBorder := range potentialBordersPerDirection[direction] {
					for _, actualBorder := range actualBordersPerType[direction.RotateDegree(180)] {
						if actualBorder.Positions.Equal(potentialBorder.Positions) {
							for pos, char := range potentialBorder.Tile {
								m[pos+actualBorder.BottomLeft] = char
							}
							delete(bordersPerTileID, id)
							continue dance
						}
					}
				}
			}
		}
	}

	seaMonster := twod.NewMapFromInput(`                  # 
#    ##    ##    ###
 #  #  #  #  #  #   `).Filter('#').SetPositive().Translate(-1i)
	fmt.Println(seaMonster)
	m = removeCouronne(m.Filter('#'))
	fmt.Println(m)
	for i := 0; i < 8; i++ {
		m.Render()
		monsters := make(twod.Map)
	dance2:
		for k := range m {
			potentialMonster := make(twod.Map)
			for k2 := range seaMonster {
				_, exist := m[k+k2]
				if !exist {
					continue dance2
				}
				//potentialMonster.Render()
				potentialMonster[k+k2] = 'O'
			}
			monsters = monsters.Merge(potentialMonster)
			if len(potentialMonster) > 0 {
				m.Merge(monsters).Render()
			}
		}
		if len(monsters) == 0 {
			m = m.RotateRight() // try every rotation
			if i%4 == 3 {
				m = m.InvertY() // try every inversion
			}
			continue
		}
		part2 = len(m) - len(monsters)
		break
	}

	return part1, part2
}

func removeCouronne(m twod.Map) twod.Map {
	cleanedMap := make(twod.Map)
	for pos, value := range m.SetPositive() {
		x := pos.X()
		y := pos.Y()
		if x%10 == 0 || x%10 == 9 || y%10 == 0 || y%10 == 9 {
			continue
		}
		newPos := twod.NewVector(x-x/10*2-1, y-y/10*2-1)
		cleanedMap[newPos] = value
	}
	return cleanedMap
}

type Border struct {
	Positions  twod.Map
	BottomLeft twod.Vector

	Tile   twod.Map
	TileID int
}

func findBorders(m twod.Map) map[twod.Vector][]*Border {
	p := twod.NewPoint(m.FindBottomLeft(), twod.UP)
	// find first pixel
	for {
		_, exist := m[p.Pos]
		if exist {
			break
		}
		p.Move(1)
	}

	// Follow borders
	seen := make(map[twod.Vector]struct{})
	borders := make(map[twod.Vector][]*Border)
	p.Steps = 0
	for {
		if _, exist := seen[p.Pos]; exist {
			break
		}
		seen[p.Pos] = struct{}{}

		b := &Border{
			Positions: make(twod.Map),
		}
		for i := 0; i < 10; i++ {
			b.Positions[p.Pos] = m[p.Pos]
			p.Move(1)
		}
		borderType := p.Speed.RotateDegree(90)

		switch borderType {
		case twod.LEFT:
			b.BottomLeft = b.Positions.FindBottomLeft() - 10
		case twod.RIGHT:
			b.BottomLeft = b.Positions.FindBottomLeft() + 1
		case twod.DOWN:
			b.BottomLeft = b.Positions.FindBottomLeft() - 10i
		case twod.UP:
			b.BottomLeft = b.Positions.FindBottomLeft() + 1i
		}
		b.Positions = b.Positions.SetPositive()

		_, exist := borders[borderType]
		if !exist {
			borders[borderType] = []*Border(nil)
		}
		borders[borderType] = append(borders[borderType], b)

		// Next movement
		p.TurnLeft()
		p.Move(1)
		if _, exist := m[p.Pos]; !exist {
			p.Move(-1)
			p.TurnRight()
			if _, exist := m[p.Pos]; !exist {
				p.Move(-1)
				p.TurnRight()
			}
		}
	}

	return borders
}

//type Tile struct {
//
//	AllRotations map[twod.Vector]twod.Map
//}

func addBorder(borders, bordersPerTile map[int][]*Rotation, rotation *Rotation) {
	_, exist := borders[rotation.Value]
	if !exist {
		borders[rotation.Value] = []*Rotation(nil)
	}
	_, exist = bordersPerTile[rotation.ID]
	if !exist {
		bordersPerTile[rotation.ID] = []*Rotation(nil)
	}
	borders[rotation.Value] = append(borders[rotation.Value], rotation)
	bordersPerTile[rotation.ID] = append(bordersPerTile[rotation.ID], rotation)
}

type Rotation struct {
	Value     int
	ID        int
	Direction twod.Vector
	Inverted  bool
}

func parse(s string) map[int]twod.Map {
	maps := strings.Split(s, "\n\n")
	m := make(map[int]twod.Map)
	for _, mapLines := range maps {
		tmp := strings.Split(mapLines, "\n")
		id := 0
		pkg.MustScanf(tmp[0], "Tile %d:", &id)
		m[id] = twod.NewMapFromInput(strings.Join(tmp[1:], "\n"))

	}
	return m
}

func main() {
	execute.RunWithPixel(run, nil, puzzle, true)
}
