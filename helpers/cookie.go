package helpers

import (
	"encoding/base64"
	"net/http"
)

func Encode(value string) string {
	return base64.StdEncoding.EncodeToString([]byte(value))
}

func Decode(value string) string {
	decoded, _ := base64.StdEncoding.DecodeString(value)
	return string(decoded)
}

func CreateCookie(w http.ResponseWriter, name string, value string) {
	cookie := http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &cookie)
}
func GetCookie(r *http.Request, name string) string {
	cookie, err := r.Cookie(name)
	if err != nil {
		return ""
	}
	return cookie.Value
}

func SetCookiesAndReturnPath(w http.ResponseWriter, cookies []http.Cookie) string {
	var path string
	for _, cookie := range cookies {
		//dont add path into cookies this is special cookie for redirection
		if cookie.Name == "Path" {
			path = cookie.Value
			continue
		}
		http.SetCookie(w, &cookie)
	}
	return path
}
