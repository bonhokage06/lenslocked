package models

type User struct {
	Name  string
	Users map[string]string
}

func (u User) GetUser() User {
	return u
}
