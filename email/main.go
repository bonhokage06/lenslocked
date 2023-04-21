package email

import (
	"log"

	"github.com/bonhokage06/lenslocked/constants"
	"github.com/go-mail/mail/v2"
)

type Email struct {
	To        string
	Subject   string
	Plaintext string
	Html      string
}

func Send(email Email) {
	from := "test@lenslocked.com"
	subject := email.Subject
	plaintext := email.Plaintext
	html := email.Html

	msg := mail.NewMessage()
	msg.SetHeader("To", email.To)
	msg.SetHeader("From", from)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/plain", plaintext)
	msg.AddAlternative("text/html", html)
	dialer := mail.NewDialer(constants.Host, constants.Port, constants.Username, constants.Password)
	err := dialer.DialAndSend(msg)
	if err != nil {
		log.Printf("error sending email: %s", err.Error())
	}
}
