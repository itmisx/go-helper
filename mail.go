package helper

import (
	"gopkg.in/gomail.v2"
)

type mailer struct {
	Host     string
	Port     int
	Username string
	Password string
}

// NewMailer
func NewMailer(host string, port int, username string, password string) mailer {
	return mailer{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
	}
}

// SendEmail 发送邮件
// from 发送方
// to 接收方
// suject 主题
// body 内容
func (m mailer) SendEmail(
	from,
	to,
	subject,
	body string,
) error {
	dialer := gomail.NewDialer(
		m.Host,
		m.Port,
		m.Username,
		m.Password,
	)
	dialer.SSL = true
	msg := gomail.NewMessage()
	msg.SetHeader("From", from)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body)
	return dialer.DialAndSend(msg)
}
