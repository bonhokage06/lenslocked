package controllers

import (
	"fmt"
	"net/http"

	"github.com/bonhokage06/lenslocked/helpers"
	"github.com/bonhokage06/lenslocked/models"
)

func (u *Users) Index(r *http.Request) (string, interface{}) {
	user := UsersResponse{
		Email: r.FormValue("email"),
	}
	return "", user
}
func (u *Users) Show(r *http.Request) (string, interface{}) {
	userModel := models.User{}
	users, err := userModel.Get()
	if err != nil {
		return "", UsersResponse{
			Errors: []string{err.Error()},
		}
	}
	user := UsersResponse{
		Users:  users,
		Errors: nil,
	}
	return "", user
}
func (u *Users) Create(r *http.Request) (string, interface{}) {
	if r.Method != "POST" {
		return "", UsersResponse{
			Errors: []string{"Method not allowed"},
		}
	}
	email := r.FormValue("email")
	if email == "" && helpers.IsValidEmail(email) {
		return "nil", UsersResponse{
			Errors: []string{"Email is required"},
		}
	}
	password := r.FormValue("password")
	if password == "" {
		return "", UsersResponse{
			Errors: []string{"Password is required"},
		}
	}
	userModel := models.User{
		Email: email,
		Hash:  password,
	}
	err := userModel.Create()
	if err != nil {
		return "", UsersResponse{
			Errors: []string{err.Error()},
		}
	}
	return fmt.Sprintf("/message?status=%s&message=%s", "Success", "User added succcessfully."), nil
}
