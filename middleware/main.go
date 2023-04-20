package middlewares

import (
	"fmt"
	"net/http"

	"github.com/bonhokage06/lenslocked/context"
	"github.com/bonhokage06/lenslocked/helpers"
	"github.com/bonhokage06/lenslocked/models"
	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/csrf"
)

// create check user middleware
func SetUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rememberToken := helpers.GetCookie(r, "remember_token")
		sessionModel := models.Session{
			RememberToken: rememberToken,
		}
		session, _ := sessionModel.Check()
		ctx := context.WithUser(r.Context(), &session)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

// create auth middleware
func IsAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := context.User(r.Context())
		fmt.Println(user)
		if user != nil {
			isLoggin := user.Email != ""
			if isLoggin {
				http.Redirect(w, r, "/auth", http.StatusFound)
				return
			} else {
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
