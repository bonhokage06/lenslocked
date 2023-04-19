package helpers

import (
	"bytes"
	"encoding/json"
	"html/template"
	"io"
	"net/http"

	"github.com/gorilla/csrf"
)

type Page struct {
	Template *template.Template
	DataFunc func(r *http.Request) ([]http.Cookie, interface{})
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
		// create DataFunc if not exist
		if s.DataFunc == nil {
			s.DataFunc = func(r *http.Request) ([]http.Cookie, interface{}) {
				return nil, nil
			}
		}
		cookies, data := s.DataFunc(r)
		if len(cookies) > 0 {
			var path string
			for _, cookie := range cookies {
				//dont add path into cookies this is special cookie for redirection
				if cookie.Name == "Path" {
					path = cookie.Value
					continue
				}
				http.SetCookie(w, &cookie)
			}
			http.Redirect(w, r, path, http.StatusFound)
			return
		}
		//add csrf token
		tpl, err := s.Template.Clone()
		if err != nil {
			http.Error(w, "There was an error rendering the page.", http.StatusInternalServerError)
		}
		tpl.Funcs(template.FuncMap{
			"csrfField": func() template.HTML {
				return csrf.TemplateField(r)
			},
			"isLogin": func() bool {
				cookie, err := r.Cookie("remember_token")
				if err != nil {
					return false
				}
				return len(cookie.Value) > 0
			},
		})
		// execute template
		var buf bytes.Buffer
		err = tpl.Execute(&buf, data)
		if err != nil {
			http.Error(w, "There was an error rendering the page.", http.StatusInternalServerError)
			return
		}
		io.Copy(w, &buf)
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
