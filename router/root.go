package router

import (
	"net/http"

	"github.com/bonhokage06/lenslocked/controllers"
	"github.com/bonhokage06/lenslocked/helpers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Router struct {
}

func (router *Router) New() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.CleanPath)
	HtmlHandler := helpers.HtmlHandler
	JsonHandler := helpers.JsonHandler
	StaticHandler := helpers.StaticHandler
	Contact := controllers.Contact{}
	Faq := controllers.Faq{}
	Users := controllers.Users{}
	r.Get("/", HtmlHandler(controllers.Html(nil, "home.gohtml", "partials/layout-parts.gohtml")))
	r.Get("/contact", HtmlHandler(controllers.Html(nil, "contact.gohtml", "partials/layout-parts.gohtml")))
	r.Get("/contact/{id}", HtmlHandler(controllers.Html(Contact.Create, "contact.gohtml", "partials/layout-parts.gohtml")))
	r.Get("/faq", HtmlHandler(controllers.Html(Faq.Create, "faq.gohtml", "partials/layout-parts.gohtml")))
	r.Get("/signup", HtmlHandler(controllers.Html(Users.Index, "users/new.gohtml", "partials/layout-parts.gohtml")))
	r.Post("/signup", JsonHandler((controllers.Json(Users.Create))))
	r.Get("/static/*", StaticHandler(controllers.Static()))
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	return r
}
