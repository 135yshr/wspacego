package wspacego

import (
	. "github.com/r7kamura/gospel"
	"testing"
)

func TestCommnad(t *testing.T) {
	Describe(t, "Command Tests", func() {
		Context("Create Instance", func() {
			It("Initialize Instance", func() {
				cmd := "cmd"
				sut := NewCommand(cmd)
				Expect(sut).To(Exist)
			})
		})
		Context("Create Instance", func() {
			It("is exists", func() {
				cmd, subcmd := "cmd", "subcmd"
				Expect(NewSubCommand(cmd, subcmd)).To(Exist)
			})
		})
		Context("Create Instance", func() {
			It("Initialize Instance", func() {
				cmd, subcmd := "cmd", "subcmd"
				sut := NewSubCommand(cmd, subcmd)
				Expect(sut.cmd).To(Equal, cmd)
				Expect(sut.subcmd).To(Equal, subcmd)
				Expect(sut.param).To(Equal, 0)
			})
		})
	})
}
