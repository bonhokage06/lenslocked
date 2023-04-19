package router

import (
	"net/http"
	"strings"

	"github.com/bonhokage06/lenslocked/controllers"
	"github.com/bonhokage06/lenslocked/helpers"
	"github.com/bonhokage06/lenslocked/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/csrf"
)

type Router struct {
}

func (router *Router) New() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.CleanPath)
	HtmlHandler := helpers.HtmlHandler
	StaticHandler := helpers.StaticHandler
	Contact := controllers.Contact{}
	Faq := controllers.Faq{}
	Users := controllers.Users{}
	Auth := controllers.Auth{}
	Message := controllers.Message{}
	r.Get("/", HtmlHandler(controllers.Html(nil, "home.gohtml", "partials/layout-parts.gohtml")))
	r.Get("/contact", HtmlHandler(controllers.Html(nil, "contact.gohtml", "partials/layout-parts.gohtml")))
	r.Get("/contact/{id}", HtmlHandler(controllers.Html(Contact.Create, "contact.gohtml", "partials/layout-parts.gohtml")))
	r.Get("/faq", HtmlHandler(controllers.Html(Faq.Create, "faq.gohtml", "partials/layout-parts.gohtml")))
	r.Get("/signup", HtmlHandler(controllers.Html(Users.Index, "users/new.gohtml", "partials/*")))
	r.Get("/signin", HtmlHandler(controllers.Html(Users.Index, "users/signin.gohtml", "partials/*")))
	r.Post("/signin", HtmlHandler((controllers.Html(Users.SignIn, "users/signin.gohtml", "partials/*"))))
	r.Post("/auth/signout", HtmlHandler(controllers.Html(Users.SignOut, "home.gohtml", "partials/*")))
	r.Post("/users/create", HtmlHandler((controllers.Html(Users.Create, "users/new.gohtml", "partials/*"))))
	r.Get("/message", HtmlHandler((controllers.Html(Message.Index, "partials/message.gohtml", "partials/layout-parts.gohtml"))))
	r.Get("/users", HtmlHandler(controllers.Html(Users.Show, "users/list.gohtml", "partials/layout-parts.gohtml")))
	r.Get("/auth", HtmlHandler(controllers.Html(Auth.Index, "auth/index.gohtml", "partials/layout-parts.gohtml")))
	r.Get("/static/*", StaticHandler(controllers.Static()))
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	csrfKey := []byte("spelspaelspel2soekslo30soe3scwade")
	csrfMiddleware := csrf.Protect(csrfKey, csrf.Secure(false))
	return csrfMiddleware(AuthMiddleware(r))
}

// create auth middleware
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//do something
		if !strings.Contains(r.URL.Path, "/auth") {
			cookie, err := r.Cookie("remember_token")
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			sessionModel := models.Session{
				RememberToken: cookie.Value,
			}
			isValidSession, err := sessionModel.Check()
			if err == nil {
				isLogin := isValidSession
				if isLogin {
					http.Redirect(w, r, "/auth", http.StatusFound)
					return
				}
			}
		}
		if strings.Contains(r.URL.Path, "/auth") {
			cookie, err := r.Cookie("remember_token")
			if err != nil {
				http.Redirect(w, r, "/", http.StatusFound)
				return
			}
			sessionModel := models.Session{
				RememberToken: cookie.Value,
			}
			isValidSession, err := sessionModel.Check()
			if err != nil || !isValidSession {
				http.Redirect(w, r, "/", http.StatusFound)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
