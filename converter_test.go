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
		Context("命令作成関数の生成", func() {
			It("スタックを操作する命令を作成する関数が作成されること", func() {
				data := byte(' ')
				sut := NewConverter()
				fn, err := sut.CreateFunction(data)
				Expect(err).To(NotExist)
				Expect(fn).To(Exist)
				// 目的のメソッドが生成されているか確認する処理を作る必要がある
				// 下記のコードだと同じということがチェックできない
				//Expect(fn).To(Equal, sut.stackManipulation)
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
			It("スタックをコピーするコマンドが作成されること", func() {
				data := []byte{'\n', ' '}
				sut := NewConverter()
				cmd, seek, err := sut.stackManipulation(data)
				Expect(err).To(NotExist)
				Expect(seek).To(Equal, len(data))
				Expect(cmd).To(Exist)
				Expect(cmd).To(Equal, NewSubCommand("stack", "copy"))
			})
			It("スタックのトップと２番目を入れ替えるコマンドが作成されること", func() {
				data := []byte{'\n', '\t'}
				sut := NewConverter()
				cmd, seek, err := sut.stackManipulation(data)
				Expect(err).To(NotExist)
				Expect(seek).To(Equal, len(data))
				Expect(cmd).To(Exist)
				Expect(cmd).To(Equal, NewSubCommand("stack", "swap"))
			})
			It("スタックのトップを削除するコマンドが作成されること", func() {
				data := []byte{'\n', '\n'}
				sut := NewConverter()
				cmd, seek, err := sut.stackManipulation(data)
				Expect(err).To(NotExist)
				Expect(seek).To(Equal, len(data))
				Expect(cmd).To(Exist)
				Expect(cmd).To(Equal, NewSubCommand("stack", "remove"))
			})
			It("定義されていない命令が指定されたときにundefinedの命令が作成されること", func() {
				data := []byte{'\t', '\n'}
				sut := NewConverter()
				cmd, seek, err := sut.stackManipulation(data)
				Expect(err).To(Exist)
				Expect(seek).To(Equal, 0)
				Expect(cmd).To(NotExist)
			})
		})
		Context("制御文に関する命令の生成", func() {
			It("ラベルを定義するコマンドが作成されること", func() {
				data := []byte{' ', ' ', '\t', ' ', ' ', '\t', '\n'}
				sut := NewConverter()
				cmd, seek, err := sut.flowControl(data)
				Expect(err).To(NotExist)
				Expect(seek).To(Equal, len(data))
				Expect(cmd).To(Exist)
				Expect(cmd).To(Equal, NewSubCommand("label", "1001"))
			})
			It("ラベルを呼び出すコマンドが作成されること", func() {
				data := []byte{' ', '\t', '\t', ' ', ' ', '\t', '\n'}
				sut := NewConverter()
				cmd, seek, err := sut.flowControl(data)
				Expect(err).To(NotExist)
				Expect(seek).To(Equal, len(data))
				Expect(cmd).To(Exist)
				Expect(cmd).To(Equal, NewSubCommand("call", "1001"))
			})
			It("ラベルを呼び出すコマンドが作成されること", func() {
				data := []byte{' ', '\n', '\t', ' ', ' ', '\t', '\n'}
				sut := NewConverter()
				cmd, seek, err := sut.flowControl(data)
				Expect(err).To(NotExist)
				Expect(seek).To(Equal, len(data))
				Expect(cmd).To(Exist)
				Expect(cmd).To(Equal, NewSubCommand("goto", "1001"))
			})
			It("スタックの値が０のときにラベルを呼び出すコマンドが作成されること", func() {
				data := []byte{'\t', ' ', '\t', ' ', ' ', '\t', '\n'}
				sut := NewConverter()
				cmd, seek, err := sut.flowControl(data)
				Expect(err).To(NotExist)
				Expect(seek).To(Equal, len(data))
				Expect(cmd).To(Exist)
				Expect(cmd).To(Equal, NewSubCommand("if stack==0 then goto", "1001"))
			})
			It("スタックの値が０未満のときにラベルを呼び出すコマンドが作成されること", func() {
				data := []byte{'\t', '\t', '\t', ' ', ' ', '\t', '\n'}
				sut := NewConverter()
				cmd, seek, err := sut.flowControl(data)
				Expect(err).To(NotExist)
				Expect(seek).To(Equal, len(data))
				Expect(cmd).To(Exist)
				Expect(cmd).To(Equal, NewSubCommand("if stack!=0 then goto", "1001"))
			})
			It("呼び出し元に戻るコマンドが作成されること", func() {
				data := []byte{'\t', '\n'}
				sut := NewConverter()
				cmd, seek, err := sut.flowControl(data)
				Expect(err).To(NotExist)
				Expect(seek).To(Equal, len(data))
				Expect(cmd).To(Exist)
				Expect(cmd).To(Equal, NewCommand("return"))
			})
			It("プログラムを終了するコマンドが作成されること", func() {
				data := []byte{'\n', '\n'}
				sut := NewConverter()
				cmd, seek, err := sut.flowControl(data)
				Expect(err).To(NotExist)
				Expect(seek).To(Equal, len(data))
				Expect(cmd).To(Exist)
				Expect(cmd).To(Equal, NewCommand("exit"))
			})
			It("解析できないパターンができたときにエラーが作成されること", func() {
				data := []byte{'\n', '\t'}
				sut := NewConverter()
				cmd, seek, err := sut.flowControl(data)
				Expect(err).To(Exist)
				Expect(seek).To(Equal, 0)
				Expect(cmd).To(NotExist)
			})
		})
		Context("演算の命令を作成", func() {
			It("足し算する命令が作成されること", func() {
				data := []byte{' ', ' '}
				sut := NewConverter()
				cmd, seek, err := sut.arithmetic(data)
				Expect(err).To(NotExist)
				Expect(seek).To(Equal, len(data))
				Expect(cmd).To(Exist)
				Expect(cmd).To(Equal, NewCommand("add"))
			})
			It("引き算する命令が作成されること", func() {
				data := []byte{' ', '\t'}
				sut := NewConverter()
				cmd, seek, err := sut.arithmetic(data)
				Expect(err).To(NotExist)
				Expect(seek).To(Equal, len(data))
				Expect(cmd).To(Exist)
				Expect(cmd).To(Equal, NewCommand("sub"))
			})
			It("掛け算する命令が作成されること", func() {
				data := []byte{' ', '\n'}
				sut := NewConverter()
				cmd, seek, err := sut.arithmetic(data)
				Expect(err).To(NotExist)
				Expect(seek).To(Equal, len(data))
				Expect(cmd).To(Exist)
				Expect(cmd).To(Equal, NewCommand("mul"))
			})
			It("割り算する命令が作成されること", func() {
				data := []byte{'\t', ' '}
				sut := NewConverter()
				cmd, seek, err := sut.arithmetic(data)
				Expect(err).To(NotExist)
				Expect(seek).To(Equal, len(data))
				Expect(cmd).To(Exist)
				Expect(cmd).To(Equal, NewCommand("div"))
			})
			It("余りを求める命令が作成されること", func() {
				data := []byte{'\t', '\t'}
				sut := NewConverter()
				cmd, seek, err := sut.arithmetic(data)
				Expect(err).To(NotExist)
				Expect(seek).To(Equal, len(data))
				Expect(cmd).To(Exist)
				Expect(cmd).To(Equal, NewCommand("mod"))
			})
			It("解析できないパターンができたときにエラーが作成されること", func() {
				data := []byte{'\t', '\n'}
				sut := NewConverter()
				cmd, seek, err := sut.arithmetic(data)
				Expect(err).To(Exist)
				Expect(seek).To(Equal, 0)
				Expect(cmd).To(NotExist)
			})
		})
		Context("ヒープ領域を操作する命令の作成", func() {
			It("ヒープ領域にスタックトップの値を保存する命令が作成されること", func() {
				data := []byte{' '}
				sut := NewConverter()
				cmd, seek, err := sut.heapAccess(data)
				Expect(err).To(NotExist)
				Expect(seek).To(Equal, len(data))
				Expect(cmd).To(Exist)
				Expect(cmd).To(Equal, NewSubCommand("heap", "push"))
			})
			It("ヒープ領域あら値を取得してスタック領域に保存する命令が作成されること", func() {
				data := []byte{'\t'}
				sut := NewConverter()
				cmd, seek, err := sut.heapAccess(data)
				Expect(err).To(NotExist)
				Expect(seek).To(Equal, len(data))
				Expect(cmd).To(Exist)
				Expect(cmd).To(Equal, NewSubCommand("heap", "pop"))
			})
			It("解析できないパターンができたときにエラーが作成されること", func() {
				data := []byte{'\n'}
				sut := NewConverter()
				cmd, seek, err := sut.heapAccess(data)
				Expect(err).To(Exist)
				Expect(seek).To(Equal, 0)
				Expect(cmd).To(NotExist)
			})
		})
		Context("I/O操作に関する命令の作成", func() {
			It("スタックの内容を文字として標準出力する命令が作成されること", func() {
				data := []byte{' ', ' '}
				sut := NewConverter()
				cmd, seek, err := sut.i_o(data)
				Expect(err).To(NotExist)
				Expect(seek).To(Equal, len(data))
				Expect(cmd).To(Exist)
				Expect(cmd).To(Equal, NewCommand("putc"))
			})
			It("スタックの内容を数字として標準出力する命令が作成されること", func() {
				data := []byte{' ', '\t'}
				sut := NewConverter()
				cmd, seek, err := sut.i_o(data)
				Expect(err).To(NotExist)
				Expect(seek).To(Equal, len(data))
				Expect(cmd).To(Exist)
				Expect(cmd).To(Equal, NewCommand("putn"))
			})
			It("標準入力の値を文字としてスタックに格納する命令が作成されること", func() {
				data := []byte{'\t', ' '}
				sut := NewConverter()
				cmd, seek, err := sut.i_o(data)
				Expect(err).To(NotExist)
				Expect(seek).To(Equal, len(data))
				Expect(cmd).To(Exist)
				Expect(cmd).To(Equal, NewCommand("getc"))
			})
			It("標準入力の値を数値としてスタックに格納する命令が作成されること", func() {
				data := []byte{'\t', '\t'}
				sut := NewConverter()
				cmd, seek, err := sut.i_o(data)
				Expect(err).To(NotExist)
				Expect(seek).To(Equal, len(data))
				Expect(cmd).To(Exist)
				Expect(cmd).To(Equal, NewCommand("getn"))
			})
			It("解析できないパターンができたときにエラーが作成されること", func() {
				data := []byte{'\t', '\n'}
				sut := NewConverter()
				cmd, seek, err := sut.i_o(data)
				Expect(err).To(Exist)
				Expect(seek).To(Equal, 0)
				Expect(cmd).To(NotExist)
			})
		})
	})
}
