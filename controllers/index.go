package controllers

import (
	"net/http"

	"github.com/bonhokage06/lenslocked/templates"
	"github.com/bonhokage06/lenslocked/views/pages"
)

func Index(dataFunc func(r *http.Request) interface{}, path ...string) Page {
	template := templates.Html{}
	tpl, err := template.ParseFs(pages.FS, path...)
	if err != nil {
		panic(err)
	}
	return Page{
		template: tpl,
		dataFnc:  dataFunc,
	}
}
func IndexJson(dataFunc func(r *http.Request) interface{}) Json {
	return Json{
		dataFnc: dataFunc,
	}
}
func IndexStatic() Static {
	fs := http.FileServer(http.FS(pages.StaticFs))
	return Static{
		fs: fs,
	}
}
