package pages

import "embed"

//go:embed *.gohtml partials/*.gohtml users/*.gohtml
var FS embed.FS

//go:embed static/js/*.js
var StaticFs embed.FS
