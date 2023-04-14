package controllers

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (c *Contact) Create(r *http.Request) interface{} {
	id := chi.URLParam(r, "id")
	return ContactData{
		Id: id,
	}
}
