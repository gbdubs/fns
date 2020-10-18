package util

import (
	"bytes"
	"text/template"
)

func RenderTemplate(templateFileName string, data interface{}) string {
	t := template.Must(template.ParseGlob(
		"/Users/gradyward/go/src/github.com/gbdubs/fns/tmpl/" + templateFileName))
	b := bytes.Buffer{}
	e := t.Execute(&b, data)
	if e != nil {
		panic(e)
	}
	return b.String()
}
