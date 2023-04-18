package controllers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/bonhokage06/lenslocked/helpers"
	"github.com/bonhokage06/lenslocked/models"
)

func (u *Users) Index(r *http.Request) ([]http.Cookie, interface{}) {
	user := UsersResponse{
		Email: r.FormValue("email"),
	}
	return nil, user
}
func (u *Users) Show(r *http.Request) ([]http.Cookie, interface{}) {
	userModel := models.User{}
	users, err := userModel.Get()
	if err != nil {
		return nil, UsersResponse{
			Errors: []string{err.Error()},
		}
	}
	user := UsersResponse{
		Users:  users,
		Errors: nil,
	}
	return nil, user
}
func (u *Users) SignIn(r *http.Request) ([]http.Cookie, interface{}) {
	if r.Method != "POST" {
		return nil, UsersResponse{
			Errors: []string{"Method not allowed"},
		}
	}
	email := r.FormValue("email")
	if email == "" && helpers.IsValidEmail(email) {
		return nil, UsersResponse{
			Errors: []string{"Email is required"},
		}
	}
	password := r.FormValue("password")
	if password == "" {
		return nil, UsersResponse{
			Errors: []string{"Password is required"},
		}
	}
	userModel := models.User{
		Email: email,
		Hash:  password,
	}
	isValid := userModel.Authenticate()
	if isValid {
		cookies := []http.Cookie{
			{
				Name:     "Email",
				Value:    email,
				MaxAge:   3600,
				HttpOnly: true,
				SameSite: http.SameSiteLaxMode,
				Path:     "/",
			},
		}
		return cookies, nil
	}
	return nil, UsersResponse{
		Errors: []string{"Invalid email or password"},
	}
}
func (u *Users) SignOut(r *http.Request) ([]http.Cookie, interface{}) {
	//delete cookie Email
	cookies := []http.Cookie{
		{
			Name:     "Email",
			Expires:  time.Now().Add(-7 * 24 * time.Hour),
			HttpOnly: true,
		},
	}
	return cookies, nil
}
func (u *Users) Create(r *http.Request) ([]http.Cookie, interface{}) {
	if r.Method != "POST" {
		return nil, UsersResponse{
			Errors: []string{"Method not allowed"},
		}
	}
	email := r.FormValue("email")
	if email == "" && helpers.IsValidEmail(email) {
		return nil, UsersResponse{
			Errors: []string{"Email is required"},
		}
	}
	password := r.FormValue("password")
	if password == "" {
		return nil, UsersResponse{
			Errors: []string{"Password is required"},
		}
	}
	userModel := models.User{
		Email: email,
		Hash:  password,
	}
	err := userModel.Create()
	if err != nil {
		if strings.Contains(err.Error(), "users_email_key") {
			return nil, UsersResponse{
				Errors: []string{"Email already exist"},
			}
		}
		return nil, UsersResponse{
			Errors: []string{err.Error()},
		}
	}
	cookies := []http.Cookie{
		{
			Name:  "Path",
			Value: fmt.Sprintf("/message?status=%s&message=%s", "Success", "User added succcessfully."),
		},
	}
	return cookies, UsersResponse{}
}
