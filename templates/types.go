package templates

import (
	"html/template"
	"net/http"
)

type Html struct {
	Writer   http.ResponseWriter
	Template *template.Template
}
