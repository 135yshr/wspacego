package wspacego

import (
	. "github.com/r7kamura/gospel"
	"path"
	"runtime"
	"testing"
)

func TestInterpretor(t *testing.T) {
	Describe(t, "インタープリターのテスト", func() {
		Context("インスタンス生成", func() {
			test_file := path.Join(current_dir(), "samples/hworld.ws")
			It("インスタンスが生成できること", func() {
				actual := NewInterpreter(test_file)
				Expect(actual).To(Exist)
				Expect(actual.path).To(Equal, test_file)
			})
		})
	})
}

func current_dir() string {
	_, fpath, _, _ := runtime.Caller(0)
	return path.Dir(fpath)
}
