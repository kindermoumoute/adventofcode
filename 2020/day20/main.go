package main

import (
	"fmt"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/adventofcode/pkg/execute"
	"github.com/kindermoumoute/adventofcode/pkg/twod"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	tiles := parse(input)

	part1, part2 := 0, 0

	//tilesMap := make(twod.Map)
	borders := make(map[int][]*Rotation)
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
		addBorder(borders, &Rotation{
			ID:        id,
			Value:     topBorder,
			Direction: twod.UP,
		})
		addBorder(borders, &Rotation{
			ID:        id,
			Value:     shiftedTopBorder,
			Direction: twod.UP,
			Shifted:   true,
		})
		addBorder(borders, &Rotation{
			ID:        id,
			Value:     botBorder,
			Direction: twod.DOWN,
		})
		addBorder(borders, &Rotation{
			ID:        id,
			Value:     shiftedBotBorder,
			Direction: twod.DOWN,
			Shifted:   true,
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
		addBorder(borders, &Rotation{
			ID:        id,
			Value:     leftBorder,
			Direction: twod.LEFT,
		})
		addBorder(borders, &Rotation{
			ID:        id,
			Value:     shiftedLeftBorder,
			Direction: twod.LEFT,
			Shifted:   true,
		})
		addBorder(borders, &Rotation{
			ID:        id,
			Value:     rightBorder,
			Direction: twod.RIGHT,
		})
		addBorder(borders, &Rotation{
			ID:        id,
			Value:     shiftedRightBorder,
			Direction: twod.RIGHT,
			Shifted:   true,
		})
	}
	//// 2D lib
	//
	//m := twod.NewMapFromInput(input)
	//m:= make( twod.Map)
	//twod.NewPoint()
	////
	adjacent := make(map[int]int)
	for _, rotations := range borders {
		fmt.Println(len(rotations))
		for _, rotation := range rotations {
			if len(rotations) > 1 {
				adjacent[rotation.ID]++
			}
			fmt.Print(rotation)
		}
	}

	part1 = 1
	for id, value := range adjacent {
		fmt.Println(id, "adjacent count ", value)
		if value == 4 {
			part1 *= id
		}
	}

	return part1, part2
}

func addBorder(borders map[int][]*Rotation, rotation *Rotation) {
	_, exist := borders[rotation.Value]
	if !exist {
		borders[rotation.Value] = []*Rotation(nil)
	}
	borders[rotation.Value] = append(borders[rotation.Value], rotation)
}

type Rotation struct {
	Value     int
	ID        int
	Direction twod.Vector
	Shifted   bool
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
	execute.Run(run, nil, puzzle, true)
}
