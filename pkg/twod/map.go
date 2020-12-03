package twod

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

type Map map[Vector]interface{}

// NewMapFromInput set every character from the input string into a Map object
// based on its coordinate. Every vector in the map is positive, which means that
// the bottom left of the input will be at the origin of the Map.
func NewMapFromInput(input string) Map {
	m := make(Map)
	lines := strings.Split(input, "\n")
	for j, line := range lines {
		for i, char := range line {
			m[NewVector(i, len(lines)-1-j)] = char
		}
	}
	return m
}

func (m Map) Clone() Map {
	clone := make(Map)
	for k, v := range m {
		clone[k] = v
	}
	return clone
}

func (m Map) Filter(matches ...interface{}) Map {
	clone := make(Map)
	for k, v := range m {
		if hasMatch(v, matches) {
			clone[k] = v
		}
	}
	return clone
}

func (m Map) RotateLeft() Map {
	clone := make(Map)
	for k, v := range m {
		clone[k.Rotate(math.Pi/2)] = v
	}
	return clone
}

func (m Map) RotateRight() Map {
	clone := make(Map)
	for k, v := range m {
		clone[k.Rotate(-math.Pi/2)] = v
	}
	return clone
}

func (m Map) Translate(move Vector) Map {
	clone := make(Map)
	for k, v := range m {
		clone[k+move] = v
	}
	return clone
}

func (m Map) InvertY() Map {
	clone := make(Map)
	for k, v := range m {
		clone[NewVector(k.X(), -k.Y())] = v
	}
	return clone
}

func (m Map) SetPositive() Map {
	return m.Translate(-m.FindBottomLeft())
}

func (m Map) ToSlice() []Vector {
	slice := []Vector(nil)
	for k := range m {
		slice = append(slice, k)
	}
	return slice
}

// SortByDist will sort the sliced map by distance from the given position (lowest to max).
// If two points have the same distance the order will be based on radian angle (lowest to max).
func (m Map) SortByDist(from Vector) []Vector {
	sortedKeys := m.ToSlice()
	sort.Slice(sortedKeys, func(i, j int) bool {
		if sortedKeys[i].DistFrom(from) < sortedKeys[j].DistFrom(from) {
			return true
		}
		if sortedKeys[i].DistFrom(from) > sortedKeys[j].DistFrom(from) {
			return false
		}
		return sortedKeys[i].AngleFrom(from) < sortedKeys[j].AngleFrom(from)
	})
	return sortedKeys
}

func (m Map) SortValuesByDist(from Vector) []interface{} {
	sortedValues := []interface{}(nil)
	for _, k := range m.SortByDist(from) {
		sortedValues = append(sortedValues, m[k])
	}
	return sortedValues
}

func (m Map) Find(matches ...interface{}) []Vector {
	sorted := m.SortByDist(0)
	founds := []Vector(nil)
	for _, k := range sorted {
		if hasMatch(m[k], matches) {
			founds = append(founds, k)
		}
	}
	return founds
}

// FindTopLeft returns a point that represent the top left corner of a squared map only composed of the match values.
// This point might not exist on the actual map.
func (m Map) FindTopLeft() Vector {
	minX, maxY := 0, 0
	xSet, ySet := false, false
	for k := range m {
		if !xSet || k.X() < minX {
			xSet = true
			minX = k.X()
		}
		if !ySet || k.Y() > maxY {
			ySet = true
			maxY = k.Y()
		}
	}
	return NewVector(minX, maxY)
}

// FindTopRight returns a point that represent the top right corner of a squared map only composed of the match values.
// This point might not exist on the actual map.
func (m Map) FindTopRight() Vector {
	maxX, maxY := 0, 0
	xSet, ySet := false, false
	for k := range m {
		if !xSet || k.X() > maxX {
			xSet = true
			maxX = k.X()
		}
		if !ySet || k.Y() > maxY {
			ySet = true
			maxY = k.Y()
		}
	}
	return NewVector(maxX, maxY)
}

// FindBottomLeft returns a point that represent the bottom left corner of a squared map only composed of the match values.
// This point might not exist on the actual map.
func (m Map) FindBottomLeft() Vector {
	minX, minY := 0, 0
	xSet, ySet := false, false
	for k := range m {
		if !xSet || k.X() < minX {
			xSet = true
			minX = k.X()
		}
		if !ySet || k.Y() < minY {
			ySet = true
			minY = k.Y()
		}
	}
	return NewVector(minX, minY)
}

// FindBottomRight returns a point that represent the bottom right corner of a squared map only composed of the match values.
// This point might not exist on the actual map.
func (m Map) FindBottomRight() Vector {
	maxX, minY := 0, 0
	xSet, ySet := false, false
	for k := range m {
		if !xSet || maxX < k.X() {
			xSet = true
			maxX = k.X()
		}
		if !ySet || minY > k.Y() {
			ySet = true
			minY = k.Y()
		}
	}
	return NewVector(maxX, minY)
}

// Width returns width of the map.
func (m Map) Width() int {
	return m.FindBottomRight().X() - m.FindBottomLeft().X() + 1
}

// Height returns height of the map.
func (m Map) Height() int {
	return m.FindBottomRight().Y() - m.FindTopRight().Y() + 1
}

func (m Map) String() string {
	topLeft := m.FindTopLeft()
	botRight := m.FindBottomRight()

	// start with headers
	s := lineHeaderString(topLeft.X(), botRight.X())

	// display map
	grad := 0
	for j := topLeft.Y(); j >= botRight.Y(); j-- {
		if grad%5 == 0 {
			s += fmt.Sprintf("%d\t ┤", j)
		} else {
			s += "\t │"
		}
		for i := topLeft.X(); i <= botRight.X(); i++ {
			v, exist := m[NewVector(i, j)]
			if exist {
				s += printValue(v)
			} else {
				s += " "
			}
		}
		s += "\n"
		grad++
	}
	return s
}

func printValue(v interface{}) string {
	ascii, isASCII := v.(int32)
	if isASCII {
		return fmt.Sprintf("%c", ascii)
	}
	return fmt.Sprint(v)
}

func lineHeaderString(leftX, rightX int) string {
	header1 := "\t  "
	header2 := "\t ┌"
	grad := 0
	for i := leftX; i <= rightX; i++ {
		if grad%5 == 0 {
			n := fmt.Sprint(i)
			i += len(n) - 1
			grad += len(n) - 1
			header1 += n
			header2 += "┴" + strings.Repeat("─", len(n)-1)
		} else {
			header1 += " "
			header2 += "─"
		}
		grad++
	}
	return header1 + "\n" + header2 + "\n"
}

func hasMatch(v interface{}, matches []interface{}) bool {
	if len(matches) == 0 {
		return true
	}
	for _, match := range matches {
		if v == match {
			return true
		}
	}
	return false
}

// TODO: render map
