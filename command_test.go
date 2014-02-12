package wspacego

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
				sut := NewCommand(cmd)
				Expect(sut).To(Exist)
			})
			It("コマンドだけ変数が書き換えられていること", func() {
				cmd := "cmd"
				sut := NewCommand(cmd)
				Expect(sut.cmd).To(Equal, cmd)
				Expect(sut.subcmd).To(Equal, "")
				Expect(sut.param).To(Equal, 0)
			})
		})
		Context("サブコマンドを作成する", func() {
			It("インスタンスが作成できること", func() {
				cmd, subcmd := "cmd", "subcmd"
				Expect(NewSubCommand(cmd, subcmd)).To(Exist)
			})
			It("コマンドとサブコマンドが指定した値で書き換えられていること", func() {
				cmd, subcmd := "cmd", "subcmd"
				sut := NewSubCommand(cmd, subcmd)
				Expect(sut.cmd).To(Equal, cmd)
				Expect(sut.subcmd).To(Equal, subcmd)
				Expect(sut.param).To(Equal, 0)
			})
		})
		Context("パラメータ付きのコマンドを作成する", func() {
			It("インスタンスが作成できること", func() {
				cmd := "cmd"
				param := 1
				sut := NewCommandWithParam(cmd, param)
				Expect(sut).To(Exist)
			})
			It("コマンドとパラメータが指定した値で書き換えられていること", func() {
				cmd := "cmd"
				param := 1
				sut := NewCommandWithParam(cmd, param)
				Expect(sut.cmd).To(Equal, cmd)
				Expect(sut.subcmd).To(Equal, "")
				Expect(sut.param).To(Equal, 1)
			})
		})
		Context("パラメータ付きのサブコマンドを作成する", func() {
			It("インスタンスが作成できること", func() {
				cmd, subcmd := "cmd", "subcmd"
				param := 1
				sut := NewSubCommandWithParam(cmd, subcmd, param)
				Expect(sut).To(Exist)
			})
			It("指定した値でメンバー変数が書き換えられていること", func() {
				cmd, subcmd := "cmd", "subcmd"
				param := 1
				sut := NewSubCommandWithParam(cmd, subcmd, param)
				Expect(sut.cmd).To(Equal, cmd)
				Expect(sut.subcmd).To(Equal, subcmd)
				Expect(sut.param).To(Equal, param)
			})
		})
		Context("文字列に変換する", func() {
			It("指定したフォーマットになっている", func() {
				cmd, subcmd := "cmd", "subcmd"
				param := 1
				sut := NewSubCommandWithParam(cmd, subcmd, param)
				Expect(fmt.Sprint(sut)).To(Equal, "cmd subcmd 1")
			})
			It("パラメータに２を渡したとき指定したフォーマットになっている", func() {
				cmd, subcmd := "cmd", "subcmd"
				param := 2
				sut := NewSubCommandWithParam(cmd, subcmd, param)
				Expect(fmt.Sprint(sut)).To(Equal, "cmd subcmd 2")
			})
		})
	})
}

func TestCommandList(t *testing.T) {
	Describe(t, "CommandList Tests", func() {
		Context("実行するコマンドの一覧を作成する", func() {
			It("インスタンスが作成できること", func() {
				sut := NewCommandList()
				Expect(sut).To(Exist)
			})
		})
		Context("リストにコマンドを追加する関数", func() {
			sut := NewCommandList()
			It("コマンドが追加できることを確認する", func() {
				sut.Add(NewCommand("test"))
				Expect(len(*sut)).To(Equal, 1)
			})
			It("コマンドが追加できることを確認する（2回目）", func() {
				sut.Add(NewCommand("test2"))
				Expect(len(*sut)).To(Equal, 2)
			})
		})
		Context("コマンドをすべて削除する関数", func() {
			sut := NewCommandList()
			sut.Add(NewCommand("test"))
			Expect(len(*sut)).To(Equal, 1)
			It("コマンドがすべて削除される", func() {
				sut.Clear()
				Expect(len(*sut)).To(Equal, 0)
			})
			It("コマンドがすべて削除される", func() {
				sut.Add(NewCommand("test"))
				sut.Add(NewCommand("test2"))
				sut.Clear()
				Expect(len(*sut)).To(Equal, 0)
			})
		})
		Context("コマンドを行番号で取得する関数", func() {
			sut := NewCommandList()
			sut.Add(NewCommand("test"))
			It("指定した１行目のコマンドを取得できること", func() {
				actual := sut.Get(1)
				Expect(actual).To(Exist)
			})
		})
	})
}
