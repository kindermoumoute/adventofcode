package twod

import (
	"math"
	"sort"
	"strings"
)

type Map map[Vector]interface{}

func NewBinaryMapFromInput(input string, origin Vector, matches ...interface{}) Map {
	m := make(Map)
	for j, line := range strings.Split(input, "\n") {
		for i, char := range line {
			for _, match := range matches {
				if char == match {
					m[NewVector(i, -j)-origin] = char
				}
			}
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

func (m Map) Center() Map {
	return m.Translate(-m.FindTopLeft())
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
	maxX, maxY := 0, 0
	xSet, ySet := false, false
	for k := range m {
		if !xSet || maxX < k.X() {
			xSet = true
			maxX = k.X()
		}
		if !ySet || maxY < k.Y() {
			ySet = true
			maxY = k.Y()
		}
	}
	return NewVector(maxX, maxY)
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
