package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestCheckPass(t *testing.T) {
	testCases := map[string]bool{
		"112233": true,
		"111122": true,
		"111233": true,
		"111111": false,
		"123444": false,
		"223450": false,
		"123789": false,
	}
	for tt, expected := range testCases {
		assert.Equal(t, expected, checkPassword([]byte(tt)), "%s fail", tt)
	}
}
