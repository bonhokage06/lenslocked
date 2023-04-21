package controllers

import (
	"fmt"
	"net/http"

	"github.com/bonhokage06/lenslocked/models"
	"github.com/go-chi/chi"
)

func (p *PasswordReset) Index(r *http.Request) ([]http.Cookie, interface{}) {
	token := chi.URLParam(r, "token")
	if len(token) == 0 {
		cookies := []http.Cookie{
			{
				Name:     "Path",
				Value:    fmt.Sprintf("/message?status=%s&message=%s", "Error", "Token is required"),
				HttpOnly: true,
				Secure:   true,
				SameSite: http.SameSiteLaxMode,
			},
		}
		return cookies, nil
	}
	//get token from the url
	passwordResetModel := models.PasswordReset{
		TokenHash: token,
	}
	passwordResetUser := passwordResetModel.Check()
	if passwordResetUser.Email == "" {
		cookies := []http.Cookie{
			{
				Name:     "Path",
				Value:    fmt.Sprintf("/message?status=%s&message=%s", "Error", "Token is invalid"),
				HttpOnly: true,
				Secure:   true,
				SameSite: http.SameSiteLaxMode,
			},
		}
		return cookies, nil
	}
	passwordResetModel.Delete()
	return nil, PasswordResetResponse{
		Email:  passwordResetUser.Email,
		Errors: nil,
	}
}
func (p *PasswordReset) Update(r *http.Request) ([]http.Cookie, interface{}) {
	//check method
	if r.Method != "POST" {
		return nil, PasswordResetResponse{
			Errors: []string{"Method not allowed"},
		}
	}
	//check email
	email := r.FormValue("email")
	if email == "" {
		return nil, PasswordResetResponse{
			Errors: []string{"Email is required"},
		}
	}
	//check password
	password := r.FormValue("password")
	if password == "" {
		return nil, PasswordResetResponse{
			Errors: []string{"Password is required"},
		}
	}
	userModal := models.User{
		Email: email,
		Hash:  password,
	}
	err := userModal.ChangePassword()
	if err != nil {
		return nil, PasswordResetResponse{
			Errors: []string{err.Error()},
		}
	}
	cookies := []http.Cookie{
		{
			Name:     "Path",
			Value:    fmt.Sprintf("/message?status=%s&message=%s", "Success", "User updated succcessfully."),
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteLaxMode,
		},
	}
	return cookies, nil
}
