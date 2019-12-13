package twod

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap_Count(t *testing.T) {
	input := `.V..#
.....
#####
....#
...V#`

	m := NewMapFromInput(input).Translate(-2+2i).Filter('#', 'V')
	assert.Equal(t, []Vector{-2 + 0i, -1 + 0i, 0 + 0i, 1 - 2i, 2 - 1i, 1 + 0i, -1 + 2i, 2 - 2i, 2 + 0i, 2 + 2i}, m.SortByDist(-2-1i))
	assert.Equal(t, 8, len(m.Find('#')))
	assert.Equal(t, []Vector{1 - 2i, -1 + 2i}, m.Find('V'))
}

func TestSortedMap(t *testing.T) {
	m := Map{
		0:          "origin",
		1 + 3i:     "A",
		244 - 15i:  "B",
		-4 - 3i:    "C",
		-4 - 3000i: "D",
		-5 - 1i:    "E",
	}
	assert.Equal(t, []interface{}{"origin", "A", "E", "C", "B", "D"}, m.SortValuesByDist(0))
}
