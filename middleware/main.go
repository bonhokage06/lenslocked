package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/bonhokage06/lenslocked/helpers"
	"github.com/bonhokage06/lenslocked/models"
	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/csrf"
)

// create auth middleware
func IsAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//do something
		rememberToken := helpers.GetCookie(r, "remember_token")
		if !strings.Contains(r.URL.Path, "/auth") {
			if len(rememberToken) == 0 {
				next.ServeHTTP(w, r)
				return
			}
			sessionModel := models.Session{
				RememberToken: rememberToken,
			}
			session, err := sessionModel.Check()
			fmt.Println(session.Email)
			if err == nil {
				isLogin := session.Email != ""
				if isLogin {
					http.Redirect(w, r, "/auth", http.StatusFound)
					return
				}
			} else {
				helpers.DeleteCookie(w, "remember_token")
				http.Redirect(w, r, "/", http.StatusFound)
			}
		}
		if strings.Contains(r.URL.Path, "/auth") {
			if len(rememberToken) == 0 {
				http.Redirect(w, r, "/", http.StatusFound)
				return
			}
			sessionModel := models.Session{
				RememberToken: rememberToken,
			}
			session, err := sessionModel.Check()
			if err != nil || session.Email == "" {
				helpers.DeleteCookie(w, "remember_token")
				http.Redirect(w, r, "/", http.StatusFound)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

func Logger() func(http.Handler) http.Handler {
	return middleware.Logger
}

func CleanPath() func(http.Handler) http.Handler {
	return middleware.CleanPath
}
func Csrf(h http.Handler) http.Handler {
	csrfKey := []byte("spelspaelspel2soekslo30soe3scwade")
	csrfMiddleware := csrf.Protect(csrfKey, csrf.Secure(false))
	return csrfMiddleware(h)
}
