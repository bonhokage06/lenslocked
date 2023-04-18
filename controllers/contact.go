package controllers

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (c *Contact) Create(r *http.Request) ([]http.Cookie, interface{}) {
	id := chi.URLParam(r, "id")
	return nil, ContactData{
		Id: id,
	}
}
