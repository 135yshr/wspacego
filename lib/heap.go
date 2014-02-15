package lib

type Heap map[int]int

func newHeap() *Heap {
	ret := make(Heap)
	return &ret
}

func (h *Heap) Push(k, v int) {
	heap := *h
	heap[k] = v
	*h = heap
}

func (h *Heap) Pop(k int) int {
	heap := *h
	return heap[k]
}
