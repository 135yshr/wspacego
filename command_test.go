package wspacego

import (
	. "github.com/r7kamura/gospel"
	"testing"
)

func TestCommnad(t *testing.T) {
	Describe(t, "Command Tests", func() {
		Context("Create Instance", func() {
			It("is exists", func() {
				cmd, subcmd := "cmd", "subcmd"
				Expect(NewCommand(cmd, subcmd)).To(Exist)
			})
		})
	})
}
