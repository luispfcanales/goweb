package gotemplate

import (
	"io"
	"log"
	"text/template"
)

//GoTemplate is template to render
type GoTemplate struct {
	tmpl *template.Template
}

//New return golangTemplate
func New(path string) (*GoTemplate, error) {
	t, err := template.New("").ParseGlob(path)
	if err != nil {
		return &GoTemplate{}, err
	}
	return &GoTemplate{t}, nil
}

//Render execute template to render webpage
func (t *GoTemplate) Render(writer io.Writer, name string, data interface{}) {
	err := t.tmpl.ExecuteTemplate(writer, name, data)
	if err != nil {
		log.Fatalf("error to execute template %v", err)
	}
}
