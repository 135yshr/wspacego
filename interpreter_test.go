package wspacego

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
		Context("ソースファイルをコマンドリストに変換する関数", func() {
			It("スタックに１をプッシュするコマンドが作成されること", func() {
				data = []byte{'P', 'u', 's', 'h', ' ', ' ', '\t', '\n'}
				sut := NewInterpreter(data)
				sut.filter()
				sut.parseCommands()
				Expect(sut.commands).To(Exist)
				Expect(sut.commands.Len()).To(Equal, 1)
				Expect(sut.commands.Get(1)).To(Equal, NewSubCommandWithParam("stack", "push", 1))
			})
			It("スタックをコピーするコマンドが作成されること", func() {
				data = []byte{'C', 'o', 'p', 'y', ' ', '\n', ' '}
				sut := NewInterpreter(data)
				sut.filter()
				sut.parseCommands()
				Expect(sut.commands).To(Exist)
				Expect(sut.commands.Len()).To(Equal, 1)
				Expect(sut.commands.Get(1)).To(Equal, NewSubCommandWithParam("stack", "copy", 0))
			})
			It("スタックを入れ替えるするコマンドが作成されること", func() {
				data = []byte{'S', 'w', 'a', 'p', ' ', '\n', '\t'}
				sut := NewInterpreter(data)
				sut.filter()
				sut.parseCommands()
				Expect(sut.commands).To(Exist)
				Expect(sut.commands.Len()).To(Equal, 1)
				Expect(sut.commands.Get(1)).To(Equal, NewSubCommandWithParam("stack", "swap", 0))
			})
		})
	})
}

func current_dir() string {
	_, fpath, _, _ := runtime.Caller(0)
	return path.Dir(fpath)
}
