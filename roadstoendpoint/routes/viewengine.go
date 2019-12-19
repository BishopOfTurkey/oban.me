package routes

import (
	"html/template"
	"io"
	"os"
	"path/filepath"
)

func renderView(w io.Writer, path string, fmap *template.FuncMap, v interface{}) error {
	_, name := filepath.Split(path)
	wd, _ := os.Getwd()

	if fmap == nil {
		fmap = &template.FuncMap{}
	}
	t, err := template.New(name).Funcs(*fmap).ParseFiles(filepath.Join(wd, path))
	if err != nil {
		return err
	}
	return t.Execute(w, v)
}
