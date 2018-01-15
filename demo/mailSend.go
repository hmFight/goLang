package demo

import (
	"fmt"
	"gopkg.in/gomail.v2"
)

//邮件发送
//doc https://github.com/go-gomail/gomail
func SendMail() {
	sender := "sned@dsds.com"
	senderNickName := "wangjia"
	senderPw := "password"
	to := "4355to@qq.com"
	host := "smtp.exmail.qq.com"

	subject := "使用Golang发送邮件"

	m := gomail.NewMessage()
	// 发件人
	m.SetAddressHeader("From", sender, senderNickName)
	// 收件人
	m.SetHeader("To",
		m.FormatAddress(to, "wdado"),
	)
	m.SetHeader("Subject", subject)                                                             // 主题
	m.SetBody("text/html", "Hello <a href = \"http://blog.csdn.net/liang19890820\">hdahds</a>") // 正文

	d := gomail.NewDialer(host, 25, sender, senderPw) // 发送邮件服务器、端口、发件人账号、发件人密码
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	} else {
		fmt.Println("send ok")
	}
}
