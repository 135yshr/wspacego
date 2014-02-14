package lib

import (
	. "github.com/r7kamura/gospel"
	"path"
	"runtime"
	"testing"
)

func TestInterpretor(t *testing.T) {
	Describe(t, "インタープリターのテスト", func() {
		data := []byte{'A', ' ', 'B', '\t', '\r', '\n'}
		Context("インスタンス生成", func() {
			It("インスタンスが生成できること", func() {
				actual := NewInterpreter(data)
				Expect(actual).To(Exist)
				Expect(actual.origin).To(Equal, data)
			})
		})
		Context("データを読み込んで命令を実行する", func() {
			It("不要なデータを排除してスペースやタブをそれぞれSとTに置き換えた文字を出力する", func() {
				data := []byte{' ', ' ', '	', ' ', ' ', ' ', ' ', ' ', '	', '\n'}
				expected := []byte{'S', 'S', 'T', 'S', 'S', 'S', 'S', 'S', 'T', '\n'}
				sut := NewInterpreter(data)
				dat, err := sut.ToChar()
				Expect(err).To(NotExist)
				Expect(dat).To(Equal, expected)
			})
			It("不要なデータを排除してスペースやタブをそれぞれSとTに置き換えた文字を出力する パート２", func() {
				data := []byte{' ', ' ', '	', ' ', ' ', ' ', ' ', '	', '	', '\n'}
				expected := []byte{'S', 'S', 'T', 'S', 'S', 'S', 'S', 'T', 'T', '\n'}
				sut := NewInterpreter(data)
				dat, err := sut.ToChar()
				Expect(err).To(NotExist)
				Expect(dat).To(Equal, expected)
			})
		})
		Context("不要な文字を排除する関数", func() {
			It("不要なデータ以外排除されていること", func() {
				sut := NewInterpreter(data)
				sut.filter()
				Expect(sut.source).To(Equal, []byte{' ', '\t', '\n'})
			})
			It("不要なデータ以外排除されていること（改行を増やす）", func() {
				data = append(data, '\n')
				sut := NewInterpreter(data)
				sut.filter()
				Expect(sut.source).To(Equal, []byte{' ', '\t', '\n', '\n'})
			})
		})
	})
}

func current_dir() string {
	_, fpath, _, _ := runtime.Caller(0)
	return path.Dir(fpath)
}
