package controllers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	emailService "github.com/bonhokage06/lenslocked/email"
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
	isValid, User := userModel.Authenticate()
	if isValid {
		sessionModel := models.Session{
			UserId: User.Id,
		}
		session, err := sessionModel.Create()
		if err != nil {
			return nil, UsersResponse{
				Errors: []string{err.Error()},
			}
		}
		cookies := []http.Cookie{
			{
				Name:     "remember_token",
				Value:    session.RememberToken,
				MaxAge:   3600,
				HttpOnly: true,
				Secure:   true,
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
	//get remember_token from cookie
	cookie, err := r.Cookie("remember_token")
	if err != nil {
		return nil, UsersResponse{
			Errors: []string{err.Error()},
		}
	}
	//delete session
	sessionModel := models.Session{
		RememberToken: cookie.Value,
	}
	err = sessionModel.Delete()
	if err != nil {
		return nil, UsersResponse{
			Errors: []string{err.Error()},
		}
	}
	//delete cookie remember_token
	cookies := []http.Cookie{
		{
			Name:     "remember_token",
			Value:    "",
			Path:     "/",
			Expires:  time.Now().Add(-time.Hour * 24 * 365),
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteLaxMode,
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
			Name:     "Path",
			Value:    fmt.Sprintf("/message?status=%s&message=%s", "Success", "User added succcessfully."),
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteLaxMode,
		},
	}
	return cookies, UsersResponse{}
}

// a function the implements forgot password and send email
func (u *Users) ForgotPassword(r *http.Request) ([]http.Cookie, interface{}) {
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
	go func() {
		emailService.Send(email)
	}()
	cookies := []http.Cookie{
		{
			Name:     "Path",
			Value:    fmt.Sprintf("/message?status=%s&message=%s", "Success", "Email sent successfully."),
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteLaxMode,
		},
	}
	return cookies, UsersResponse{}
}
