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
	ContactHandler := controllers.Contact{}
	FaqHandler := controllers.Faq{}
	UsersHandler := controllers.Users{}
	r.Get("/", HtmlHandler(controllers.Index(nil, "home.gohtml", "partials/layout-parts.gohtml")))
	r.Get("/contact", HtmlHandler(controllers.Index(nil, "contact.gohtml", "partials/layout-parts.gohtml")))
	r.Get("/contact/{id}", HtmlHandler(controllers.Index(ContactHandler.Create, "contact.gohtml", "partials/layout-parts.gohtml")))
	r.Get("/faq", HtmlHandler(controllers.Index(FaqHandler.Create, "faq.gohtml", "partials/layout-parts.gohtml")))
	r.Get("/signup", HtmlHandler(controllers.Index(UsersHandler.Index, "users/new.gohtml", "partials/layout-parts.gohtml")))
	r.Post("/signup", JsonHandler((controllers.IndexJson(UsersHandler.Create))))
	r.Get("/static/*", StaticHandler(controllers.IndexStatic()))
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	return r
}
