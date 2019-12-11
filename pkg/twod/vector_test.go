package twod

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVector(t *testing.T) {
	type cases []struct {
		x, y, xb, yb          int
		expectedDist          int
		expectedEuclidianDist float64
	}
	tests := cases{
		{
			0, 0, 1, 1,
			2,
			math.Hypot(1, 1),
		},
		{
			0, 0, 3, 5,
			8,
			math.Hypot(3, 5),
		},
	}
	for _, tt := range tests {
		t.Run("testnewvector", func(t *testing.T) {
			a := NewVector(tt.x, tt.y)
			assert.Equal(t, tt.x, a.X())
			assert.Equal(t, tt.y, a.Y())

			b := NewVector(tt.xb, tt.yb)
			assert.Equal(t, tt.xb, b.X())
			assert.Equal(t, tt.yb, b.Y())

			assert.Equal(t, tt.expectedDist, a.DistFrom(b))
			assert.Equal(t, tt.expectedEuclidianDist, a.EuclideanDistFrom(b))
		})
	}
}
