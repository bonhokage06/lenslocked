package controllers

import (
	"net/http"

	"github.com/bonhokage06/lenslocked/helpers"
	"github.com/bonhokage06/lenslocked/templates"
	"github.com/bonhokage06/lenslocked/views/pages"
)

func Index(dataFunc func(r *http.Request) interface{}, path ...string) helpers.Page {
	template := templates.Html{}
	tpl, err := template.ParseFs(pages.FS, path...)
	if err != nil {
		panic(err)
	}
	return helpers.Page{
		Template: tpl,
		DataFunc: dataFunc,
	}
}
func IndexJson(dataFunc func(r *http.Request) interface{}) helpers.Json {
	return helpers.Json{
		DataFunc: dataFunc,
	}
}
func IndexStatic() helpers.Static {
	fs := http.FileServer(http.FS(pages.StaticFs))
	return helpers.Static{
		Fs: fs,
	}
}
