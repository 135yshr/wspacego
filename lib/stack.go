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
