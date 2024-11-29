package mail

import (
	"fmt"
	"net/smtp"
)

type Mail struct {
	from     string
	password string
	host     string
	port     string
}

type Mailer interface {
	Send(typ, to, subject, content string) error
}

func New() *Mail {
	return &Mail{
		host: "smtp.gmail.com",
		port: "587",
	}
}

func (m *Mail) Send(typ, to, subject, content string) error {
	message := []byte(subject + "\n" + content)

	auth := smtp.PlainAuth("", m.from, m.password, m.host)

	err := smtp.SendMail(m.host+":"+m.port, auth, m.from, []string{to}, message)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	return nil
}
