package templates

import (
	"errors"
	"fmt"
	"html/template"
	"io/fs"
	"strings"
)

func (t *Html) ParseFs(fs fs.FS, patterns ...string) (*template.Template, error) {
	//split patterns[0] to get the name of the template
	fileName := strings.Split(patterns[0], "/")
	tpl := template.New(fileName[len(fileName)-1])
	tpl = tpl.Funcs(template.FuncMap{
		"csrfField": func() (template.HTML, error) {
			return `<-- todo csrf field -->`, errors.New("csrfField not implemented yet")
		},
		"isLogin": func() bool {
			return false
		},
	})
	tpl, err := tpl.ParseFS(fs, patterns...)
	if err != nil {
		return nil, fmt.Errorf("error parsing template: %v", err)
	}
	return tpl, nil
}
