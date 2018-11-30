package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
	assert.Equal(t, 3, Max(3, 1, -200))
	assert.Equal(t, 5, Max(3, 1, 5))
	assert.Equal(t, -100, Max(-155, -100, -200))

	assert.Equal(t, -200, Min(3, 1, -200))
	assert.Equal(t, 1, Min(3, 1, 5))

	assert.Equal(t, 3, Abs(3))
	assert.Equal(t, 3, Abs(-3))
}
