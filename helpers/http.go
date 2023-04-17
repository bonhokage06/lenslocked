package helpers

import (
	"encoding/json"
	"html/template"
	"net/http"
)

type Page struct {
	Template *template.Template
	DataFunc func(r *http.Request) interface{}
}
type Json struct {
	DataFunc func(r *http.Request) interface{}
}
type Static struct {
	Fs http.Handler
}

func Headers(w http.ResponseWriter, contentType string) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(http.StatusOK)
}

func HtmlHandler(s Page) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if s.DataFunc == nil {
			s.DataFunc = func(r *http.Request) interface{} {
				return nil
			}
		}
		data := s.DataFunc(r)
		err := s.Template.Execute(w, data)
		if err != nil {
			panic(err)
		}
	}
}
func JsonHandler(s Json) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if s.DataFunc == nil {
			s.DataFunc = func(r *http.Request) interface{} {
				return nil
			}
		}
		data := s.DataFunc(r)
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
		s.Fs.ServeHTTP(w, r)
	}
}
