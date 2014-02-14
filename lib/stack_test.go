package lib

import (
	. "github.com/r7kamura/gospel"
	"testing"
)

func TestPush(t *testing.T) {
	Describe(t, "Stack Push/Pop", func() {
		Context("Create Instance", func() {
			It("not nil", func() {
				Expect(NewStack()).To(Exist)
			})
		})
		Context("push 1", func() {
			It("evaluates actual == expected", func() {
				sut := NewStack()
				sut.Push(1)
				Expect(sut.Pop()).To(Equal, 1)
			})
		})
		Context("push 2", func() {
			It("evaluates actual == expected", func() {
				sut := NewStack()
				sut.Push(2)
				Expect(sut.Pop()).To(Equal, 2)
			})
		})
		Context("push 1; push 2", func() {
			It("evaluates actual == expected", func() {
				sut := NewStack()
				sut.Push(1)
				sut.Push(2)
				Expect(sut.Pop()).To(Equal, 2)
				Expect(sut.Pop()).To(Equal, 1)
			})
		})
	})
}
