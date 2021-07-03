package container_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yuyamada/goder/pkg/container"
)

func TestUnionFind(t *testing.T) {
	u := container.NewUnionFind(4)
	u.Unite(0, 1)
	u.Unite(1, 2)
	assert.True(t, u.IsSame(0, 0))
	assert.True(t, u.IsSame(0, 1))
	assert.True(t, u.IsSame(0, 2))
	assert.False(t, u.IsSame(0, 3))
	assert.True(t, u.IsSame(1, 1))
	assert.True(t, u.IsSame(1, 2))
	assert.False(t, u.IsSame(1, 3))
	assert.True(t, u.IsSame(2, 2))
	assert.False(t, u.IsSame(2, 3))
	assert.True(t, u.IsSame(3, 3))
}
