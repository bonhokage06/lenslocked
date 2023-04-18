package controllers

import "net/http"

func (m *Message) Index(r *http.Request) ([]http.Cookie, interface{}) {

	return nil, MessageResponse{
		Message: Message{
			Status: r.FormValue("status"),
			Body:   r.FormValue("message"),
		},
	}
}
