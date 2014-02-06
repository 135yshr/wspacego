package wspacego

type Heap map[int]int

func (h Heap) Push(k, v int) {
	h[k] = v
}

func (h Heap) Pop(k int) int {
	return h[k]
}
