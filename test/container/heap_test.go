package container_test

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yuyamada/goder/pkg/container"
)

func TestIntHeap(t *testing.T) {
	h := &container.IntHeap{}
	data := []int{1, 3, 5, 2, 4}
	expected := []int{1, 2, 3, 4, 5}
	for _, v := range data {
		heap.Push(h, v)
	}
	actual := make([]int, 0, len(data))
	for range data {
		actual = append(actual, heap.Pop(h).(int))
	}
	assert.Equal(t, expected, actual)
}
