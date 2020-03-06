package lib

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"regexp"
	"time"
)

var t *template.Template
var RegexEmail = regexp.MustCompile(`^\w.+@(\w)+\.\w+$`)

// ParseTemplate parse html file from views folder with data
func ParseTemplate(w http.ResponseWriter, pathfile string, data interface{}) error {
	t := template.New("base").Funcs(template.FuncMap{
		"formatDate": func(value time.Time, layout string) string {
			return value.Format(layout)
		},
	})
	t, err := t.ParseFiles("views/base.html", "views/"+pathfile+".html")
	if err != nil {
		return err
	}

	err = t.ExecuteTemplate(w, "base", data)
	if err != nil {
		return err
	}
	return nil
}

// GetPort get value from PORT env
func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("Err when open PORT env: ", port)
	}

	return ":" + port
}