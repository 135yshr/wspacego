package lib

import (
	. "github.com/r7kamura/gospel"
	"testing"
)

func TestPush(t *testing.T) {
	Describe(t, "Stack Push/Pop", func() {
		Context("インスタンス生成", func() {
			It("インスタンスが作成されること", func() {
				Expect(newStack()).To(Exist)
			})
		})
		Context("push/popメソッド", func() {
			It("１を格納して１を取得できること", func() {
				sut := newStack()
				sut.Push(1)
				Expect(sut.Pop()).To(Equal, 1)
			})
			It("２を格納して２を取得できること", func() {
				sut := newStack()
				sut.Push(2)
				Expect(sut.Pop()).To(Equal, 2)
			})
			It("１と２を格納して２→１の順に取得できること", func() {
				sut := newStack()
				sut.Push(1)
				sut.Push(2)
				Expect(sut.Pop()).To(Equal, 2)
				Expect(sut.Pop()).To(Equal, 1)
			})
		})
		Context("１番目と２番めの値を入れ替える", func() {
			It("値が入れ替わること", func() {
				sut := newStack()
				sut.Push(1)
				sut.Push(10)
				sut.Swap()
				Expect(sut.Pop()).To(Equal, 1)
				Expect(sut.Pop()).To(Equal, 10)
			})
			It("１番目と２番目の値が入れ替わること", func() {
				sut := newStack()
				sut.Push(1)
				sut.Push(10)
				sut.Push(20)
				sut.Swap()
				Expect(sut.Pop()).To(Equal, 10)
				Expect(sut.Pop()).To(Equal, 20)
				Expect(sut.Pop()).To(Equal, 1)
			})
			It("格納した値が１つだけのとき内容が変わらないこと", func() {
				sut := newStack()
				sut.Push(1)
				sut.Swap()
				Expect(sut.Pop()).To(Equal, 1)
			})
		})
		Context("スタックのn番目の値をトップにコピーする", func() {
			It("1番目を指定したときスタックのトップに値がコピーされること", func() {
				sut := newStack()
				sut.Push(1)
				sut.Push(2)
				sut.Copy(0)
				Expect(sut.Pop()).To(Equal, 1)
			})
			It("2番目を指定したときスタックのトップに値がコピーされること", func() {
				sut := newStack()
				sut.Push(1)
				sut.Push(2)
				sut.Copy(1)
				Expect(sut.Pop()).To(Equal, 2)
			})
			It("3番目を指定したときエラーが返ってくること", func() {
				sut := newStack()
				sut.Push(1)
				sut.Push(2)
				err := sut.Copy(3)
				Expect(err).To(Exist)
				Expect(sut.Pop()).To(Equal, 2)
			})
		})
	})
}
