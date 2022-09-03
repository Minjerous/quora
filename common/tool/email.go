package tool

import (
	"crypto/tls"
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

type Email struct {
	ServerEmail string
	PassWord    string
}

func SendCodeByEmail(Email, code string, emailCfg Email) error {

	e := email.NewEmail()
	e.From = fmt.Sprintf("Get <%s>", emailCfg.ServerEmail)
	e.To = []string{Email}
	e.Subject = "Quora邮箱验证码"
	e.HTML = []byte("你的验证码为：<h1>" + code + "</h1>" + "<p>请在5分钟内完成验证")
	err := e.SendWithTLS("smtp.qq.com:465", smtp.PlainAuth("", emailCfg.ServerEmail, emailCfg.PassWord, "smtp.qq.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.qq.com"})
	return err
}
