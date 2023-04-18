package controllers

import "net/http"

func (m *Message) Index(r *http.Request) (string, interface{}) {

	return "", MessageResponse{
		Message: Message{
			Status: r.FormValue("status"),
			Body:   r.FormValue("message"),
		},
	}
}
