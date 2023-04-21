package router

import (
	"net/http"

	"github.com/bonhokage06/lenslocked/controllers"
	middleware "github.com/bonhokage06/lenslocked/middleware"
	"github.com/go-chi/chi"
)

type Router struct {
}

func (router *Router) New() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger())
	r.Use(middleware.CleanPath())
	r.Use(middleware.Csrf)
	r.Use(middleware.SetUser)
	r.Use(middleware.IsAuth)
	HtmlHandler := middleware.HtmlHandler
	StaticHandler := middleware.StaticHandler
	Contact := controllers.Contact{}
	Faq := controllers.Faq{}
	Users := controllers.Users{}
	Auth := controllers.Auth{}
	Message := controllers.Message{}
	Passwordreset := controllers.PasswordReset{}
	r.Get("/", HtmlHandler(controllers.Html(nil, "home.gohtml", "partials/layout-parts.gohtml")))
	r.Get("/contact", HtmlHandler(controllers.Html(nil, "contact.gohtml", "partials/layout-parts.gohtml")))
	r.Get("/contact/{id}", HtmlHandler(controllers.Html(Contact.Create, "contact.gohtml", "partials/layout-parts.gohtml")))
	r.Get("/faq", HtmlHandler(controllers.Html(Faq.Create, "faq.gohtml", "partials/layout-parts.gohtml")))
	r.Get("/signup", HtmlHandler(controllers.Html(Users.Index, "users/new.gohtml", "partials/*")))
	r.Get("/signin", HtmlHandler(controllers.Html(Users.Index, "users/signin.gohtml", "partials/*")))
	r.Get("/reset-password", HtmlHandler(controllers.Html(Users.Index, "users/forgot.gohtml", "partials/*")))
	r.Get("/reset-password/{token}", HtmlHandler(controllers.Html(Passwordreset.Index, "users/reset.gohtml", "partials/*")))
	r.Get("/reset-password/{token}", HtmlHandler(controllers.Html(Passwordreset.Index, "users/reset.gohtml", "partials/*")))
	r.Post("/signin", HtmlHandler((controllers.Html(Users.SignIn, "users/signin.gohtml", "partials/*"))))
	r.Post("/auth/signout", HtmlHandler(controllers.Html(Users.SignOut, "home.gohtml", "partials/*")))
	r.Post("/users/create", HtmlHandler((controllers.Html(Users.Create, "users/new.gohtml", "partials/*"))))
	r.Post("/users/forgot", HtmlHandler((controllers.Html(Users.ForgotPassword, "users/forgot.gohtml", "partials/*"))))
	r.Post("/users/reset", HtmlHandler((controllers.Html(Passwordreset.Update, "users/reset.gohtml", "partials/*"))))
	r.Get("/message", HtmlHandler((controllers.Html(Message.Index, "partials/message.gohtml", "partials/layout-parts.gohtml"))))
	r.Get("/users", HtmlHandler(controllers.Html(Users.Show, "users/list.gohtml", "partials/layout-parts.gohtml")))
	r.Get("/auth", HtmlHandler(controllers.Html(Auth.Index, "auth/index.gohtml", "partials/layout-parts.gohtml")))
	r.Get("/static/*", StaticHandler(controllers.Static()))
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	return r
}
