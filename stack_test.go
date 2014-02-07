package wspacego

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
			sut := new(Stack)
			sut.Push(1)
			It("evaluates actual == expected", func() {
				Expect(sut.Pop()).To(Equal, 1)
			})
		})
		Context("push 2", func() {
			sut := new(Stack)
			sut.Push(2)
			It("evaluates actual == expected", func() {
				Expect(sut.Pop()).To(Equal, 2)
			})
		})
		Context("push 1; push 2", func() {
			var sut Stack
			sut.Push(1)
			sut.Push(2)
			It("evaluates actual == expected", func() {
				Expect(sut.Pop()).To(Equal, 2)
				Expect(sut.Pop()).To(Equal, 1)
			})
		})
	})
}
