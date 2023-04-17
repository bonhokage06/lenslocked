package controllers

import (
	"encoding/json"
	"html/template"
	"net/http"
)

type Page struct {
	template *template.Template
	dataFnc  func(r *http.Request) interface{}
}
type Json struct {
	dataFnc func(r *http.Request) interface{}
}
type Static struct {
	fs http.Handler
}

func HtmlHandler(s Page) http.HandlerFunc {
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
func JsonHandler(s Json) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if s.dataFnc == nil {
			s.dataFnc = func(r *http.Request) interface{} {
				return nil
			}
		}
		data := s.dataFnc(r)
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(data)
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
