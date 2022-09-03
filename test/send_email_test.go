package test

import (
	"quora/common/tool"
	"testing"
)

func TestSendCodeByEmail(t *testing.T) {
	tool.SendCodeByEmail("1725014728@qq.com", "asdasd", tool.Email{
		ServerEmail: "1725014728@qq.com",
		PassWord:    "XXXXXXXXX",
	})
}
