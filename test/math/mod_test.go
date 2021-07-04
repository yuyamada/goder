package math_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yuyamada/goder/pkg/math"
)

func TestMod_Mul(t *testing.T) {
	m := math.NewMod(7)
	for _, tt := range []struct{ x, y, expected int }{
		{4, 5, 6}, {7, 7, 0}, {0, 0, 0},
	} {
		assert.Equal(t, tt.expected, m.Mul(tt.x, tt.y))
	}
}

func TestMod_Pow(t *testing.T) {
	m := math.NewMod(7)
	for _, tt := range []struct{ base, n, expected int }{
		{2, 3, 1}, {4, 5, 2}, {1, 0, 1}, {0, 1, 0}, {0, 0, 1}, {4, -2, 4},
	} {
		assert.Equal(t, tt.expected, m.Pow(tt.base, tt.n))
	}
}

func TestMod_Inv(t *testing.T) {
	m := math.NewMod(7)
	for _, tt := range []struct{ x, expected int }{
		{2, 4}, {1, 1},
	} {
		assert.Equal(t, tt.expected, m.Inv(tt.x))
	}
}

func TestMod_Factorial(t *testing.T) {
	m := math.NewMod(7)
	for _, tt := range []struct{ x, expected int }{
		{2, 2}, {4, 3}, {0, 1},
	} {
		assert.Equal(t, tt.expected, m.Factorial(tt.x))
	}
}

func TestMod_InvFactorial(t *testing.T) {
	m := math.NewMod(7)
	m.InitInvFactorials(4)
	for _, tt := range []struct{ x, expected int }{
		{0, 1}, {1, 1}, {2, 4}, {3, 6}, {4, 5}, {5, 1}, {6, 6},
	} {
		t.Run(fmt.Sprintf("x=%d", tt.x), func(t *testing.T) {
			assert.Equal(t, tt.expected, m.InvFactorial(tt.x))
		})
	}
}

func TestMod_Combination(t *testing.T) {
	m := math.NewMod(7)
	m.InitInvFactorials(6)
	for _, tt := range []struct{ n, k, expected int }{
		{3, 0, 1}, {3, 1, 3}, {3, 2, 3}, {3, 3, 1}, {3, 4, 0}, {0, 0, 1},
	} {
		t.Run(fmt.Sprintf("n=%d k=%d", tt.n, tt.k), func(t *testing.T) {
			assert.Equal(t, tt.expected, m.Combination(tt.n, tt.k))
		})
	}
}
