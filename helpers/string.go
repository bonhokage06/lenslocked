package helpers

import (
	"net/mail"
)

// check if valid email
func IsValidEmail(email string) bool {
	//generate regex that check valid email
	//return true if valid email
	_, err := mail.ParseAddress(email)
	return err != nil
}
