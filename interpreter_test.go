package wspacego

import (
	. "github.com/r7kamura/gospel"
	"testing"
)

func TestInterpretor(t *testing.T) {
	Describe(t, "インタープリターのテスト", func() {
		Context("インスタンス生成", func() {
			It("インスタンスが生成できること", func() {
				Expect(NewInterpreter()).To(Exist)
			})
		})
	})
}
