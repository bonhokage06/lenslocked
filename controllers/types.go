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
	IsLogin   bool
	Questions []Questions
}

type Message struct {
	Status string
	Body   string
}
type MessageResponse struct {
	Message Message
	IsLogin bool
}
type Users struct {
}
type UsersResponse struct {
	Email  string `json:"email"`
	Users  []models.User
	Errors []string `json:"errors"`
}

type Auth struct {
}
type AuthResponse struct {
	Errors []string `json:"errors"`
}
