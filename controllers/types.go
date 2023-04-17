package controllers

import (
	"html/template"
	"net/http"
)

type Controller interface {
	Create(r *http.Request) interface{}
}

type Home struct {
}
type Contact struct {
}
type ContactData struct {
	Id string
}
type Faq struct {
}
type Questions struct {
	Question string
	Answer   template.HTML
}

type FaqResponse struct {
	Questions []Questions
}
