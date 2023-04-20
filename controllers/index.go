package controllers

import (
	"net/http"

	middlewares "github.com/bonhokage06/lenslocked/middleware"
	"github.com/bonhokage06/lenslocked/templates"
	"github.com/bonhokage06/lenslocked/views/pages"
)

func Html(dataFunc func(r *http.Request) ([]http.Cookie, interface{}), path ...string) middlewares.Page {
	htmlTemplate := templates.Html{}
	tpl, err := htmlTemplate.ParseFs(pages.FS, path...)
	if err != nil {
		panic(err)
	}
	return middlewares.Page{
		Template: tpl,
		DataFunc: dataFunc,
	}
}
func Json(dataFunc func(r *http.Request) interface{}) middlewares.Json {
	return middlewares.Json{
		DataFunc: dataFunc,
	}
}
func Static() middlewares.Static {
	fs := http.FileServer(http.FS(pages.StaticFs))
	return middlewares.Static{
		Fs: fs,
	}
}
