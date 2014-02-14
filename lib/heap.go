package lib

type Heap map[int]int

func (h Heap) Push(k, v int) {
	h[k] = v
}

func (h Heap) Pop(k int) int {
	return h[k]
}

func NewHeap() *Heap {
	ret := make(Heap)
	return &ret
}
