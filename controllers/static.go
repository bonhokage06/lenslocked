package controllers

import (
	"html/template"
	"net/http"
)

type Page struct {
	template *template.Template
	dataFnc  func(r *http.Request) interface{}
}
type Static struct {
	fs http.Handler
}

func Handler(s Page) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if s.dataFnc == nil {
			s.dataFnc = func(r *http.Request) interface{} {
				return nil
			}
		}
		data := s.dataFnc(r)
		err := s.template.Execute(w, data)
		if err != nil {
			panic(err)
		}
	}
}

func StaticHandler(s Static) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/javascript")
		w.Header().Set("Referrer-Policy", "no-referrer")
		s.fs.ServeHTTP(w, r)
	}
}
