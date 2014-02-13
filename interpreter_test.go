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
				Expect(sut.commands.Get(1)).To(Equal, NewSubCommand("stack", "copy"))
			})
			It("スタックを入れ替えるコマンドが作成されること", func() {
				data = []byte{'S', 'w', 'a', 'p', ' ', '\n', '\t'}
				sut := NewInterpreter(data)
				sut.filter()
				sut.parseCommands()
				Expect(sut.commands).To(Exist)
				Expect(sut.commands.Len()).To(Equal, 1)
				Expect(sut.commands.Get(1)).To(Equal, NewSubCommand("stack", "swap"))
			})
			It("スタックを削除するコマンドが作成されること", func() {
				data = []byte{'R', 'e', 'm', 'o', 'v', 'e', ' ', '\n', '\n'}
				sut := NewInterpreter(data)
				sut.filter()
				sut.parseCommands()
				Expect(sut.commands).To(Exist)
				Expect(sut.commands.Len()).To(Equal, 1)
				Expect(sut.commands.Get(1)).To(Equal, NewSubCommand("stack", "remove"))
			})
			It("定義されていない命令が指定されたときにundefinedの命令が作成されること", func() {
				data = []byte{'u', 'n', 'k', 'n', 'o', 'w', 'n', ' ', '\t', '\n'}
				sut := NewInterpreter(data)
				sut.filter()
				err := sut.parseCommands()
				Expect(sut.commands).To(NotExist)
				Expect(err).To(Exist)
			})
			It("ラベルを定義するコマンドが作成されること", func() {
				data = []byte{'L', 'a', 'b', 'l', '\n', ' ', ' ', '\t', ' ', ' ', '\t', '\n'}
				sut := NewInterpreter(data)
				sut.filter()
				sut.parseCommands()
				Expect(sut.commands).To(Exist)
				Expect(sut.commands.Len()).To(Equal, 1)
				Expect(sut.commands.Get(1)).To(Equal, NewSubCommand("label", "1001"))
			})
			It("ラベルを呼び出すコマンドが作成されること", func() {
				data = []byte{'C', 'a', 'l', 'l', '\n', ' ', '\t', '\t', ' ', ' ', '\t', '\n'}
				sut := NewInterpreter(data)
				sut.filter()
				sut.parseCommands()
				Expect(sut.commands).To(Exist)
				Expect(sut.commands.Len()).To(Equal, 1)
				Expect(sut.commands.Get(1)).To(Equal, NewSubCommand("call", "1001"))
			})
			It("ラベルを呼び出すコマンドが作成されること", func() {
				data = []byte{'G', 'o', 't', 'o', '\n', ' ', '\n', '\t', ' ', ' ', '\t', '\n'}
				sut := NewInterpreter(data)
				sut.filter()
				sut.parseCommands()
				Expect(sut.commands).To(Exist)
				Expect(sut.commands.Len()).To(Equal, 1)
				Expect(sut.commands.Get(1)).To(Equal, NewSubCommand("goto", "1001"))
			})
			It("スタックの値が０のときにラベルを呼び出すコマンドが作成されること", func() {
				data = []byte{'=', '=', '0', 'G', 'o', 't', 'o', '\n', '\t', ' ', '\t', ' ', ' ', '\t', '\n'}
				sut := NewInterpreter(data)
				sut.filter()
				sut.parseCommands()
				Expect(sut.commands).To(Exist)
				Expect(sut.commands.Len()).To(Equal, 1)
				Expect(sut.commands.Get(1)).To(Equal, NewSubCommand("if stack==0 then goto", "1001"))
			})
			It("スタックの値が０未満のときにラベルを呼び出すコマンドが作成されること", func() {
				data = []byte{'<', '0', 'G', 'o', 't', 'o', '\n', '\t', '\t', '\t', ' ', ' ', '\t', '\n'}
				sut := NewInterpreter(data)
				sut.filter()
				sut.parseCommands()
				Expect(sut.commands).To(Exist)
				Expect(sut.commands.Len()).To(Equal, 1)
				Expect(sut.commands.Get(1)).To(Equal, NewSubCommand("if stack<0 then goto", "1001"))
			})
			It("呼び出し元に戻るコマンドが作成されること", func() {
				data = []byte{'R', 'e', 't', 'u', 'r', 'n', '\n', '\t', '\n'}
				sut := NewInterpreter(data)
				sut.filter()
				sut.parseCommands()
				Expect(sut.commands).To(Exist)
				Expect(sut.commands.Len()).To(Equal, 1)
				Expect(sut.commands.Get(1)).To(Equal, NewCommand("return"))
			})
			It("プログラムを終了するコマンドが作成されること", func() {
				data = []byte{'E', 'x', 'i', 't', '\n', '\n', '\n'}
				sut := NewInterpreter(data)
				sut.filter()
				sut.parseCommands()
				Expect(sut.commands).To(Exist)
				Expect(sut.commands.Len()).To(Equal, 1)
				Expect(sut.commands.Get(1)).To(Equal, NewCommand("exit"))
			})
			It("解析できないパターンができたときにエラーが作成されること", func() {
				data = []byte{'u', 'n', 'k', 'o', 'w', 'n', '\n', '\n', '\t'}
				sut := NewInterpreter(data)
				sut.filter()
				err := sut.parseCommands()
				Expect(sut.commands).To(NotExist)
				Expect(err).To(Exist)
			})
			It("足し算する命令が作成されること", func() {
				data = []byte{'a', 'd', 'd', '\t', ' ', ' ', ' '}
				sut := NewInterpreter(data)
				sut.filter()
				sut.parseCommands()
				Expect(sut.commands).To(Exist)
				Expect(sut.commands.Len()).To(Equal, 1)
				Expect(sut.commands.Get(1)).To(Equal, NewCommand("add"))
			})
			It("引き算する命令が作成されること", func() {
				data = []byte{'s', 'u', 'b', '\t', ' ', ' ', '\t'}
				sut := NewInterpreter(data)
				sut.filter()
				sut.parseCommands()
				Expect(sut.commands).To(Exist)
				Expect(sut.commands.Len()).To(Equal, 1)
				Expect(sut.commands.Get(1)).To(Equal, NewCommand("sub"))
			})
			It("掛け算する命令が作成されること", func() {
				data = []byte{'m', 'u', 'l', '\t', ' ', ' ', '\n'}
				sut := NewInterpreter(data)
				sut.filter()
				sut.parseCommands()
				Expect(sut.commands).To(Exist)
				Expect(sut.commands.Len()).To(Equal, 1)
				Expect(sut.commands.Get(1)).To(Equal, NewCommand("mul"))
			})
			It("割り算する命令が作成されること", func() {
				data = []byte{'d', 'i', 'v', '\t', ' ', '\t', ' '}
				sut := NewInterpreter(data)
				sut.filter()
				sut.parseCommands()
				Expect(sut.commands).To(Exist)
				Expect(sut.commands.Len()).To(Equal, 1)
				Expect(sut.commands.Get(1)).To(Equal, NewCommand("div"))
			})
			It("余りを求める命令が作成されること", func() {
				data = []byte{'m', 'o', 'd', '\t', ' ', '\t', '\t'}
				sut := NewInterpreter(data)
				sut.filter()
				sut.parseCommands()
				Expect(sut.commands).To(Exist)
				Expect(sut.commands.Len()).To(Equal, 1)
				Expect(sut.commands.Get(1)).To(Equal, NewCommand("mod"))
			})
			It("解析できないパターンができたときにエラーが作成されること", func() {
				data = []byte{'u', 'n', 'k', 'o', 'w', 'n', '\t', ' ', '\t', '\n'}
				sut := NewInterpreter(data)
				sut.filter()
				err := sut.parseCommands()
				Expect(sut.commands).To(NotExist)
				Expect(err).To(Exist)
			})
			It("ヒープエリアに値をプッシュする命令が作成されること", func() {
				data = []byte{'h', 'e', 'a', 'p', '\t', 'p', 'u', 's', 'h', '\t', ' '}
				sut := NewInterpreter(data)
				sut.filter()
				sut.parseCommands()
				Expect(sut.commands).To(Exist)
				Expect(sut.commands.Len()).To(Equal, 1)
				Expect(sut.commands.Get(1)).To(Equal, NewSubCommand("heap", "push"))
			})
			It("ヒープエリアから値をポップする命令が作成されること", func() {
				data = []byte{'h', 'e', 'a', 'p', '\t', 'p', 'o', 'p', '\t', '\t'}
				sut := NewInterpreter(data)
				sut.filter()
				sut.parseCommands()
				Expect(sut.commands).To(Exist)
				Expect(sut.commands.Len()).To(Equal, 1)
				Expect(sut.commands.Get(1)).To(Equal, NewSubCommand("heap", "pop"))
			})
			It("解析できないパターンができたときにエラーが作成されること", func() {
				data = []byte{'u', 'n', 'k', 'o', 'w', 'n', '\t', '\t', '\n'}
				sut := NewInterpreter(data)
				sut.filter()
				err := sut.parseCommands()
				Expect(sut.commands).To(NotExist)
				Expect(err).To(Exist)
			})
			It("スタックの内容を文字として標準出力する命令が作成されること", func() {
				data = []byte{'p', 'u', 't', 'c', '\t', '\n', ' ', ' '}
				sut := NewInterpreter(data)
				sut.filter()
				sut.parseCommands()
				Expect(sut.commands).To(Exist)
				Expect(sut.commands.Len()).To(Equal, 1)
				Expect(sut.commands.Get(1)).To(Equal, NewCommand("putc"))
			})
			It("スタックの内容を数字として標準出力する命令が作成されること", func() {
				data = []byte{'p', 'u', 't', 'n', '\t', '\n', ' ', '\t'}
				sut := NewInterpreter(data)
				sut.filter()
				sut.parseCommands()
				Expect(sut.commands).To(Exist)
				Expect(sut.commands.Len()).To(Equal, 1)
				Expect(sut.commands.Get(1)).To(Equal, NewCommand("putn"))
			})
			It("標準入力の値を文字としてスタックに格納する命令が作成されること", func() {
				data = []byte{'g', 'e', 't', 'c', '\t', '\n', '\t', ' '}
				sut := NewInterpreter(data)
				sut.filter()
				sut.parseCommands()
				Expect(sut.commands).To(Exist)
				Expect(sut.commands.Len()).To(Equal, 1)
				Expect(sut.commands.Get(1)).To(Equal, NewCommand("getc"))
			})
		})
	})
}

func current_dir() string {
	_, fpath, _, _ := runtime.Caller(0)
	return path.Dir(fpath)
}
