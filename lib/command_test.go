package lib

import (
	"fmt"
	. "github.com/r7kamura/gospel"
	"testing"
)

func TestCommnad(t *testing.T) {
	Describe(t, "Command Tests", func() {
		Context("コマンドだけ指定する", func() {
			It("インスタンスが作成できること", func() {
				cmd := "cmd"
				sut := newCommand(cmd)
				Expect(sut).To(Exist)
			})
			It("コマンドだけ変数が書き換えられていること", func() {
				cmd := "cmd"
				sut := newCommand(cmd)
				Expect(sut.cmd).To(Equal, cmd)
				Expect(sut.subcmd).To(Equal, "")
				Expect(sut.param).To(Equal, 0)
			})
		})
		Context("サブコマンドを作成する", func() {
			It("インスタンスが作成できること", func() {
				cmd, subcmd := "cmd", "subcmd"
				Expect(newSubCommand(cmd, subcmd)).To(Exist)
			})
			It("コマンドとサブコマンドが指定した値で書き換えられていること", func() {
				cmd, subcmd := "cmd", "subcmd"
				sut := newSubCommand(cmd, subcmd)
				Expect(sut.cmd).To(Equal, cmd)
				Expect(sut.subcmd).To(Equal, subcmd)
				Expect(sut.param).To(Equal, 0)
			})
		})
		Context("パラメータ付きのコマンドを作成する", func() {
			It("インスタンスが作成できること", func() {
				cmd := "cmd"
				param := 1
				sut := newCommandWithParam(cmd, param)
				Expect(sut).To(Exist)
			})
			It("コマンドとパラメータが指定した値で書き換えられていること", func() {
				cmd := "cmd"
				param := 1
				sut := newCommandWithParam(cmd, param)
				Expect(sut.cmd).To(Equal, cmd)
				Expect(sut.subcmd).To(Equal, "")
				Expect(sut.param).To(Equal, 1)
			})
		})
		Context("パラメータ付きのサブコマンドを作成する", func() {
			It("インスタンスが作成できること", func() {
				cmd, subcmd := "cmd", "subcmd"
				param := 1
				sut := newSubCommandWithParam(cmd, subcmd, param)
				Expect(sut).To(Exist)
			})
			It("指定した値でメンバー変数が書き換えられていること", func() {
				cmd, subcmd := "cmd", "subcmd"
				param := 1
				sut := newSubCommandWithParam(cmd, subcmd, param)
				Expect(sut.cmd).To(Equal, cmd)
				Expect(sut.subcmd).To(Equal, subcmd)
				Expect(sut.param).To(Equal, param)
			})
		})
		Context("文字列に変換する", func() {
			It("指定したフォーマットになっている", func() {
				cmd, subcmd := "cmd", "subcmd"
				param := 1
				sut := newSubCommandWithParam(cmd, subcmd, param)
				Expect(fmt.Sprint(sut)).To(Equal, "cmd subcmd 1")
			})
			It("パラメータに２を渡したとき指定したフォーマットになっている", func() {
				cmd, subcmd := "cmd", "subcmd"
				param := 2
				sut := newSubCommandWithParam(cmd, subcmd, param)
				Expect(fmt.Sprint(sut)).To(Equal, "cmd subcmd 2")
			})
		})
	})
}

func TestCommandList(t *testing.T) {
	Describe(t, "CommandList Tests", func() {
		Context("実行するコマンドの一覧を作成する", func() {
			It("インスタンスが作成できること", func() {
				sut := newCommandList()
				Expect(sut).To(Exist)
			})
		})
		Context("リストにコマンドを追加する関数", func() {
			sut := newCommandList()
			It("コマンドが追加できることを確認する", func() {
				sut.Add(newCommand("test"))
				Expect(sut.Len()).To(Equal, 1)
			})
			It("コマンドが追加できることを確認する（2回目）", func() {
				sut.Add(newCommand("test2"))
				Expect(sut.Len()).To(Equal, 2)
			})
		})
		Context("コマンドをすべて削除する関数", func() {
			sut := newCommandList()
			sut.Add(newCommand("test"))
			Expect(sut.Len()).To(Equal, 1)
			It("コマンドがすべて削除される", func() {
				sut.Clear()
				Expect(sut.Len()).To(Equal, 0)
			})
			It("コマンドがすべて削除される", func() {
				sut.Add(newCommand("test"))
				sut.Add(newCommand("test2"))
				sut.Clear()
				Expect(sut.Len()).To(Equal, 0)
			})
		})
		Context("コマンドを行番号で取得する関数", func() {
			sut := newCommandList()
			sut.Add(newCommand("test"))
			sut.Add(newCommand("test2"))
			It("指定した１行目のコマンドを取得できること", func() {
				actual := sut.Get(1)
				Expect(actual).To(Exist)
				Expect(actual.cmd).To(Equal, "test")
			})
			It("指定した２行目のコマンドを取得できること", func() {
				actual := sut.Get(2)
				Expect(actual).To(Exist)
				Expect(actual.cmd).To(Equal, "test2")
			})
			It("指定した3行目のコマンドが存在しないときにnilが返ってくること", func() {
				actual := sut.Get(3)
				Expect(actual).To(NotExist)
			})
			It("指定した０行目のコマンドが存在しないときにnilが返ってくること", func() {
				actual := sut.Get(0)
				Expect(actual).To(NotExist)
			})
			It("０未満の値を指定したときにnilが返ってくること", func() {
				actual := sut.Get(-1)
				Expect(actual).To(NotExist)
				actual = sut.Get(-2)
				Expect(actual).To(NotExist)
			})
		})
		Context("コマンドリストから目的のコマンドを見つける", func() {
			sut := newCommandList()
			sut.Add(newCommand("test"))
			sut.Add(newCommand("test2"))
			sut.Add(newCommand("test3"))
			It("コマンドリストから２番目のコマンドのキーを取得できること", func() {
				key, err := sut.Search(newCommand("test2"))
				Expect(err).To(NotExist)
				Expect(key).To(Equal, 2)
			})
			It("コマンドリストから３番目のコマンドのキーを取得できること", func() {
				key, err := sut.Search(newCommand("test3"))
				Expect(err).To(NotExist)
				Expect(key).To(Equal, 3)
			})
			It("コマンドリストから１番目のコマンドのキーを取得できること", func() {
				key, err := sut.Search(newCommand("test"))
				Expect(err).To(NotExist)
				Expect(key).To(Equal, 1)
			})
			It("存在しないコマンドを指定されたときエラーが発生すること", func() {
				key, err := sut.Search(newCommand("not defined"))
				Expect(err).To(Exist)
				Expect(key).To(Equal, -1)
			})
		})
	})
}
