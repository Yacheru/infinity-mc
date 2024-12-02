package email

import (
	"net/smtp"
	"net/textproto"

	"github.com/jordan-wright/email"
)

type Sender interface {
	SendMail(to []string, msg []byte, subject string) error
}

type SMTP struct {
	auth smtp.Auth

	sender         string
	senderPassword string

	SMTPServer string
	SMTPPort   string
}

func NewSMPTServer(sender string, senderPassword string, smtpServer string, smtpPort string) *SMTP {
	auth := smtp.PlainAuth("", sender, senderPassword, smtpServer)

	return &SMTP{
		auth:           auth,
		sender:         sender,
		senderPassword: senderPassword,
		SMTPServer:     smtpServer,
		SMTPPort:       smtpPort,
	}
}

func (s *SMTP) SendMail(to []string, msg []byte, subject string) error {
	e := email.Email{
		To:      to,
		From:    s.sender,
		HTML:    msg,
		Subject: subject,
		Headers: textproto.MIMEHeader{},
	}
	addr := s.SMTPServer + ":" + s.SMTPPort

	err := e.Send(addr, s.auth)
	if err != nil {
		return err
	}

	return nil
}
