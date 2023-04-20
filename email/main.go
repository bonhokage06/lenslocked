package email

import (
	"log"

	"github.com/bonhokage06/lenslocked/constants"
	"github.com/go-mail/mail/v2"
)

func Send(to string) {
	from := "test@lenslocked.com"
	subject := "This is a test email"
	plaintext := "This is the body of the email"
	html := `<h1>Hello there buddy!</h1><p>This is the email</p><p>Hope you enjoy it</p>`

	msg := mail.NewMessage()
	msg.SetHeader("To", to)
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
