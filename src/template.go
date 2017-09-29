package src

import (
	"html/template"
	"io"
	"os"
	"path/filepath"

	"github.com/teambition/gear"
	"github.com/teambition/gear/logging"
)

var _ gear.Renderer = &Template{}

// Template struct implement gear Renderer interface
type Template struct {
}

// Render implement gear Renderer interface Render
func (t *Template) Render(ctx *gear.Context, w io.Writer, name string, data interface{}) (err error) {
	dir, _ := os.Getwd()
	name = filepath.Join(dir, "view", name+".html")
	tmpl := template.Must(template.ParseFiles(name))

	err = tmpl.Execute(w, data)
	if err != nil {
		logging.Println(err)
	}

	return
}
