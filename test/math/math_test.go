package math_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yuyamada/goder/pkg/math"
)

func TestAbs(t *testing.T) {
	for _, tt := range []struct{ data, expected int }{
		{1, 1}, {0, 0}, {-1, 1},
	} {
		assert.Equal(t, math.Abs(tt.data), tt.expected)
	}
}

func TestMax(t *testing.T) {
	x, y, z := 1, 0, -1
	expected := 1
	assert.Equal(t, expected, math.Max(x, y, z))
}

func TestMin(t *testing.T) {
	x, y, z := 1, 0, -1
	expected := -1
	assert.Equal(t, expected, math.Min(x, y, z))
}

func TestPow(t *testing.T) {
	base, n := 5, 5
	expected := 3125
	assert.Equal(t, expected, math.Pow(base, n))
}

func TestGCD(t *testing.T) {
	for _, tt := range []struct{ x, y, expected int }{
		{17, 16, 1}, {8, 4, 4}, {10, 0, 10}, {0, 10, 10}, {0, 0, 0},
	} {
		assert.Equal(t, tt.expected, math.GCD(tt.x, tt.y))
	}
}
