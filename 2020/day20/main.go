package main

import (
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
	twod.SetFPS(60)
	bordersPerTileID, tileMap, seaMap := parse(input)

	part1, part2 := 0, 0
buildingMap:
	for len(bordersPerTileID) > 0 {
		//removeBorders(seaMap).Render()
		actualBordersPerType := findBorders(seaMap)
		for _, direction := range twod.FourDirections {
			for id, potentialBordersPerDirection := range bordersPerTileID {
				for _, potentialBorder := range potentialBordersPerDirection[direction] {
					for _, actualBorder := range actualBordersPerType[direction.RotateDegree(180)] {
						if actualBorder.Positions.Equal(potentialBorder.Positions) {
							tileMap[actualBorder.BottomLeft] = id
							for pos, char := range potentialBorder.Tile {
								seaMap[pos+actualBorder.BottomLeft] = char
							}
							delete(bordersPerTileID, id)
							continue buildingMap
						}
					}
				}
			}
		}
	}
	part1 = tileMap[tileMap.FindTopLeft()].(int)
	part1 *= tileMap[tileMap.FindTopRight()].(int)
	part1 *= tileMap[tileMap.FindBottomRight()].(int)
	part1 *= tileMap[tileMap.FindBottomLeft()].(int)

	seaMonster := twod.NewMapFromInput(`                  # 
#    ##    ##    ###
 #  #  #  #  #  #   `).Filter('#')
	seaMap = removeBorders(seaMap.Filter('#'))
	for i := 0; i < 8; i++ {
		//seaMap.Render()
		matches := seaMap.FindPattern(seaMonster, false)
		if len(matches) != 0 {
			part2 = len(seaMap)
			for _, pattern := range matches {
				part2 -= len(pattern)
			}
			break
		}

		seaMap = seaMap.RotateRight() // try every rotation
		if i%4 == 3 {
			seaMap = seaMap.InvertY() // try every inversion
		}
	}

	return part1, part2
}

func removeBorders(m twod.Map) twod.Map {
	cleanedMap := make(twod.Map)
	for pos, value := range m.SetPositive() {
		x, y := pos.X(), pos.Y()
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

// findBorders will return all borders sorted by side (UP,DOWN,LEFT,RIGHT).
func findBorders(m twod.Map) map[twod.Vector][]*Border {
	// find first pixel
	p := twod.NewPoint(m.FindBottomLeft(), twod.UP)
	for _, exist := m[p.Pos]; !exist; _, exist = m[p.Pos] {
		p.Move(1)
	}

	// Follow borders
	seen := make(map[twod.Vector]struct{})
	borders := make(map[twod.Vector][]*Border)
	for exist := false; !exist; _, exist = seen[p.Pos] {
		seen[p.Pos] = struct{}{}

		b := &Border{Positions: make(twod.Map)}
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

		_, exist = borders[borderType]
		if !exist {
			borders[borderType] = []*Border(nil)
		}
		borders[borderType] = append(borders[borderType], b)

		// Next movement
		p.TurnLeft()
		p.Move(1)
		if _, exist = m[p.Pos]; !exist {
			p.Move(-1)
			p.TurnRight()
			if _, exist = m[p.Pos]; !exist {
				p.Move(-1)
				p.TurnRight()
			}
		}
	}

	return borders
}

func parse(s string) (map[int]map[twod.Vector][]*Border, twod.Map, twod.Map) {
	maps := strings.Split(s, "\n\n")

	bordersPerTileID := make(map[int]map[twod.Vector][]*Border)
	tileMap := make(twod.Map)
	var seaMap twod.Map
	for _, mapLines := range maps {
		tmp := strings.Split(mapLines, "\n")
		id := 0
		pkg.MustScanf(tmp[0], "Tile %d:", &id)
		tile := twod.NewMapFromInput(strings.Join(tmp[1:], "\n"))
		if seaMap == nil {
			tileMap[0] = id
			seaMap = tile
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

	return bordersPerTileID, tileMap, seaMap
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
