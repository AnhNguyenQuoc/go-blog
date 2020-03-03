package lib

import (
	"html/template"
	"net/http"
)

var t *template.Template

func ParseTemplate(w http.ResponseWriter, pathfile string, data interface{}) error {
	t, err := t.ParseFiles("views/base.html", pathfile+".html")
	if err != nil {
		return err
	}

	t.ExecuteTemplate(w, "base", data)
	return nil
}
