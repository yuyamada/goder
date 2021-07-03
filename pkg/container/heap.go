package container

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	tmp := *h
	*h = tmp[:len(tmp)-1]
	return tmp[len(tmp)-1]
}

type Ordered interface {
	Less(Ordered) bool
}

type PriorityQueue []Ordered

func (q PriorityQueue) Len() int {
	return len(q)
}

func (q PriorityQueue) Less(i, j int) bool {
	return q[i].Less(q[j])
}

func (q PriorityQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *PriorityQueue) Push(x interface{}) {
	*q = append(*q, x.(Ordered))
}

func (q *PriorityQueue) Pop() interface{} {
	tmp := *q
	*q = tmp[:len(tmp)-1]
	return tmp[len(tmp)-1]
}
