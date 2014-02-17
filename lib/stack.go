package lib

import (
	"fmt"
)

type Stack []int

func newStack() *Stack {
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

func (s *Stack) Swap() error {
	st := *s
	size := len(st)
	if size < 2 {
		return fmt.Errorf("Value does not contain only one.")
	}
	st[size-1], st[size-2] = st[size-2], st[size-1]
	*s = st
	return nil
}

func (s *Stack) Copy(n int) error {
	st := *s
	size := len(st)
	if n < 0 || size <= n {
		return fmt.Errorf("out of range! [%d]", n)
	}
	v := st[n]
	st.Push(v)
	*s = st
	return nil
}

func (s *Stack) Move(n int) error {
	st := *s
	size := len(st)
	v := st[n]
	for idx := n; idx < size-1; idx++ {
		st[idx] = st[idx+1]
	}
	st[size-1] = v
	*s = st
	return nil
}
