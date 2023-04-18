package controllers

import (
	"html/template"
	"net/http"

	"github.com/bonhokage06/lenslocked/models"
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

type Message struct {
	Status string
	Body   string
}

type Users struct {
}
type UsersResponse struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Users    []models.User
	Errors   []string `json:"errors"`
	Message  Message
}

type MessageResponse struct {
	Message Message
}
