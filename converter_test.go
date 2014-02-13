package wspacego

import (
	. "github.com/r7kamura/gospel"
	"testing"
)

func TestConerter(t *testing.T) {
	Describe(t, "whitespace のソースを文字や読みやすい文字列に変換する", func() {
		Context("インスタンスの生成", func() {
			It("インスタンスが作成されること", func() {
				Expect(NewConverter()).To(Exist)
			})
		})
		Context("スタックに関連する命令の生成", func() {
			It("スタックに１をプッシュするコマンドが作成されること", func() {
				data := []byte{' ', '\t', '\n'}
				sut := NewConverter()
				cmd, seek, err := sut.stackManipulation(data)
				Expect(err).To(NotExist)
				Expect(seek).To(Equal, len(data))
				Expect(cmd).To(Exist)
				Expect(cmd).To(Equal, NewSubCommandWithParam("stack", "push", 1))
			})
			It("スタックに2をプッシュするコマンドが作成されること", func() {
				data := []byte{' ', '\t', ' ', '\n'}
				sut := NewConverter()
				cmd, seek, err := sut.stackManipulation(data)
				Expect(err).To(NotExist)
				Expect(seek).To(Equal, len(data))
				Expect(cmd).To(Exist)
				Expect(cmd).To(Equal, NewSubCommandWithParam("stack", "push", 2))
			})
			It("スタックに4をプッシュするコマンドが作成されること", func() {
				data := []byte{' ', '\t', ' ', ' ', '\n'}
				sut := NewConverter()
				cmd, seek, err := sut.stackManipulation(data)
				Expect(err).To(NotExist)
				Expect(seek).To(Equal, len(data))
				Expect(cmd).To(Exist)
				Expect(cmd).To(Equal, NewSubCommandWithParam("stack", "push", 4))
			})
		})
	})
}
