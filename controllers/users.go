package controllers

import (
	"net/http"
)

type Users struct {
}
type UsersResponse struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *Users) Index(r *http.Request) interface{} {
	user := UsersResponse{
		Email: r.FormValue("email"),
	}
	return user
}
func (u *Users) Create(r *http.Request) interface{} {
	user := UsersResponse{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}
	return user
}
