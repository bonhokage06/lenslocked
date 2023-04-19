package controllers

import "net/http"

func (a *Auth) Index(r *http.Request) ([]http.Cookie, interface{}) {
	_, err := r.Cookie("Email")
	if err != nil {
		cookies := []http.Cookie{
			{
				Name:  "Path",
				Value: "/",
				Path:  "/",
			},
		}
		return cookies, nil
	}
	return nil, nil
}
