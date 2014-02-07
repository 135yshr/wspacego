package wspacego

import (
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
			})
		})
	})
}