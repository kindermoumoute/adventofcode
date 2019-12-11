package twod

import (
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

func (m Map) Find(match interface{}) []Vector {
	sorted := m.SortByDist(0)
	founds := []Vector(nil)
	for _, k := range sorted {
		if m[k] == match {
			founds = append(founds, k)
		}
	}
	return founds
}
