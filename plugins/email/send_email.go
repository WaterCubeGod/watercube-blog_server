package email

import (
	"gopkg.in/gomail.v2"
	"gvb_server/global"
)

type Subject string

const (
	Code  Subject = "平台验证码"
	Note  Subject = "操作通知"
	Alarm Subject = "告警通知"
)

type EmailApi struct {
	Subject Subject
}

func (a *EmailApi) SendEmail(name, body string) error {
	return send(name, string(a.Subject), body)
}

func NewCode() EmailApi {
	return EmailApi{
		Subject: Code,
	}
}

func NewNote() EmailApi {
	return EmailApi{
		Subject: Note,
	}
}

func NewAlarm() EmailApi {
	return EmailApi{
		Subject: Alarm,
	}
}

func send(name, subject, body string) error {
	e := global.CONFIG.Email
	return sendMail(
		e.User,
		e.Password,
		e.Host,
		e.Port,
		name,
		e.DefaultFromEmail,
		subject,
		body,
	)
}

func sendMail(userName, authCode, host string, port int, mailTo, sendName string, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(userName, sendName)) // 谁发的
	m.SetHeader("To", mailTo)                                // 发给谁
	m.SetHeader("Subject", subject)                          // 主题
	m.SetBody("text/html", body)
	d := gomail.NewDialer(host, port, userName, authCode)
	err := d.DialAndSend(m)
	return err
}
