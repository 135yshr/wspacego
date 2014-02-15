package lib

import (
	. "github.com/r7kamura/gospel"
	"testing"
)

func TestConerter(t *testing.T) {
	Describe(t, "whitespace のソースを文字や読みやすい文字列に変換する", func() {
		Context("命令作成関数の生成", func() {
			It("スタックを操作する命令を生成する関数が作成されること", func() {
				data := byte(' ')
				fn, err := createFunction(data)
				Expect(err).To(NotExist)
				Expect(fn).To(Exist)
				// TODO:目的のメソッドが生成されているか確認する処理を作る必要がある
				//      下記のコードだと同じということがチェックできない
				//Expect(fn).To(Equal, stackManipulation)
			})
			It("制御文を作成する関数が生成されること", func() {
				data := byte('\n')
				fn, err := createFunction(data)
				Expect(err).To(NotExist)
				Expect(fn).To(Exist)
				// TODO:目的のメソッドが生成されているか確認する処理を作る必要がある
				//      下記のコードだと同じということがチェックできない
				//Expect(fn).To(Equal, stackManipulation)
			})
			It("演算やヒープ領域の操作と入出力に関する命令を作成する関数が生成されること", func() {
				data := byte('\t')
				fn, err := createFunction(data)
				Expect(err).To(NotExist)
				Expect(fn).To(Exist)
				// TODO:目的のメソッドが生成されているか確認する処理を作る必要がある
				//      下記のコードだと同じということがチェックできない
				//Expect(fn).To(Equal, stackManipulation)
			})
			It("解析できないパターンができたときにエラーが作成されること", func() {
				data := byte('A')
				fn, err := createFunction(data)
				Expect(err).To(Exist)
				Expect(fn).To(NotExist)
			})
		})
		Context("スタックに関連する命令の生成", func() {
			It("スタックに１をプッシュするコマンドが作成されること", func() {
				data := []byte{' ', '\t', '\n'}
				cmd, seek, err := stackManipulation(data)
				Expect(err).To(NotExist)
				Expect(seek).To(Equal, len(data))
				Expect(cmd).To(Exist)
				Expect(cmd).To(Equal, newSubCommandWithParam("stack", "push", 1))
			})
			It("スタックに2をプッシュするコマンドが作成されること", func() {
				data := []byte{' ', '\t', ' ', '\n'}
				cmd, seek, err := stackManipulation(data)
				Expect(err).To(NotExist)
				Expect(seek).To(Equal, len(data))
				Expect(cmd).To(Exist)
				Expect(cmd).To(Equal, newSubCommandWithParam("stack", "push", 2))
			})
			It("スタックに4をプッシュするコマンドが作成されること", func() {
				data := []byte{' ', '\t', ' ', ' ', '\n'}
				cmd, seek, err := stackManipulation(data)
				Expect(err).To(NotExist)
				Expect(seek).To(Equal, len(data))
				Expect(cmd).To(Exist)
				Expect(cmd).To(Equal, newSubCommandWithParam("stack", "push", 4))
			})
			It("スタックをコピーするコマンドが作成されること", func() {
				data := []byte{'\n', ' '}
				cmd, seek, err := stackManipulation(data)
				Expect(err).To(NotExist)
				Expect(seek).To(Equal, len(data))
				Expect(cmd).To(Exist)
				Expect(cmd).To(Equal, newSubCommand("stack", "copy"))
			})
			It("スタックのトップと２番目を入れ替えるコマンドが作成されること", func() {
				data := []byte{'\n', '\t'}
				cmd, seek, err := stackManipulation(data)
				Expect(err).To(NotExist)
				Expect(seek).To(Equal, len(data))
				Expect(cmd).To(Exist)
				Expect(cmd).To(Equal, newSubCommand("stack", "swap"))
			})
			It("スタックのトップを削除するコマンドが作成されること", func() {
				data := []byte{'\n', '\n'}
				cmd, seek, err := stackManipulation(data)
				Expect(err).To(NotExist)
				Expect(seek).To(Equal, len(data))
				Expect(cmd).To(Exist)
				Expect(cmd).To(Equal, newSubCommand("stack", "remove"))
			})
			It("定義されていない命令が指定されたときにundefinedの命令が作成されること", func() {
				data := []byte{'\t', '\n'}
				cmd, seek, err := stackManipulation(data)
				Expect(err).To(Exist)
				Expect(seek).To(Equal, 0)
				Expect(cmd).To(NotExist)
			})
		})
		Context("制御文に関する命令の生成", func() {
			It("ラベルを定義するコマンドが作成されること", func() {
				data := []byte{' ', ' ', '\t', ' ', ' ', '\t', '\n'}
				cmd, seek, err := flowControl(data)
				Expect(err).To(NotExist)
				Expect(seek).To(Equal, len(data))
				Expect(cmd).To(Exist)
				Expect(cmd).To(Equal, newSubCommand("label", "1001"))
			})
			It("ラベルを呼び出すコマンドが作成されること", func() {
				data := []byte{' ', '\t', '\t', ' ', ' ', '\t', '\n'}
				cmd, seek, err := flowControl(data)
				Expect(err).To(NotExist)
				Expect(seek).To(Equal, len(data))
				Expect(cmd).To(Exist)
				Expect(cmd).To(Equal, newSubCommand("call", "1001"))
			})
			It("ラベルを呼び出すコマンドが作成されること", func() {
				data := []byte{' ', '\n', '\t', ' ', ' ', '\t', '\n'}
				cmd, seek, err := flowControl(data)
				Expect(err).To(NotExist)
				Expect(seek).To(Equal, len(data))
				Expect(cmd).To(Exist)
				Expect(cmd).To(Equal, newSubCommand("goto", "1001"))
			})
			It("スタックの値が０のときにラベルを呼び出すコマンドが作成されること", func() {
				data := []byte{'\t', ' ', '\t', ' ', ' ', '\t', '\n'}
				cmd, seek, err := flowControl(data)
				Expect(err).To(NotExist)
				Expect(seek).To(Equal, len(data))
				Expect(cmd).To(Exist)
				Expect(cmd).To(Equal, newSubCommand("if stack==0 then goto", "1001"))
			})
			It("スタックの値が０未満のときにラベルを呼び出すコマンドが作成されること", func() {
				data := []byte{'\t', '\t', '\t', ' ', ' ', '\t', '\n'}
				cmd, seek, err := flowControl(data)
				Expect(err).To(NotExist)
				Expect(seek).To(Equal, len(data))
				Expect(cmd).To(Exist)
				Expect(cmd).To(Equal, newSubCommand("if stack!=0 then goto", "1001"))
			})
			It("呼び出し元に戻るコマンドが作成されること", func() {
				data := []byte{'\t', '\n'}
				cmd, seek, err := flowControl(data)
				Expect(err).To(NotExist)
				Expect(seek).To(Equal, len(data))
				Expect(cmd).To(Exist)
				Expect(cmd).To(Equal, newCommand("return"))
			})
			It("プログラムを終了するコマンドが作成されること", func() {
				data := []byte{'\n', '\n'}
				cmd, seek, err := flowControl(data)
				Expect(err).To(NotExist)
				Expect(seek).To(Equal, len(data))
				Expect(cmd).To(Exist)
				Expect(cmd).To(Equal, newCommand("exit"))
			})
			It("解析できないパターンができたときにエラーが作成されること", func() {
				data := []byte{'\n', '\t'}
				cmd, seek, err := flowControl(data)
				Expect(err).To(Exist)
				Expect(seek).To(Equal, 0)
				Expect(cmd).To(NotExist)
			})
		})
		Context("演算の命令を作成", func() {
			It("足し算する命令が作成されること", func() {
				data := []byte{' ', ' '}
				cmd, seek, err := arithmetic(data)
				Expect(err).To(NotExist)
				Expect(seek).To(Equal, len(data))
				Expect(cmd).To(Exist)
				Expect(cmd).To(Equal, newCommand("add"))
			})
			It("引き算する命令が作成されること", func() {
				data := []byte{' ', '\t'}
				cmd, seek, err := arithmetic(data)
				Expect(err).To(NotExist)
				Expect(seek).To(Equal, len(data))
				Expect(cmd).To(Exist)
				Expect(cmd).To(Equal, newCommand("sub"))
			})
			It("掛け算する命令が作成されること", func() {
				data := []byte{' ', '\n'}
				cmd, seek, err := arithmetic(data)
				Expect(err).To(NotExist)
				Expect(seek).To(Equal, len(data))
				Expect(cmd).To(Exist)
				Expect(cmd).To(Equal, newCommand("mul"))
			})
			It("割り算する命令が作成されること", func() {
				data := []byte{'\t', ' '}
				cmd, seek, err := arithmetic(data)
				Expect(err).To(NotExist)
				Expect(seek).To(Equal, len(data))
				Expect(cmd).To(Exist)
				Expect(cmd).To(Equal, newCommand("div"))
			})
			It("余りを求める命令が作成されること", func() {
				data := []byte{'\t', '\t'}
				cmd, seek, err := arithmetic(data)
				Expect(err).To(NotExist)
				Expect(seek).To(Equal, len(data))
				Expect(cmd).To(Exist)
				Expect(cmd).To(Equal, newCommand("mod"))
			})
			It("解析できないパターンができたときにエラーが作成されること", func() {
				data := []byte{'\t', '\n'}
				cmd, seek, err := arithmetic(data)
				Expect(err).To(Exist)
				Expect(seek).To(Equal, 0)
				Expect(cmd).To(NotExist)
			})
		})
		Context("ヒープ領域を操作する命令の作成", func() {
			It("ヒープ領域にスタックトップの値を保存する命令が作成されること", func() {
				data := []byte{' '}
				cmd, seek, err := heapAccess(data)
				Expect(err).To(NotExist)
				Expect(seek).To(Equal, len(data))
				Expect(cmd).To(Exist)
				Expect(cmd).To(Equal, newSubCommand("heap", "push"))
			})
			It("ヒープ領域あら値を取得してスタック領域に保存する命令が作成されること", func() {
				data := []byte{'\t'}
				cmd, seek, err := heapAccess(data)
				Expect(err).To(NotExist)
				Expect(seek).To(Equal, len(data))
				Expect(cmd).To(Exist)
				Expect(cmd).To(Equal, newSubCommand("heap", "pop"))
			})
			It("解析できないパターンができたときにエラーが作成されること", func() {
				data := []byte{'\n'}
				cmd, seek, err := heapAccess(data)
				Expect(err).To(Exist)
				Expect(seek).To(Equal, 0)
				Expect(cmd).To(NotExist)
			})
		})
		Context("I/O操作に関する命令の作成", func() {
			It("スタックの内容を文字として標準出力する命令が作成されること", func() {
				data := []byte{' ', ' '}
				cmd, seek, err := i_o(data)
				Expect(err).To(NotExist)
				Expect(seek).To(Equal, len(data))
				Expect(cmd).To(Exist)
				Expect(cmd).To(Equal, newCommand("putc"))
			})
			It("スタックの内容を数字として標準出力する命令が作成されること", func() {
				data := []byte{' ', '\t'}
				cmd, seek, err := i_o(data)
				Expect(err).To(NotExist)
				Expect(seek).To(Equal, len(data))
				Expect(cmd).To(Exist)
				Expect(cmd).To(Equal, newCommand("putn"))
			})
			It("標準入力の値を文字としてスタックに格納する命令が作成されること", func() {
				data := []byte{'\t', ' '}
				cmd, seek, err := i_o(data)
				Expect(err).To(NotExist)
				Expect(seek).To(Equal, len(data))
				Expect(cmd).To(Exist)
				Expect(cmd).To(Equal, newCommand("getc"))
			})
			It("標準入力の値を数値としてスタックに格納する命令が作成されること", func() {
				data := []byte{'\t', '\t'}
				cmd, seek, err := i_o(data)
				Expect(err).To(NotExist)
				Expect(seek).To(Equal, len(data))
				Expect(cmd).To(Exist)
				Expect(cmd).To(Equal, newCommand("getn"))
			})
			It("解析できないパターンができたときにエラーが作成されること", func() {
				data := []byte{'\t', '\n'}
				cmd, seek, err := i_o(data)
				Expect(err).To(Exist)
				Expect(seek).To(Equal, 0)
				Expect(cmd).To(NotExist)
			})
		})
	})
}
