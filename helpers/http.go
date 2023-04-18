package helpers

import (
	"encoding/base64"
	"encoding/json"
	"html/template"
	"net/http"
	"time"
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
				cookie.Value = Encode(cookie.Value)
				http.SetCookie(w, &cookie)
			}
			http.Redirect(w, r, path, http.StatusFound)
			return
		}
		//check if loggin if url is not /auth
		if r.URL.Path != "/auth" {
			cookie, err := r.Cookie("Email")
			if err == nil {
				isLogin := len(Decode(cookie.Value)) > 0
				if isLogin {
					http.Redirect(w, r, "/auth", http.StatusFound)
					return
				}
			}
		}
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

func SetFlash(w http.ResponseWriter, name string, value string) {
	c := &http.Cookie{Name: name, Value: encode([]byte(value))}
	http.SetCookie(w, c)
}

func GetFlash(w http.ResponseWriter, r *http.Request, name string) ([]byte, error) {
	c, err := r.Cookie(name)
	if err != nil {
		switch err {
		case http.ErrNoCookie:
			return nil, nil
		default:
			return nil, err
		}
	}
	value, err := decode(c.Value)
	if err != nil {
		return nil, err
	}
	dc := &http.Cookie{Name: name, MaxAge: -1, Expires: time.Unix(1, 0)}
	http.SetCookie(w, dc)
	return value, nil
}

// -------------------------

func encode(src []byte) string {
	return base64.URLEncoding.EncodeToString(src)
}

func decode(src string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(src)
}
