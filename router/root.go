package router

import (
	"net/http"

	"github.com/bonhokage06/lenslocked/controllers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Router struct {
}

func (router *Router) New() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.CleanPath)
	PageHandler := controllers.Handler
	StaticHandler := controllers.StaticHandler
	ContactHandler := controllers.Contact{}
	FaqHandler := controllers.Faq{}
	r.Get("/", PageHandler(controllers.Index(nil, "home.gohtml", "partials/layout-parts.gohtml")))
	r.Get("/contact", PageHandler(controllers.Index(nil, "contact.gohtml", "partials/layout-parts.gohtml")))
	r.Get("/contact/{id}", PageHandler(controllers.Index(ContactHandler.Create, "contact.gohtml", "partials/layout-parts.gohtml")))
	r.Get("/faq", PageHandler(controllers.Index(FaqHandler.Create, "faq.gohtml", "partials/layout-parts.gohtml")))
	r.Get("/static/*", StaticHandler(controllers.IndexStatic()))
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	return r
}
