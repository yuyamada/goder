package container_test

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yuyamada/goder/pkg/container"
)

func TestIntHeap(t *testing.T) {
	data := []int{1, 3, 5, 2, 4}
	expected := []int{1, 2, 3, 4, 5}
	h := &container.IntHeap{}
	for _, v := range data {
		heap.Push(h, v)
	}
	actual := make([]int, 0, len(data))
	for range data {
		actual = append(actual, heap.Pop(h).(int))
	}
	assert.Equal(t, expected, actual)
}

type Pair struct{ x, y int }

func (p Pair) Less(other container.Ordered) bool {
	if p.x == other.(Pair).x {
		return p.y < other.(Pair).y
	}
	return p.x < other.(Pair).x
}

func TestPriorityQueue(t *testing.T) {
	data := []Pair{
		{1, 2}, {1, 1}, {2, 1}, {2, 2},
	}
	expected := []Pair{
		{1, 1}, {1, 2}, {2, 1}, {2, 2},
	}
	q := &container.PriorityQueue{}
	for _, v := range data {
		heap.Push(q, v)
	}
	actual := make([]Pair, 0, len(data))
	for range data {
		actual = append(actual, heap.Pop(q).(Pair))
	}
	assert.Equal(t, expected, actual)
}
