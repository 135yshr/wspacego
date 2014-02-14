package lib

type Stack []int

func NewStack() *Stack {
	return new(Stack)
}

func (s *Stack) Pop() int {
	tmp := *s
	n := len(tmp) - 1
	ret := tmp[n]
	*s = tmp[0:n]
	return ret
}

func (s *Stack) Push(n int) {
	*s = append(*s, n)
}

func (s *Stack) Swap() {
	st := *s
	size := len(st)
	if size > 1 {
		st[size-1], st[size-2] = st[size-2], st[size-1]
		*s = st
	}
}
