package helpers

import "net/http"

func Headers(w http.ResponseWriter, contentType string) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(http.StatusOK)
}
