package lib

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
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

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("Err when open PORT env: ", port)
	}

	return ":" + port
}
