package models

type User struct {
	Email    string
	Password string
}

func (u User) GetUser() User {
	return u
}
