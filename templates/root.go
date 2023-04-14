package templates

import (
	"fmt"
	"html/template"
	"io/fs"

	"github.com/bonhokage06/lenslocked/constants"
	"github.com/bonhokage06/lenslocked/helpers"
)

func (t *Html) Parse(path string) (*template.Template, error) {
	tpl, err := t.Template.ParseFiles(path)
	if err != nil {
		return nil, fmt.Errorf("error parsing template: %v", err)
	}
	return tpl, nil
}
func (t *Html) ParseFs(fs fs.FS, patterns ...string) (*template.Template, error) {
	tpl, err := t.Template.ParseFS(fs, patterns...)
	if err != nil {
		return nil, fmt.Errorf("error parsing template: %v", err)
	}
	return tpl, nil
}
func (t *Html) Execute(data interface{}) error {
	helpers.Headers(t.Writer, constants.TextHtml)
	err := t.Template.Execute(t.Writer, data)
	if err != nil {
		return fmt.Errorf("error executing template: %v", err)
	}
	return nil
}
