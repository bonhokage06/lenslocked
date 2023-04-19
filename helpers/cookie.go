package helpers

import (
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"time"
)

func Encode(value string) string {
	tokenHash := sha256.Sum256([]byte(value))
	return base64.URLEncoding.EncodeToString(tokenHash[:])
}

func Decode(value string) string {
	decoded, _ := base64.URLEncoding.DecodeString(value)
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

// delete a cookie
func DeleteCookie(w http.ResponseWriter, name string) {
	cookie := http.Cookie{
		Name:     name,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   -1,
		Expires:  time.Now().Add(-time.Hour * 24 * 365),
	}
	http.SetCookie(w, &cookie)
}
