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
	// err := passwordResetModel.Delete()
	// if err != nil {
	// 	return nil, PasswordResetResponse{
	// 		Errors: []string{err.Error()},
	// 	}
	// }
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
	return nil, PasswordResetResponse{
		Email:  passwordResetUser.Email,
		Errors: nil,
	}
}
