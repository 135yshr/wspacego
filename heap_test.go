package wspacego

import (
	. "github.com/r7kamura/gospel"
	"testing"
)

func TestNewHeap(t *testing.T) {
	Describe(t, "Initialize Heap", func() {
		Context("new instance", func() {
			It("not nil", func() {
				Expect(NewHeap()).To(Exist)
			})
		})
	})
}

func TestHeap(t *testing.T) {
	Describe(t, "Heap struct Tests", func() {
		Context("push k:0 v:0", func() {
			sut := make(Heap)
			sut.Push(0, 0)
			It("actual == expected", func() {
				Expect(sut.Pop(0)).To(Equal, 0)
			})
		})
		Context("push k:0 v:1", func() {
			sut := make(Heap)
			sut.Push(0, 1)
			It("actual == expected", func() {
				Expect(sut.Pop(0)).To(Equal, 1)
			})
		})
		Context("push k:v 0:1 1:2", func() {
			sut := make(Heap)
			sut.Push(0, 1)
			sut.Push(1, 2)
			It("actual == expected", func() {
				Expect(sut.Pop(0)).To(Equal, 1)
				Expect(sut.Pop(1)).To(Equal, 2)
			})
		})
		Context("push k:v 0:1 0:2", func() {
			sut := make(Heap)
			sut.Push(0, 1)
			sut.Push(0, 2)
			It("actual == expected", func() {
				Expect(sut.Pop(0)).To(Equal, 2)
			})
		})
	})
}
