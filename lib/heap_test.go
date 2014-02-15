package lib

import (
	. "github.com/r7kamura/gospel"
	"testing"
)

func TestnewHeap(t *testing.T) {
	Describe(t, "Initialize Heap", func() {
		Context("new instance", func() {
			It("not nil", func() {
				Expect(newHeap()).To(Exist)
			})
		})
	})
}

func TestHeap(t *testing.T) {
	Describe(t, "Heap struct Tests", func() {
		Context("push k:0 v:0", func() {
			It("actual == expected", func() {
				sut := newHeap()
				sut.Push(0, 0)
				Expect(sut.Pop(0)).To(Equal, 0)
			})
		})
		Context("push k:0 v:1", func() {
			It("actual == expected", func() {
				sut := newHeap()
				sut.Push(0, 1)
				Expect(sut.Pop(0)).To(Equal, 1)
			})
		})
		Context("push k:v 0:1 1:2", func() {
			It("actual == expected", func() {
				sut := newHeap()
				sut.Push(0, 1)
				sut.Push(1, 2)
				Expect(sut.Pop(0)).To(Equal, 1)
				Expect(sut.Pop(1)).To(Equal, 2)
			})
		})
		Context("push k:v 0:1 0:2", func() {
			It("actual == expected", func() {
				sut := newHeap()
				sut.Push(0, 1)
				sut.Push(0, 2)
				Expect(sut.Pop(0)).To(Equal, 2)
			})
		})
	})
}
